package main

import (
	"net"
	"os"
	"os/signal"

	"github.com/MiteshSharma/project/cmd/server"

	"github.com/MiteshSharma/project/core/config"
	"github.com/MiteshSharma/project/core/logger"
)

var version = "1.0.0"
var commit = ""
var branch = "master"
var startTime = ""
var buildNo = "0"

func main() {
	setting := config.NewSetting(buildNo, version, commit, branch, startTime)
	server := server.NewServer(setting)
	defer server.StopServer()

	server.StartServer()

	err := sendSystemdNotification()

	if err != nil {
		server.ServerParam.Logger.Debug("Systemd notification error.", logger.Error(err))
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	server.ServerParam.Logger.Debug("Stopped server.")

	server.StopServer()

	os.Exit(0)
}

func sendSystemdNotification() error {
	notifySocket := os.Getenv("NOTIFY_SOCKET")
	if notifySocket != "" {
		state := "READY=1"
		socketAddr := &net.UnixAddr{
			Name: notifySocket,
			Net:  "unixgram",
		}
		conn, err := net.DialUnix(socketAddr.Net, nil, socketAddr)
		if err != nil {
			return err
		}
		defer conn.Close()
		_, err = conn.Write([]byte(state))
		return err
	}
	return nil
}
