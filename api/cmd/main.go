package main

import (
	"api/config"
	"api/db"
	"fmt"
	"log"
	"time"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Println("Error while reading a config")
		log.Fatal(err)
	}
	err = db.RunMigrations(cfg)
	if err != nil {
		log.Println("Error while doing migrations")
		log.Fatal(err)
	}
	fmt.Println("Success starting app...")
	time.Sleep(time.Minute)
}
