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

var (
	interval           *int
	gardenNetwork      *string
	gardenAddr         *string
	silkFile           *string
	datastoreFileOwner *string
	datastoreFileGroup *string
)

func init() {
	interval = flag.Int("n", 30, "sync interval in seconds.")
	gardenNetwork = flag.String("gardenNetwork", "", "garden network type.")
	gardenAddr = flag.String("gardenAddr", "", "garden address.")
	silkFile = flag.String("silkFile", "", "silk file.")
	datastoreFileOwner = flag.String("silkFileOwner", "", "owner of silk file")
	datastoreFileGroup = flag.String("silkFileGroup", "", "group owner of silk file")
}

func main() {
	flag.Parse()

	log.Println("===== Properties =====")
	log.Println("interval:", *interval)
	log.Println("gardenNetwork:", *gardenNetwork)
	log.Println("gardenAddr:", *gardenAddr)

	gardenClient := client.New(connection.New(*gardenNetwork, *gardenAddr))
	WaitForGarden(gardenClient)
	store := makeDatastore()

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
			desiredLogConfig, err := getGardenLogConfig(c)
			if err != nil {
				continue
			}

			sc := storeContainers[c.Handle()]
			actualLogConfig, err := getSilkLogConfig(sc)
			if err != nil {
				continue
			}

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
			err = store.Add(sc.Handle, sc.IP, sc.Metadata)
			if err != nil {
				log.Printf("Error updating log config %s", err.Error())
			}
		}
	}
}

func WaitForGarden(gardenClient garden.Client) {
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

}
func makeDatastore() *datastore.Store {
	store := &datastore.Store{
		Serializer: &serial.Serial{},
		Locker: &filelock.Locker{
			FileLocker: filelock.NewLocker(*silkFile + "_lock"),
			Mutex:      new(sync.Mutex),
		},
		DataFilePath:    *silkFile,
		VersionFilePath: *silkFile + "_version",
		LockedFilePath:  *silkFile + "_lock",
		FileOwner:       *datastoreFileOwner,
		FileGroup:       *datastoreFileGroup,
		CacheMutex:      new(sync.RWMutex),
	}
	return store
}

func getGardenLogConfig(c garden.Container) (executor.LogConfig, error) {
	props, err := c.Properties()
	if err != nil {
		log.Printf("Garden container: %s: error retrieving properties: %s", c.Handle(), err.Error())
		return executor.LogConfig{}, err
	}
	logConfigStr, ok := props["log_config"]
	var desiredLogConfig executor.LogConfig
	if ok {
		err := json.Unmarshal([]byte(logConfigStr), &desiredLogConfig)
		if err != nil {
			log.Printf("Garden container: %s: error unmarshalling container log config from datastore: %s", c.Handle(), err.Error())
			return executor.LogConfig{}, err
		}
	}
	log.Printf("Garden container: log config: %#v", desiredLogConfig)
	return desiredLogConfig, nil
}

func getSilkLogConfig(sc datastore.Container) (executor.LogConfig, error) {
	var actualLogConfig executor.LogConfig
	logConfigStr, ok := sc.Metadata["log_config"].(string)
	if ok {
		err := json.Unmarshal([]byte(logConfigStr), &actualLogConfig)
		if err != nil {
			log.Printf("Datastore container: %s: error unmarshalling container log config from datastore: %s", sc.Handle, err.Error())
			return executor.LogConfig{}, err
		}
	}
	log.Printf("Datastore container: log config: %#v", actualLogConfig)
	return actualLogConfig, nil
}
