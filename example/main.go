package main

import (
	"log"

	"github.com/dmitrymomot/random"
)

func main() {
	str := random.String(16)
	log.Println(str)
}
