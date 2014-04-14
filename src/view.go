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

func getFavoritesList(params martini.Params) []byte {
	dbmap := initDb()
	defer dbmap.Db.Close()

	log.Println("11")

	var objects []UserPlaceMTM
	_, err := dbmap.Select(&objects, "Select * from userPlacesMTM where userid=$1", params["id"])
	checkErr(err, "Select favorites failed")

	b, err := json.Marshal(objects)
	checkErr(err, "json.Marshal failed")

	return b
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

func getUserDetails(params martini.Params) []byte {
	dbmap := initDb()
	defer dbmap.Db.Close()

	obj, err := dbmap.Get(User{}, params["id"])

	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(obj)

	if err != nil {
		log.Fatal(err)
	}

	return b
	
}

func getAllCategories() []byte {
	dbmap := initDb()
	defer dbmap.Db.Close()

	var objects []Category
	_, err := dbmap.Select(&objects, "Select * from categories")

	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(objects)

	if err != nil {
		log.Fatal(err)
	}

	return b
}

//TODO:привести вывод ошибок и возврат ошибок к нормальному виду
//TODO:привести ответ сервера к нормальному виду
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
	useridString, placeidString, ratioString, feedback, isfavoriteString := r.FormValue("userid"), r.FormValue("placeid"), r.FormValue("ratio"), r.FormValue("feedback"), r.FormValue("isfavorite")
	
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
	isfavorite, err := strconv.ParseBool(isfavoriteString)
	if err != nil {
		log.Println("can't isfavorite ratio to bool")
		return nil
	}

	return &UserPlaceMTM{
		UserId:  int64(userid),
		PlaceId: int64(placeid),
		Ratio:  int8(ratio),
		Feedback: feedback,
		IsFavorite: isfavorite,
	}
}

func deleteUser(params martini.Params) string {
	dbmap := initDb()
	defer dbmap.Db.Close()
	
	obj, err := dbmap.Get(User{}, params["id"])
	checkErr(err, "can't get user by id")

	count, err := dbmap.Delete(obj)
	checkErr(err, "can't delete user")

	return "deleteUser " + string(count) + " rows succesfull"	//FIXME to string
}

func putRation() string {
	return "putRation "
}

func putUser(r *http.Request) string {
	dbmap := initDb()
	defer dbmap.Db.Close()

	user := getPutDataUser(r)
	id := &user.Id

	log.Println("userid: ", id)

	object,err := dbmap.Get(User{}, id)

	checkErr(err, "can't get updating user")

	object = user
	if user != nil {
		count, err := dbmap.Update(object)
		checkErr(err, "can't update user's data")
		return "updated " + strconv.FormatInt(count, 10) + " rows"
	} else{
		log.Println("can't get data from request")
	}


	return "putUser"
}
func getPutDataUser(r *http.Request) *User {//сделать так, чтобы обязательным был только id, остальные параметры не обязательны
	idString, login, password := r.FormValue("id"), r.FormValue("login"), r.FormValue("password") 
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println("can't convert userid to integer")
		return nil
	}

	return &User{
		Id: int64(id),
		Login: login,
		Password: password,
		Photo: "",		//FIXME
		Kilometers: int64(42),	//FIXME
		TasksCount: int64(42),	//FIXME
	}
}

func getLogout(params martini.Params) string {
	return "getLogout " + params["id"]
}
func getLogin(params martini.Params) string {
	return "getLogin " + params["id"]
}


