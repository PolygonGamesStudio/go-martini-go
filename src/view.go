package main

import (
	"github.com/codegangsta/martini"
)

func getRouteDetail(params martini.Params) string {
	return "getRouteDetail " + params["id"]
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
