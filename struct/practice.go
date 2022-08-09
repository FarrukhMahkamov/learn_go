package main

import "fmt"

type Person struct {
	firstName     string
	lastName      string
	favoriteGames []string
}

func main() {
	firstPerson := Person{
		firstName: "Farruh",
		lastName:  "Mahkamov",
		favoriteGames: []string{
			"Mount and blade",
			"Mobile Legeneds Bang Bang",
			"Counter Strike Global Offensive",
		},
	}

	secondPerson := Person{
		firstName: "Jonh",
		lastName:  "Wick",
		favoriteGames: []string{
			"Forza Horizon V",
			"GTA V",
			"Counter Strike Global Offensive",
		},
	}

	p1map := map[string]Person{
		firstPerson.firstName:  firstPerson,
		secondPerson.firstName: secondPerson,
	}

	for i, v := range p1map {
		fmt.Println(i)
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for j, b := range v.favoriteGames {
			fmt.Println(j, b)
		}
	}

	fmt.Println(p1map)

	// fmt.Println(firstPerson, secondPerson)

	fmt.Println(firstPerson.firstName, firstPerson.lastName)
	for i, v := range firstPerson.favoriteGames {
		fmt.Println(i, v)
	}

	fmt.Println(secondPerson.firstName, secondPerson.lastName)
	for i, v := range secondPerson.favoriteGames {
		fmt.Println(i, v)
	}
}
