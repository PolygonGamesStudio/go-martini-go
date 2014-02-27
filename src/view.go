package main

import (
	"encoding/json"
	"github.com/codegangsta/martini"
	"log"
)

func getRouteDetail(params martini.Params) []byte {
	type Message struct {
		Id   string
		Name string
		Body string
		Time int64
	}
	m := Message{params["id"], "Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
func getFavoritesList() string {
	return "getFavoritsList "
}
func getPlaceDetail() string {
	return "getPlaceDetail "
}
func getHistoryList() string {
	return "getHistoryList "
}
func postRatioDetail(params martini.Params) string {
	return "postRatioDetail " + params["id"]
}
func deleteUser(params martini.Params) string {
	return "deleteUser " + params["id"]
}
func putRation() string {
	return "putRation "
}
func putUser() string {
	return "putUser "
}

func getLogout(params martini.Params) string {
	return "getLogout " + params["id"]
}
func getLogin(params martini.Params) string {
	return "getLogin " + params["id"]
}
