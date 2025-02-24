package main

import (
	"fmt"

	"github.com/mgarcia88/maycms/domain"
)

var c = domain.NewContent()

func main() {
	c.ContentText = "Lorem ipsum lorem ipsum lorem ipsum"
	c.Title = "Lorem ipsum"
	fmt.Println(c.ContentText)
	fmt.Println(c.Title)

}
