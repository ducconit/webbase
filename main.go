package main

import (
	"codebase/cmd"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	if err := cmd.Run(os.Args[1:]); err != nil {
		log.Fatalln(err)
	}
}
