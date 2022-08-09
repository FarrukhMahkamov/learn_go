package main

import "fmt"

func main() {
	maps := map[string]int{
		"James":    21,
		"Farruh":   33,
		"Nat":      45,
		"Xurshida": 20,
		"Xushnoza": 20,
	}
	maps["Zilola"] = 33 // adding element to map

	delete(maps, "Xurshida") //removing element from map

	// If given index is not exists it returns 0
	// fmt.Println(maps["Nasd"])

	// if v, ok := maps["Zilola"]; ok {
	// 	fmt.Println("This person is exists", v)
	// } else {
	// 	fmt.Println("This person not exists yet")
	// }

	for i, r := range maps {
		fmt.Println(i, r)
	}
}
