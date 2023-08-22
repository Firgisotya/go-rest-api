package main

import (
	"os"
	"github.com/Firgisotya/go-rest-api/config"
	"github.com/Firgisotya/go-rest-api/config/command"
)

func main() {
	functionRun := os.Args[1]

	switch functionRun {
	case "db:migrate":
		command.Migrate()
	case "db:seed":
		command.Seed(config.DB)
	}

}
