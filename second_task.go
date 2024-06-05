package main

func square(n int) int {
	return n * n
}
func process(num1, num2, num3, num4 int, square func(int) int) {
	println(square(num1))
	println(square(num2))
	println(square(num3))
	println(square(num4))
}

func main() {
	num1 := 1
	num2 := 2
	num3 := 3
	num4 := 4
	process(num1, num2, num3, num4, square)
	// output
	// 1
	// 4
	// 9
	// 16
}
