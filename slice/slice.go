package main

import "fmt"

func main() {
	// x := make([]int, 10, 100)

	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))
	// fmt.Println(x)

	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
	// fmt.Println(x[1:4])

	// x = append(x, 123, 23423, 5345, 546, 456, 45, 645, 645, 645, 645, 64, 56, 456, 45)
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }

	jb := []string{"James", "Nick", "Carter"}
	bm := []string{"Block", "Changer", "Johnatan"}

	xp := [][]string{jb, bm}

	fmt.Println(xp)
}
