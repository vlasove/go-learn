package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("empty args")
	}
	log.Println("arg is:", os.Args[1])
}
