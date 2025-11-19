package main

import (
	"RedRock/core/runc"
	"RedRock/core/utils/banner"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	banner.GetBanner().ShowBanner()

	runc.GoWork()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server ShutDown, Bye!")
}
