package main

import (
	"RedRock/core/runc"
	"RedRock/core/utils/banner"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	banner.GetBanner().ShowBanner()

	runc.GoWork()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server ShutDown, Bye!")
}
