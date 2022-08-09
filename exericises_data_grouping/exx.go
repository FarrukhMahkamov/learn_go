package main

import "fmt"

func main() {
	// firstEx()
	// secondEx()
	// thirdEx()
	// fourthEx()
	// fifthEx()
	// sixthEx()
	seventhEx()
}

func firstEx() {
	arr := [5]string{
		"First",
		"Second",
		"Third",
		"Fourth",
		"Fifth",
	}

	fmt.Println(arr)

	for _, v := range arr {
		fmt.Println(v)
	}

	fmt.Printf("%T\n", arr)
}

func secondEx() {
	slice := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	for i, v := range slice {
		fmt.Println("Index", i, " - ", v)
	}

	fmt.Printf("%T\n", slice)

	fmt.Println(slice[2:5])
}

func thirdEx() {
	//1
	x := []int{
		42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	}
	//2
	x = append(x, 52)
	//3
	x = append(x, 53, 54, 55)
	//4
	y := []int{
		55, 56, 57, 58, 59,
	}

	x = append(x, y...)
	fmt.Println(x)
}

func fourthEx() {
	x := []int{
		42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	}

	x = append(x[:3], x[6:]...)

	fmt.Println(x)
}

func fifthEx() {
	states := make([]string, 55, 55)
	states = []string{
		"Alaska",
		"Alabama",
		"Arkansas",
		"American Samoa",
		"Arizona",
		"California",
		"Colorado",
		"Connecticut",
		"District of Columbia",
		"Delaware",
		"Florida",
		"Georgia",
		"Guam",
		"Hawaii",
		"Iowa",
		"Idaho",
		"Illinois",
		"Indiana",
		"Kansas",
		"Kentucky",
		"Louisiana",
		"Massachusetts",
		"Maryland",
		"Maine",
		"Michigan",
		"Minnesota",
		"Missouri",
		"Mississippi",
		"Montana",
		"North Carolina",
		" North Dakota",
		"Nebraska",
		"New Hampshire",
		"New Jersey",
		"New Mexico",
		"Nevada",
		"New York",
		"Ohio",
		"Oklahoma",
		"Oregon",
		"Pennsylvania",
		"Puerto Rico",
		"Rhode Island",
		"South Carolina",
		"South Dakota",
		"Tennessee",
		"Texas",
		"Utah",
		"Virginia",
		"Virgin Islands",
		"Vermont",
		"Washington",
		"Wisconsin",
		"West Virginia",
		"Wyoming",
	}
	//1
	fmt.Println(states)
	fmt.Println(len(states))
	fmt.Println(cap(states))

	//2
	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}
}

func sixthEx() {
	artists := []string{
		"James",
		"Michael",
		"Anton",
		"Justin",
		"Eminem",
	}

	justPeople := []string{
		"Farruh",
		"Umid",
		"Ozod",
		"Nigora",
		"Gulchexra",
	}
	//2
	mix := [][]string{artists, justPeople}
	//3
	fmt.Println(mix)
	//4

	//5
	for i, mx := range mix {
		fmt.Println("Record ", i)
		for j, val := range mx {
			fmt.Printf("\t Index position: %v \t value: \t %v \n", j, val)
		}
	}
}

func seventhEx() {
	people := map[string][]string{
		"James Bond": []string{
			"Shaken, not stirred",
			"Martinis",
			"Womans",
		},
		"Just Me": []string{
			"Golang",
			"Music",
			"Games",
			"Womens",
		},
		"Emine": []string{
			"Music",
			"Rap",
			"Womans",
			"Hip Hop",
		},
	}

	for i, firs := range people {
		fmt.Println("THis is record for: ", i)
		for j, sec := range firs {
			fmt.Println("\t", j, sec)
		}
	}
}
