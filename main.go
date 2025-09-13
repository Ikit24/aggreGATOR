package main

import (
	"fmt"
	"log"
	"github.com/Ikit24/aggreGATOR/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	if err := cfg.SetUser("Attila"); err != nil {
		log.Fatal(err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.DBURL)
}
