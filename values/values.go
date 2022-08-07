package main

import "fmt"

// Go имеет различные типы значений, включая строки,
// целые числа, числа с плавающей запятой, логические
// значения и т. д. Вот несколько основных примеров. z
func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
