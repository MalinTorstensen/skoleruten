package main

import (
	"fmt"
	"log"

	skoleruten "skoleruten"
)

func main() {
	fridager, err := skoleruten.DefaultFridagerForYear(2026)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Loaded %d holidays for 2026\n", len(fridager))
}
