package main

import (
	"database/sql"
	"encoding/json"
	"github.com/codegangsta/martini"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	"log"
)

// func initDb() *gorp.DbMap {
// TODO: вынести работу с базой из view
func initDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("postgres", "user=postgres dbname=testdb host=localhost password=maxim321 sslmode=disable")
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(Creator{}, "creators").SetKeys(true, "Id")
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id")
	dbmap.AddTableWithName(RouteNodes{}, "routeNodes").SetKeys(true, "Id")
	dbmap.AddTableWithName(Category{}, "categories").SetKeys(true, "Id")
	dbmap.AddTableWithName(Place{}, "places").SetKeys(true, "Id")
	dbmap.AddTableWithName(UserPlaceMTM{}, "userPlacesMTM")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

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

func getPlaceDetail(params martini.Params) []byte {
	dbmap := initDb()
	defer dbmap.Db.Close()

	obj, err := dbmap.Get(Place{}, params["id"])
	myPlace := obj.(*Place)

	if obj == nil {
		log.Println("vse huevo")
	}

	if err != nil {
		log.Fatal(err)
	}

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
