package main

import (
	"flag"
	"fmt"
	"iptables-logger/config"
	"iptables-logger/merger"
	"iptables-logger/parser"
	"iptables-logger/repository"
	"iptables-logger/runner"
	"iptables-logger/taillogger"
	"lib/common"
	"lib/datastore"
	"lib/serial"
	"log"
	"os"
	"sync"
	"time"

	"github.com/cloudfoundry/dropsonde"
	"github.com/hpcloud/tail"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/grouper"
	"github.com/tedsuo/ifrit/sigmon"

	"iptables-logger/rotatablesink"

	"io"

	"code.cloudfoundry.org/cf-networking-helpers/metrics"
	"code.cloudfoundry.org/filelock"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagerflags"
)

const (
	dropsondeOrigin = "iptables-logger"
	emitInterval    = 30 * time.Second
	jobPrefix       = "iptables-logger"
	logPrefix       = "cfnetworking"
)

func main() {
	configFilePath := flag.String("config-file", "", "path to config file")
	flag.Parse()
	conf, err := config.New(*configFilePath)
	if err != nil {
		log.Fatalf("%s.iptables-logger: reading config: %s", logPrefix, err)
	}

	logger, sink := lagerflags.NewFromConfig(fmt.Sprintf("%s.%s", logPrefix, jobPrefix), common.GetLagerConfig())
	sink.SetMinLevel(lager.DEBUG)

	logger.Info("starting")

	tailConfig := tail.Config{
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: io.SeekEnd,
		},
		MustExist: true,
		Follow:    true,
		Poll:      true,
		ReOpen:    true,
	}

	if conf.LogTimestampFormat == "rfc3339" {
		tailConfig.Logger = taillogger.Shim{Logger: logger}
	}

	t, err := tail.TailFile(conf.KernelLogFile, tailConfig)
	if err != nil {
		logger.Fatal("tail-input", err)
	}

	logger.Info("started tailing file")

	kernelLogParser := &parser.KernelLogParser{}

	store := &datastore.Store{
		Serializer: &serial.Serial{},
		Locker: &filelock.Locker{
			FileLocker: filelock.NewLocker(conf.ContainerMetadataFile + "_lock"),
			Mutex:      new(sync.Mutex),
		},
		DataFilePath:    conf.ContainerMetadataFile,
		VersionFilePath: conf.ContainerMetadataFile + "_version",
		LockedFilePath:  conf.ContainerMetadataFile + "_lock",
		CacheMutex:      new(sync.RWMutex),
	}
	containerRepo := &repository.ContainerRepo{
		Store: store,
	}
	logMerger := &merger.Merger{
		ContainerRepo: containerRepo,
		HostIp:        conf.HostIp,
		HostGuid:      conf.HostGuid,
	}
	iptablesLogger := lager.NewLogger(fmt.Sprintf("%s.iptables", logPrefix))
	outputLogFile, err := os.OpenFile(conf.OutputLogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatal("open-output-log-file", err)
	}

	iptablesSink, err := rotatablesink.NewRotatableSink(
		outputLogFile.Name(),
		lager.DEBUG,
		rotatablesink.DefaultFileWriterFunc(rotatablesink.DefaultFileWriter),
		rotatablesink.DefaultDestinationFileInfo{},
		logger,
		conf.LogTimestampFormat == "rfc3339",
	)

	if err != nil {
		logger.Fatal("rotatable-sink", err)
	}
	iptablesLogger.RegisterSink(iptablesSink)

	err = dropsonde.Initialize(conf.MetronAddress, dropsondeOrigin)
	if err != nil {
		log.Fatalf("%s: initializing dropsonde: %s", logPrefix, err)
	}

	uptimeSource := metrics.NewUptimeSource()
	metricsEmitter := metrics.NewMetricsEmitter(logger, emitInterval, uptimeSource)

	runner := &runner.Runner{
		Lines:          t.Lines,
		Parser:         kernelLogParser,
		Logger:         logger,
		Merger:         logMerger,
		IPTablesLogger: iptablesLogger,
	}

	members := grouper.Members{
		{"metrics_emitter", metricsEmitter},
		{"iptables_runner", runner},
	}

	monitor := ifrit.Invoke(sigmon.New(grouper.NewOrdered(os.Interrupt, members)))
	<-monitor.Wait()
}
