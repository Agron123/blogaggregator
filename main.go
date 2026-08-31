package main

import (
	"blogaggregator/internal/config"
	"blogaggregator/internal/database"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config %v", err)
	}

	dbURL := cfg.DBURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	dbQueries := database.New(db)

	s := state{
		cfg: &cfg,
		db:  dbQueries,
	}

	c := commands{
		handlers: map[string]func(*state, command) error{},
	}

	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerUsers)
	c.register("agg", handlerAgg)

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
