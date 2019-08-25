package main

import "fmt"

func main() {
	cities := map[string]int{
		"Bangkok": 10240,
		"Sarakam": 32000,
	}

	for key, value := range cities {
		fmt.Printf("%s, %d\n", key, value)
	}
}
