package main

import (
	"fmt"

	"github.com/10Narratives/ready-to-do/server/internal/app"
	"github.com/10Narratives/ready-to-do/server/internal/config"
)

func main() {
	_, err := app.New(config.MustLoad())
	if err != nil {
		panic(fmt.Sprintf("cannot initialize server application: %s", err.Error()))
	}

}
