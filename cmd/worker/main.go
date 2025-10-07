package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Println("worker tick...")
		time.Sleep(10 * time.Second)
	}
}
