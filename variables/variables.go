package main

import "fmt"

func main() {

	//var объявляет 1 или более переменных.
	var text string = "Its variable which created by var"

	//Вы можете объявить несколько переменных одновременно
	var b, c int = 1, 2

	//Go выведет тип инициализированных переменных.
	var d = true

	//Переменные, объявленные без соответствующей инициализации,
	//имеют нулевое значение. Например, нулевое значение для int равно 0.
	var e int

	// e = 25

	//Синтаксис := является сокращением для объявления и инициализации
	// переменной, например. для var f string = "apple" в этом случае.
	f := "This is variable that created by :="

	fmt.Println(f)

	fmt.Println(e)

	fmt.Println(d)

	fmt.Println(b, c)

	fmt.Println(text)
}
