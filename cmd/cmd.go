package main

import (
	"flag"
	"github.com/hazcod/intigriti-slack-announce/config"
	"github.com/hazcod/intigriti-slack-announce/findingchecker"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	defaultConfName = "isa.yaml"
	defaultLogLevel = logrus.InfoLevel
	defaultCheckInterval = 15

	ClientVersion = "intigriti-slack-announce/1.0"
)

func main() {
	confPath := flag.String("conf", defaultConfName, "The path to the configuration file.")
	logLevelStr := flag.String("loglevel", defaultLogLevel.String(), "The log level from debug (5) to error (1).")
	flag.Parse()

	if _, err := os.Stat(*confPath); os.IsNotExist(err) {
		log.Fatalf("configuration file not found: %s", *confPath)
	}

	logLevel, err := logrus.ParseLevel(*logLevelStr)
	if err != nil {
		log.Fatalf("invalid log level: %v", err)
	}
	logrus.SetLevel(logLevel)

	config, err := config.ParseConfig(*confPath)
	if err != nil {
		log.Fatal(err)
	}
	config.ConfigPath = *confPath

	if config.CheckInterval == 0 {
		config.CheckInterval = defaultCheckInterval
	}

	if err := findingchecker.RunChecker(config, ClientVersion); err != nil {
		log.Fatalf("could not run checker: %v", err)
	}

	logrus.Debug("waiting for SIGTERM")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	logrus.Info("shutting down")
}
