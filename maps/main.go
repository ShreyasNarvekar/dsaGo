package main

import "fmt"

func main() {
	// colors := map[string]string{ //creating map with some data.
	// 	"red":   "#ff0000",			//red is the key and #ff0000 is the value of the key.
	// 	"green": "#4bf745",
	// }

	// var colors map[string]string	//creating empty map.

	// colors := make(map[string]string) // creating empty map using make(map[type]type).
	// colors["shreyas"] = "narvekar"    //putting values
	// fmt.Println(colors)
	// delete(colors, "shreyas") //deleting values by passing map,key
	// fmt.Println(colors)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}

}
