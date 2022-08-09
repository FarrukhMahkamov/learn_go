package main

import "fmt"

func main() {
	// var a [5]int
	// fmt.Println("emp:", a)

	// a[4] = 100
	// fmt.Println("Noemp:", a)

	// fmt.Println("len:", len(a)) // length of a  equal to 5

	// b := [5]int{1, 2, 3, 4, 5}
	// fmt.Println("dcl:", b)

	// for i := 0; i < len(b); i++ {
	// 	if b[i]%2 == 0 {
	// 		fmt.Println(b[i])
	// 	}
	// }

	// var bornYear int = 2000
	// var now int = 2022

	// for bornYear <= now {
	// 	fmt.Println(bornYear)
	// 	bornYear++
	// }

	// for {
	// 	if bornYear > now {
	// 		break
	// 	}
	// 	fmt.Println(bornYear)
	// 	bornYear++
	// }

	for i := 1; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Println(i)
		}
	}
}
