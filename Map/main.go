package main

import "fmt"

func main() {
	// var colors map[string]string

	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	colors["red"] = "#adsfjk"
	colors["blue"] = "#weiurn"
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, "is", hex)
	}
}
