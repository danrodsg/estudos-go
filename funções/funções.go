package main

import "fmt"

func main() {

	fmt.Println(soma(42, 13))
	fmt.Println(subtração(42, 13))
	fmt.Println(multiplicação(5, -1))
	fmt.Println(divisão(42, 13))
	fmt.Println("Dan Rodrigues")

}

func nome(Dan, Rodrigues string) string {
	return "nome"
}

func divisão(x int, y int) int {
	return x / y
}

func multiplicação(x int, y int) int {
	return x * y
}

func subtração(x int, y int) int {
	return x - y
}

func soma(x int, y int) int {
	return x + y

}