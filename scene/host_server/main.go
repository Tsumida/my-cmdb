package main

import (
	"context"
	"fmt"
	"my-cmdb/scene/host_server/service"
	"os"

	"my-cmdb/utils"
)

var (
	debug_mode = true
)

func main() {
	// grpc server
	svcName := "host-manager"
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	service.RunService(ctx)

	if debug_mode {
		pid := os.Getpid()
		fmt.Printf("pid=%d\n", pid)
	}

	utils.SingalHandler(svcName)

}
