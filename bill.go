package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	total float64
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		total: 0.0,
		items: make(map[string]float64),
		tip:   0.0,
	}
	return b
}

func (b *bill) addItem(i string) {
	b.total += Menu[i]
	b.items[i] += Menu[i]
}

func (b *bill) addItems(i []string) {
	for _, v := range i {
		(*b).total += Menu[v]
		(*b).items[v] = Menu[v]
	}
}

func (b *bill) format() string {
	s := "Name: " + b.name + "\n"
	for k, v := range b.items {
		s += fmt.Sprintf("%-25v...%v\n", k+":", v)
	}
	s += fmt.Sprintf("%-25v...%0.2f\n", "total:", b.total) //b is auto-dereferenced as it is a struct
	return s
}

func (b *bill) save() {
	d := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", d, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved successfully")
}
