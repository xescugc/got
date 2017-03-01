package main

import (
	"log"

	"./commands"
)

var ()

func init() {
}

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
