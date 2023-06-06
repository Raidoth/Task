package service

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"test/service/config"
	"test/service/database"
	"test/service/grpc"
	"test/service/model"
)

func Start() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	service := model.Service{
		Cfg: config.InitService("service/resources/config.json"),
	}
	database.ConnectionDb(service) //connect database
	wg.Add(1)
	go grpc.Start(&service, &wg, ctx) //start grpc

	<-quit //wait graceful shutdown
	cancel()
	wg.Wait()
	database.DisconnectionDB()
}
