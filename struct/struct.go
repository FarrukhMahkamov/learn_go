package main

import "fmt"

type Person struct {
	name     string
	username string
	age      int
}

type secretAgent struct {
	Person
	ltk bool
}

func main() {
	//1
	p := Person{
		"Farruh",
		"Rlly",
		22,
	}

	//2
	s := Person{
		name:     "Farruh",
		username: "Khanzo",
		age:      22,
	}

	//3 Omitted fields will be zero-valued.
	z := Person{
		name:     "Farruh",
		username: "Khanzo",
	}

	fmt.Println(p, s, z)
	fmt.Println(z.username)

	//Embeded struct

	secretAgent := secretAgent{
		Person: Person{
			name:     "Farruh",
			username: "Khanzo",
			age:      22,
		},
		ltk: true,
	}

	fmt.Println(secretAgent)
	fmt.Println(secretAgent.Person.username)

}
