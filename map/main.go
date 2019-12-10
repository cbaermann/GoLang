package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b745",
		"white": "#ffffff",
	}

	// delete(colors, "white")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	//for key, value in range of colors
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

//map vs struct
//map- all keys must be of same type, all values must be of same time
//map-all different keys are index'd, can iterate over key value pairs
//can not iterate over all key value pairs in struct
//map is a reference type, struct is a value type.
//general guideline, use map when using collection of closley related properties.
//map-do not need to know of a list of keys or field names at compiled time. can add and remove over time when you want.
