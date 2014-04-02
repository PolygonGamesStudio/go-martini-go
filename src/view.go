package main

import (
	"database/sql"
	"encoding/json"
	"github.com/codegangsta/martini"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"net/http"
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
func getFavoritesList() string {	//stars >= 4
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

func getHistoryList(params martini.Params) []byte {
	dbmap := initDb()
	defer dbmap.Db.Close()

	var objects []UserPlaceMTM
	_, err := dbmap.Select(&objects, "Select * from userPlacesMTM where userid=$1", params["id"])


	b, err := json.Marshal(objects)

	if objects == nil {
		log.Println("vse huevo")
	}

	if err != nil {
		log.Fatal(err)
	}

	checkErr(err, "Select failed")
	
	return b
}

//TODO:привести вывод ошибок и возврат ошибок к нормальному виду
func postRatioDetail(r *http.Request) string {
	dbmap := initDb()
	defer dbmap.Db.Close()

	userPlacesMTM := getPostDataRatio(r)

	if userPlacesMTM != nil{
		err := dbmap.Insert(userPlacesMTM)
		checkErr(err, "can't insert data to userPlacesMTM")
		if err != nil {
			return "error\n"
		} else{
			return "successful inserted\n"
		}
	} else {
		log.Println("can't get data from request")
		return "error\n"
	}
		
}
func getPostDataRatio(r *http.Request) *UserPlaceMTM {
	useridString, placeidString, ratioString, feedback := r.FormValue("userid"), r.FormValue("placeid"), r.FormValue("ratio"), r.FormValue("feedback")
	
	userid, err := strconv.Atoi(useridString)
	if err != nil {
		log.Println("can't convert userid to integer")
		return nil
	}
	placeid, err := strconv.Atoi(placeidString)
	if err != nil {
		log.Println("can't convert placeid to integer")
		return nil
	}
	ratio, err := strconv.Atoi(ratioString)
	if err != nil {
		log.Println("can't convert ratio to integer")
		return nil
	}

	return &UserPlaceMTM{
		UserId:  int64(userid),
		PlaceId: int64(placeid),
		Ratio:  int8(ratio),
		Feedback: feedback,
	}
}

func deleteUser(params martini.Params) string {
	return "deleteUser " + params["id"]
}
func putRation() string {
	return "putRation "
}
func putUser() string {
	return ""
}

func getLogout(params martini.Params) string {
	return "getLogout " + params["id"]
}
func getLogin(params martini.Params) string {
	return "getLogin " + params["id"]
}


