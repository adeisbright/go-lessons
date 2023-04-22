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
		"Lekki":    {"Lekki", "Lagos"},
		"Choba":    {"Choba", "Rivers"},
		"Uyo":      {"Uyo", "Akwa-Ibom"},
		"Oraifite": {"Oraifite", "Anambra"},
		"Akure":    {"Akure", "Ondo"},
		"Akenfa":   {"Akenfa", "Bayelsa"},
		"Ekosodin": {"Ekosodin", "Edo"},
	}

	var example = map[string]int{
		"a": 24,
	}

	fmt.Println(places)
	fmt.Println(places["Akenfa"])
	fmt.Println(example["a"])
}
