package main

type Product struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Specs map[string]string `json:"specs"`
	Price int64 `json:"price"`
	Feedback []string `json:"feedback"`
	Images [][]byte `json:"images"`
	Collections []string `json:"collections"`
	Tags []string `json:"tags"`
}

type Products []Product