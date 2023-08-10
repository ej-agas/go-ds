package main

import "fmt"

func main() {
	table := NewHashTable()
	list := []string{
		"John",
		"Erik",
		"Joe",
		"Jane",
		"Laika",
		"Momo",
	}

	for _, v := range list {
		table.Insert(v)
	}

	fmt.Println("John:", table.Search("John"))
	table.Delete("John")
	fmt.Println("John:", table.Search("John"))
	fmt.Println("Momo:", table.Search("Momo"))
}
