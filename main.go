package main

import (
	"log"

	"github.com/xescugc/got/commands"
)

func init() {
}

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
