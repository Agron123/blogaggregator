package main

import (
	"blogaggregator/internal/config"
	"log"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config %v", err)
	}

	s := state{
		cfg: &cfg,
	}

	c := commands{
		handlers: map[string]func(*state, command) error{},
	}

	c.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments")
	}

	cName := os.Args[1]
	cArgs := os.Args[2:]

	cmd := command{
		name: cName,
		args: cArgs,
	}

	err = c.run(&s, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
