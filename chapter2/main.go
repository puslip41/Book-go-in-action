package main

import (
	"log"
	"os"

	_ "puslip41/go_in_action/chapter2/matchers"
	"puslip41/go_in_action/chapter2/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("Nikolas")
}
