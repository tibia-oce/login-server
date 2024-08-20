package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/tibia-oce/login-server/src/api"
	"github.com/tibia-oce/login-server/src/configs"
	grpc_login_server "github.com/tibia-oce/login-server/src/grpc"
	"github.com/tibia-oce/login-server/src/logger"
	"github.com/tibia-oce/login-server/src/server"
)

var numberOfServers = 2
var initDelay = 200

func main() {

	logLevel := configs.GetLogLevel()
	logger.Init(logLevel)

	logger.Info("Welcome to OTBR Login Server")

	var wg sync.WaitGroup
	wg.Add(numberOfServers)

	gConfigs, err := configs.GetGlobalConfigs()
	if err != nil {
		logger.Panic(err)
	}

	go startServer(&wg, gConfigs, grpc_login_server.Initialize(gConfigs))
	go startServer(&wg, gConfigs, api.Initialize(gConfigs))

	time.Sleep(time.Duration(initDelay) * time.Millisecond)
	gConfigs.Display()

	// wait until WaitGroup is done
	wg.Wait()
}

func startServer(
	wg *sync.WaitGroup,
	gConfigs configs.GlobalConfigs,
	server server.ServerInterface,
) {
	if server == nil {
		logger.Error(fmt.Errorf("server is nil"))
		wg.Done()
		return
	}

	logger.Info(fmt.Sprintf("Starting %s server...", server.GetName()))
	err := server.Run(gConfigs)
	if err != nil {
		logger.Error(fmt.Errorf("server %s encountered an error: %v", server.GetName(), err))
	} else {
		logger.Info(fmt.Sprintf("Server %s stopped gracefully", server.GetName()))
	}
	wg.Done()
}
