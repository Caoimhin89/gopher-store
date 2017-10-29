package main

type Product struct {
	Name string
	Description string
	Specs map[string]string
	Price int64
	Feedback []string
	Images [][]byte
}

type Products []Product