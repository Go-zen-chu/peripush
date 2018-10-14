package main

import (
	"github.com/go-zen-chu/peripush/pkg/application/command"
	"log"
)

func main() {
	cmd := command.NewCommand()
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}