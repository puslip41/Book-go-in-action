package main

import (
	"log"
	"os"

	_ "puslip41/book-go-in-action/chapter2/matchers"
	"puslip41/book-go-in-action/chapter2/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("Sherlock Holmes")
}
