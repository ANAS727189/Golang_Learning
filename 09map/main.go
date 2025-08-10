package main

import "fmt"

func main() {
	lang := make(map[string]string)
	lang["en"] = "English"
	lang["fr"] = "French"
	lang["es"] = "Spanish"
	lang["de"] = "German"

	fmt.Println("Languages:", lang)
	fmt.Println("EN stands for:", lang["en"])

	delete(lang, "fr")

	for key, val := range lang {
		fmt.Printf("Key: %s, Value: %s\n", key, val)
	}
}
