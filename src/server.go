package main

import (
	"github.com/codegangsta/martini"
)

func main() {
	m := martini.Classic()

	m.Get("/place/:id", getPlaceDetail)
	m.Get("/favorites", getFavoritesList)
	m.Get("/route/:id", getRouteDetail)
	m.Get("/history/:id", getHistoryList)

	m.Post("/ratio/", postRatioDetail)

	m.Delete("/user/:id", deleteUser)

	m.Put("/ratio/:id", putRation)
	m.Put("/user", putUser)

	m.Get("/logout", getLogout)
	m.Post("/login", getLogin)

	m.Run()
}
