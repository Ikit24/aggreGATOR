package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/Ikit24/aggreGATOR/internal/config"
    "github.com/Ikit24/aggreGATOR/internal/database"
    _ "github.com/lib/pq"
)

type state struct {
	db	*database.Queries
	cfg	*config.Config
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

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "not enough arguments")
		os.Exit(1)
	}
	name := os.Args[1]
	args := os.Args[2:]

	if err := cmds.run(s, command{Name: name, Args: args}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
