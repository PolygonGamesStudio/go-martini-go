package main

import (
	"github.com/codegangsta/martini"
)

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world - GET"
	})

	m.Post("/", func() string {
		return "Hello world - POST"
		// create something
	})

	m.Put("/", func() string {
		return "Hello world - PUT"
		// replace something
	})

	m.Delete("/", func() string {
		return "Hello world - DELETE"
		// destroy something
	})

	m.Run()
}
