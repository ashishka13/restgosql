package main

import (
	"log"
	"restgo/controller"
	"restgo/utils"
)

func main() {
	log.Println("hello world")
	err := utils.InitializeDatabase()
	if err != nil {
		log.Panic("database not initialized")
	}
	controller.MyController()
}
