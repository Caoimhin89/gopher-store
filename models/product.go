package main

type Product struct {
	Id string
	Name string
	Description string
	Specs map[string]string
	Price int64
	Feedback []string
	Images [][]byte
	Collections []string
	Tags []string
}

type Products []Product