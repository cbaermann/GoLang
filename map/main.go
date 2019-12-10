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
