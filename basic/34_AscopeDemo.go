package main

import (
	"log"
	"os"
)

func main() {

}

var cwd string

func init() {
	cwd, err := os.Getwd() // NOTE: wrong!
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}
