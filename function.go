package main

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	a, b := swap(42, 13)
	println(a, b)
}
