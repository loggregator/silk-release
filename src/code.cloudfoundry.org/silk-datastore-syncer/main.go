package main

import (
	"encoding/json"
	"flag"
	"log"
	"reflect"
	"sync"
	"time"

	"code.cloudfoundry.org/executor"
	"code.cloudfoundry.org/filelock"
	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/garden/client"
	"code.cloudfoundry.org/garden/client/connection"
	"code.cloudfoundry.org/lib/datastore"
	"code.cloudfoundry.org/lib/serial"
)

const (
	datastoreFilePath  = "/var/vcap/data/container-metadata/store.json"
	datastoreFileOwner = "vcap"
	datastoreFileGroup = "vcap"
)

var (
	interval      *int
	gardenNetwork *string
	gardenAddr    *string
)

func init() {
	interval = flag.Int("n", 30, "sync interval in seconds.")
	gardenNetwork = flag.String("gardenNetwork", "", "garden network type.")
	gardenAddr = flag.String("gardenAddr", "", "garden address.")
}

func main() {
	flag.Parse()

	log.Println("===== Properties =====")
	log.Println("interval:", *interval)
	log.Println("gardenNetwork:", *gardenNetwork)
	log.Println("gardenAddr:", *gardenAddr)

	gardenClient := client.New(connection.New(*gardenNetwork, *gardenAddr))

	// wait for garden to come alive
	for {
		log.Println("Attempting to ping Garden")
		err := gardenClient.Ping()
		if err == nil {
			break
		}
		switch err.(type) {
		case nil:
			break
		case garden.UnrecoverableError:
			log.Fatalln("Garden: unrecoverable:", err)
		default:
			log.Println("Garden: cannot connect:", err)
			time.Sleep(1 * time.Second)
		}
	}

	store := &datastore.Store{
		Serializer: &serial.Serial{},
		Locker: &filelock.Locker{
			FileLocker: filelock.NewLocker(datastoreFilePath + "_lock"),
			Mutex:      new(sync.Mutex),
		},
		DataFilePath:    datastoreFilePath,
		VersionFilePath: datastoreFilePath + "_version",
		LockedFilePath:  datastoreFilePath + "_lock",
		FileOwner:       datastoreFileOwner,
		FileGroup:       datastoreFileGroup,
		CacheMutex:      new(sync.RWMutex),
	}

	duration := time.Duration(*interval) * time.Second
	for {
		time.Sleep(duration)
		log.Println("Starting sync loop")

		gardenContainers, err := gardenClient.Containers(nil)
		if err != nil {
			log.Println("Garden: error retrieving containers:", err)
			continue
		}

		storeContainers, err := store.ReadAll()
		if err != nil {
			log.Println("Datastore: error retrieving containers from datastore:", err)
			continue
		}

		for _, c := range gardenContainers {
			var desiredLogConfig executor.LogConfig
			props, err := c.Properties()
			if err != nil {
				log.Printf("Garden container: %s: error retrieving properties: %s", c.Handle(), err.Error())
				continue
			}
			logConfigStr, ok := props["log_config"]
			if ok {
				err := json.Unmarshal([]byte(logConfigStr), &desiredLogConfig)
				if err != nil {
					log.Printf("Garden container: %s: error unmarshalling container log config from datastore: %s", c.Handle(), err.Error())
					continue
				}
			}
			log.Printf("Garden container: log config: %#v", desiredLogConfig)

			sc := storeContainers[c.Handle()]
			var actualLogConfig executor.LogConfig
			logConfigStr, ok = sc.Metadata["log_config"].(string)
			if ok {
				err := json.Unmarshal([]byte(logConfigStr), &actualLogConfig)
				if err != nil {
					log.Printf("Datastore container: %s: error unmarshalling container log config from datastore: %s", sc.Handle, err.Error())
					continue
				}
			}
			log.Printf("Datastore container: log config: %#v", actualLogConfig)

			if reflect.DeepEqual(desiredLogConfig, actualLogConfig) {
				log.Println("They are equal.. moving on")
				continue
			}

			log.Printf("Datastore container: %s: reconciling with Garden container", sc.Handle)

			b, err := json.Marshal(desiredLogConfig)
			if err != nil {
				log.Printf("Garden container: %s: error marshalling container log config: %s", sc.Handle, err.Error())
				continue
			}
			sc.Metadata["log_config"] = string(b)
			store.Add(sc.Handle, sc.IP, sc.Metadata)
		}
	}
}
