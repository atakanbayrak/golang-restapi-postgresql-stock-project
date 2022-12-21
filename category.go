package main

var categories = []Category{}

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
