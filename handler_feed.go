package main

import (
	"fmt"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("addfeed requires 2 arguments")
	}
	
	return nil
}
