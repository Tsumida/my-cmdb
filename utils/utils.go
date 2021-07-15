package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Block until receive some signal.
func SingalHandler(svcName string) {
	fmt.Printf("service %s up\n", svcName)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-c
	fmt.Printf("\nservice %s down\n", svcName)
}
