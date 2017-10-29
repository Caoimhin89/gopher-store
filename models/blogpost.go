package main

import "time"

type Blogpost struct {
	Id string
	Title string
	Subtitle string
	Excerpt string
	PubDate time.Time
	Author string
	Body string
	CoverPhoto []byte
	ContentPhotos [][]byte
}

type Blog []Blogpost