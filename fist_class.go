package main

import "fmt"

func main() {
	showMe := func(a, b int) {
		fmt.Printf("sum := %d", a+b)
	}

	showMe(1, 2)
}
