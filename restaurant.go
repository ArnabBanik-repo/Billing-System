package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(p string, r *bufio.Reader) (string, error) {
	fmt.Print(p)
	i, e := r.ReadString('\n')
	return strings.TrimSpace(i), e
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Enter bill name: ", reader)
	b := newBill(name)
	return b
}

func options(bill *bill) {
	r := bufio.NewReader(os.Stdin)
	ch, _ := getInput("Enter your choice (a - add new item, s - save, e - abort): ", r)
	switch ch {
	case "a":
		displayMenu()
		itm, _ := getInput("Item name: ", r)
		if Menu[itm] == 0 {
			fmt.Println("Please choose an item from the menu")
			options(bill)
		} else {
			bill.addItem(itm)
			options(bill)
		}
	case "s":
		bill.save()
		return
	case "e":
		return
	default:
		fmt.Println("Invalid choice")
		options(bill)
	}

}

func main() {
	bill := createBill()
	options(&bill)
}
