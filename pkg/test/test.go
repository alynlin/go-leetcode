package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	signalChan := make(chan os.Signal)
	list := []int{1, 2, 3}

	for _, v := range list {
		v := v
		go func() {
			fmt.Println(v)
		}()
	}
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	log.Println("Shutdown signal received, exiting...")
	close(signalChan)
}
