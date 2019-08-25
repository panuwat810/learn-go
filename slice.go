package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(p)
	fmt.Println(p[0:3])
	fmt.Println(p[3:5])
	fmt.Println(p[4:])
}
