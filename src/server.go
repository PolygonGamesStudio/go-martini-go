package main

import (
	"github.com/codegangsta/martini"
)

func main() {
	m := martini.Classic()

	// TODO: регулярку для ID
	m.Get("/place", getPlaceDetail)
	m.Get("/favorites", getFavoritsList)
	m.Get("/route/{id}", getRouteDetail)
	m.Get("/history/", getHistoryList)

	m.Post("/ratio/{id}", postRatioDetail)

	m.Delete("/user/{id}", deleteUser)

	m.Put("/ratio/{id}", putRation)
	m.Put("/user/{id}", putUser)

	m.Get("/logout", getLogout)
	m.Post("/login", getLogin)

	m.Run()
}
