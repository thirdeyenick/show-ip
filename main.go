package main

import (
	"log"
	"time"
)

func main() {
	log.Printf("failing server in 5 seconds...")
	time.Sleep(5 * time.Second)
	log.Fatal("failed...")
}
