package main

import "fmt"

func main() {

	fmt.Println("1.Необходимо вывести на экран числа от 1 до 5. На экране должно быть:")
	fmt.Println("1 2 3 4 5")
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("2.Необходимо вывести на экран числа от 5 до 1. На экране должно быть:")
	fmt.Println("5 4 3 2 1")
	for i := 5; i >= 1; i-- {
		fmt.Println(i)
	}

	fmt.Println("3.Необходимо вывести на экран таблицу умножения на 3:")

	for i := 1; i <= 10; i++ {
		fmt.Println("3 *", i, "=", 3*i)
	}
}