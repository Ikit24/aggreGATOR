package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/Ikit24/aggreGATOR/internal/config"
    "github.com/Ikit24/aggreGATOR/internal/database"
    "github.com/google/uuid"
    _ "github.com/lib/pq"
)

type state struct {
	db	*database.Queries
	cfg	*config.Config
}

type command struct {
	name string
	args []string
}

type handlerFunc func(*state, command) error

type commands struct {
	m map[string]handlerFunc
}

func (c *commands) register(name string, f handlerFunc) {
	if c.m == nil {
		c.m = make(map[string]handlerFunc)
	}
	c.m[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.m[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return handler(s, cmd)
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("login expects a single argument, the username.")
	}
	username := cmd.args[0]
	
	_, err := s.db.GetUser(context.Background(),username)
	if err != nil {
		return fmt.Errorf("user does not exist: %s", username)
	}
	if err := s.cfg.SetUser(username); err != nil {
		return err
	}
	fmt.Println("user set to:", username)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("register expects a single argument, the username.")
	}
	username := cmd.args[0]
	
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	_, err := s.db.CreateUser(context.Background(),params)
	if err != nil {
		return fmt.Errorf("user already exists: %s", username)
	}
	
	if err := s.cfg.SetUser(username); err != nil {
		return err
	}

	fmt.Println("user created:", username)
	return nil
}

func main() {
	cfgVal, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	dbURL := cfgVal.DBURL

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := &state{
		cfg: &cfgVal,
		db:	 database.New(db),
	}

	cmds := &commands{m: make(map[string]handlerFunc)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "not enough arguments")
		os.Exit(1)
	}
	name := os.Args[1]
	args := os.Args[2:]

	if err := cmds.run(s, command{name: name, args: args}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
