package main

import "fmt"

func mapper() {
	fmt.Println("Working with Map")

	type Place struct {
		name  string
		state string
	}
	/*
		To create a map
		map<keyType><ValueType>
	*/
	var places = map[string]Place{
		"Lekki":    Place{"Lekki", "Lagos"},
		"Choba":    Place{"Choba", "Rivers"},
		"Uyo":      Place{"Uyo", "Akwa-Ibom"},
		"Oraifite": Place{"Oraifite", "Anambra"},
		"Akure":    Place{"Akure", "Ondo"},
		"Akenfa":   Place{"Akenfa", "Bayelsa"},
		"Ekosodin": Place{"Ekosodin", "Edo"},
	}

	var example = map[string]int{
		"a": 24,
	}

	fmt.Println(places)
	fmt.Println(places["Akenfa"])
	fmt.Println(example["a"])
}
