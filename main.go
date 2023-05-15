package main

import (
	"os"
	"os/signal"
	"syscall"
	"tradex.com/server_temp/zeebe"
)

func main() {
	shutdownBarrier := make(chan bool, 1)
	SetupShutdownBarrier(shutdownBarrier)

	client := zeebe.InitZeebeClient()
	defer zeebe.MustCloseClient(client)
	zeebe.MustStartWorker(client)

	<-shutdownBarrier
}

func SetupShutdownBarrier(done chan bool) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()
}
