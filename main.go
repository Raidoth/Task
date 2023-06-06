package main

import (
	"log"
	"test/service"
)

func main() {
	log.Println("Start service")
	service.Start()
	log.Println("Stop service")

}
