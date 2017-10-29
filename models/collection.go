package main

type Collection struct {
	Id string
	Name string
	Description string
	Photo []byte
}

type Collections []Collection