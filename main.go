package main

import (
	"blogaggregator/internal/config"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config %v", err)
	}

	err = cfg.SetUser("Dovydas")
	if err != nil {
		log.Fatalf("Error setting username %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading the file %v", err)
	}

	fmt.Printf("Read config again: %+v\n", cfg)

}
