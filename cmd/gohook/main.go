package main

import (
	"gohook/internal/config"
	"gohook/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.Init()

	go server.Start()
	log.Println("Server started, addr", config.Get().Address)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
