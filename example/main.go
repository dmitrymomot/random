package main

import (
	"log"

	"github.com/dmitrymomot/random/v2"
)

func main() {
	str := random.String(16)
	log.Println(str)
}
