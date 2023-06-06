package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"test/service/generatedproto"
	"test/service/model"

	"google.golang.org/grpc"
)

func Start(service *model.Service, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	log.Println("gRPC start listen")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", service.Cfg.PortGrpc))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serv := grpc.NewServer()

	go func() {
		<-ctx.Done()
		log.Println("gRPC stop")
		serv.GracefulStop()
	}()
	generatedproto.RegisterSearchServer(serv, &generatedproto.ProtoService{})
	if err := serv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
