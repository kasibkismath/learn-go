package main

import (
	"fmt"

	"log"

	"kasib/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// fmt.Println(quote.Glass())
	message, err := greetings.Hellos([]string{"Kasib Kismath", "Nazeeha Nimnaz", "Rawdah Kasib"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
