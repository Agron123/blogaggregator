package main

import (
	"blogaggregator/internal/config"
	"blogaggregator/internal/database"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	function, ok := c.handlers[cmd.name]
	if !ok {
		return errors.New("command does not exist")
	}

	err := function(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f

}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("the login handler expects a single argument, the username")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println("username has been set")

	return nil

}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("the register handler expects a single argument, the name")
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}
	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("Created user %v", user)

	return nil

}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetDB(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("database reset successfully")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}

	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	rss, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Print(rss)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return errors.New("the add feed handler expects two arguments, the name and url")
	}

	currentUser := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return err
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf("Created feed: %v", feed)

	return nil

}
