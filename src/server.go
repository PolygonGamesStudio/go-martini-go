package main

import (
	"github.com/codegangsta/martini"
)

func main() {
	m := martini.Classic()

	m.Get("/place/:id", getPlaceDetail)
	m.Get("/favorites", getFavoritesList)
	m.Get("/route/:id", getRouteDetail)
	m.Get("/history/", getHistoryList)

	m.Post("/ratio/:id", postRatioDetail)

	m.Delete("/user/:id", deleteUser)

	m.Put("/ratio/:id", putRation)
	m.Put("/user/:id", putUser)

	m.Get("/logout", getLogout)
	m.Post("/login", getLogin)

	m.Run()
}
