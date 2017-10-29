package main

type Collection struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Photo []byte `json:"photo"`
}

type Collections []Collection