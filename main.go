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
		log.Println("database not initialized")
	}
	controller.MyController()
}
