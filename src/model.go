package main

type Place struct {
	Id              int64
	Description     string
	PriceCategory   int64
	Logo            string
	EndPointAddress string
	Creator
	CoordinatesGPS string
	IsActive       bool
	Route          string
	Category
}

type Creator struct {
	Id       int64
	Login    string
	Password string
}

type User struct {
	Id       int64
	Login    string
	Password string
}

type RouteNodes struct {
	Id          int64
	Description string
}

type Category struct {
	Id               int64
	Name             string
	BriefDescription string
	logo             string
}

type UserPlaceMTM struct {
	User
	Place
	Ratio    int8
	Feedback string
}
