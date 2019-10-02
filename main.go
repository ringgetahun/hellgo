package main

import (
	"fmt"
	"os"
)

func translate(input string) string {
	var transilations = make(map[string]string)
	transilations["en"] = "Hello"
	transilations["de"] = "Guten Tag"
	transilations["fr"] = "Bonjure"
	return transilations[input]
}
func main() {
	var locale string
	if len(os.Args) == 1 {
		fmt.Println("Enter language code:")
		fmt.Scanf("%s", &locale)
	} else {
		locale = os.Args[1]
	}
	output := translate(locale)
	if output == " " {
		output = "Yo"
	}
	/*var transilations=make(map[string]string)
	transilations["en"]="Hello"
	transilations["de"]="Guten Tag"
	transilations["fr"]="Bonjure"
	return translations[input] */

	/*var languages = [4]string{"en", "es", "de", "fr"}
	locale = languages[1]


	fmt.Scanf("Enter language code:")
	fmt.Scanf("%s", &locale)
	/*
		if locale == "en" {
			greeting = "Hello"
		} else if locale == "es" {
			greeting = "Hola"
		} else if locale == "de" {
			greeting = "Guten Tag"
		} else {
			greeting = "Yo"
		}


	switch locale {
	case "en":
		greeting = "Hello"
	case "es":
		greeting = "Hola"
	case "de":
		greeting = "Guten Tag"
	case "fr":
		greeting = "Bonjour"
	default:
		greeting = "Yo"
	}

	fmt.Printf(greeting + ", Go!\n")
	*/
	fmt.Println(output, "Go!")
}
