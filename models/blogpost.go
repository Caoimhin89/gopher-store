package main

import "time"

type Blogpost struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Subtitle string `json:"subtitle"`
	Excerpt string `json:"excerpt"`
	PubDate time.Time `json:"pubDate"`
	Author string `json:"author"`
	Body string `json:"body"`
	CoverPhoto []byte `json:"coverPhoto"`
	ContentPhotos [][]byte `json:"contentPhotos"`
}

type Blog []Blogpost