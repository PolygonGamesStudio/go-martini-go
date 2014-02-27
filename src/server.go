package main

import (
	"github.com/codegangsta/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world - GET \n"
	})

	m.Post("/", func() string {
		return "Hello world - POST \n"
		// create something
	})

	m.Put("/", func() string {
		return "Hello world - PUT \n"
		// replace something
	})

	m.Delete("/", func() string {
		return "Hello world - DELETE \n"
		// destroy something
	})

	m.Run()
}
