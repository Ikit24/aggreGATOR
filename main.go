package main

import (
	"fmt"
	"log"
	"github.com/Ikit24/aggreGATOR/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfgVal, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	s := &state{cfg: &cfgVal}

	_ = s
}
