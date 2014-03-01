package main

import (
	"database/sql"
	"github.com/codegangsta/martini"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	"log"
)

var (
	Dbm *gorp.DbMap
)

func main() {
	m := martini.Classic()

	m.Get("/place", getPlaceDetail)
	m.Get("/favorites", getFavoritesList)
	m.Get("/route/:id", getRouteDetail)
	m.Get("/history/", getHistoryList)

	m.Post("/ratio/:id", postRatioDetail)

	m.Delete("/user/:id", deleteUser)

	m.Put("/ratio/:id", putRation)
	m.Put("/user/:id", putUser)

	m.Get("/logout", getLogout)
	m.Post("/login", getLogin)

	dbMap := initDb()
	defer dbMap.Db.Close()

	m.Run()
}

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
