package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	programmingLanguage := make(map[string]string)
	programmingLanguage["JS"] = "JavaScript"
	programmingLanguage["PY"] = "Python"
	programmingLanguage["CPP"] = "C++"

	fmt.Println(programmingLanguage)
	fmt.Println("JS stands for", programmingLanguage["JS"])

	// Remove item based on Key
	delete(programmingLanguage, "JS")
	fmt.Println("Map after deletion", programmingLanguage)

	for key, value := range programmingLanguage {
		fmt.Printf("For key %v value is %v\n", key, value)
	}
}
