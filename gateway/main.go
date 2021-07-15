package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Tsumida/my-cmdb/utils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	host_svc "my-cmdb/api/host_server"
)

func main() {
	svcName := "gateway"
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	conn, err := grpc.DialContext(
		ctx,
		"localhost:8880",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("failed to build grpc conn")
	}

	gwMux := runtime.NewServeMux()

	if err := host_svc.RegisterHostManagementHandler(ctx, gwMux, conn); err != nil {
		log.Fatalln("failed to register host-server handler")
	}

	go http.ListenAndServe("localhost:8881", gwMux)

	utils.SingalHandler(svcName)
}
