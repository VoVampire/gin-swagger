package main

import (
	. "gin-swagger/gin-router"
	"log"
)

func main() {
	err := GinEngine().Run()
	if err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}
