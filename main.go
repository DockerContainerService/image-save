package main

import (
	"github.com/DockerContainerService/image-save/cmd"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"strconv"
)

var logLevel string

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&prefixed.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})
	if logLevel == "" {
		logLevel = "5"
	}
	log, err := strconv.ParseUint(logLevel, 10, 32)
	if err != nil {
		logrus.Fatalf("Log level error: %+v", err)
	}
	logrus.SetLevel(logrus.Level(log))
}

func main() {
	cmd.Execute()
}
