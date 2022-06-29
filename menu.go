package main

import "fmt"

var Menu = map[string]float64{
	"tandoori": 8.99,
	"paneer":   9.55,
	"egg-roll": 7.50,
	"rice":     4.95,
	"roti":     4.50,
}

func displayMenu() {
	fmt.Println("Menu")
	for k, v := range Menu {
		fmt.Printf("%-25v...%v\n", k, v)
	}
}
