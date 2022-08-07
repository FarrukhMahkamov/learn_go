package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func main() {

	// fmt.Println("1.Необходимо вывести на экран числа от 1 до 5. На экране должно быть:")
	// fmt.Println("1 2 3 4 5")
	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// fmt.Println("2.Необходимо вывести на экран числа от 5 до 1. На экране должно быть:")
	// fmt.Println("5 4 3 2 1")
	// for i := 5; i >= 1; i-- {
	// 	fmt.Println(i)
	// }

	// fmt.Println("3.Необходимо вывести на экран таблицу умножения на 3:")

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("3 *", i, "=", 3*i)
	// }

	// fmt.Println("4.Напишите программу, где пользователь вводит любое целое положительное число. А программа суммирует все числа от 1 до введенного пользователем числа.")
	// var number int

	// fmt.Println("Вводите любое целое положительное число")
	// fmt.Scanln(&number)

	// result := 0
	// for i := 1; i <= number; i++ {
	// 	result = result + i
	// }
	// fmt.Println(result)

	// fmt.Println("Вывести все квадраты натуральных чисел, не превосходящие данного числа N.")
	// var number int

	// fmt.Println("Введите предел")
	// fmt.Scanln(&number)

	// for i := 1; i <= number; i++ {
	// 	if i*i <= number {
	// 		fmt.Println(i, i*i)
	// 	}
	// }

	// fmt.Println("Вывести все квадраты натуральных чисел, не превосходящие данного числа N.")
	// var firstNumber int
	// var secondNumber int

	// fmt.Println("Введите первое число")
	// fmt.Scanln(&firstNumber)

	// fmt.Println("Введите второе число")
	// fmt.Scanln(&secondNumber)

	// if firstNumber > secondNumber {
	// 	for i := firstNumber; i >= secondNumber; i-- {
	// 		fmt.Println(i, i*i*i)
	// 	}
	// } else {
	// 	for i := firstNumber; i <= secondNumber; i++ {
	// 		fmt.Println(i, i*i*i)
	// 	}
	// }

	// fmt.Println("Вычислить факториал числа, которое ввел пользователь.")
	// var number float32

	// fmt.Println("Введите число")
	// fmt.Scanln(&number)

	// result := 1
	// for i := 1; float32(i) <= number; i++ {
	// 	result = result * i
	// }

	// fmt.Println(result)

	fmt.Println("Вывести на экран ряд чисел Фибоначчи, состоящий из N элементов. Значение N вводится c клавиатуры.")
	var n int
	var z int

	fmt.Println("Введите число")

	fmt.Scanln(&n)
	CallClear()

	x := 0
	y := 1

	for i := 1; i <= n; i++ {
		z = x + y // 1
		x = y     // 1
		y = z     // 1
		fmt.Println(z)
	}
}
