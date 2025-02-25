package main

import (
	"fmt"
	"maycms/domain"
)

var c = domain.Content{}

func main() {
	c.ContentText = "Lorem ipsum lorem ipsum lorem ipsum"
	c.Title = "Lorem ipsum"
	fmt.Println(c.ContentText)
	fmt.Println(c.Title)

}
