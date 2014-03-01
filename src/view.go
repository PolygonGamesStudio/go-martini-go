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

var (
	myPlace Place
)

func getPlaceDetail(params martini.Params) []byte {
	myPlace, err := dbmap.Get(Place{}, params["id"])

	// myPlace := Place{}
	// err = dbmap.SelectOne(&myPlace, "select * from posts where post_id=$1", myplaceparams["id"])
	// checkErr(err, "SelectOne failed")
	b, err := json.Marshal(myPlace)
	if err != nil {
		log.Fatal(err)
	}
	return b
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
