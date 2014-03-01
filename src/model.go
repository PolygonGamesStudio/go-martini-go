package main

type Place struct {
	Id              int64
	Description     string
	PriceCategory   int64
	Logo            string
	EndPointAddress string
	CreatorId       int64
	CoordinatesGPS  string
	IsActive        bool
	Route           string
	CategoryId      int64
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
	Logo             string
}

type UserPlaceMTM struct {
	UserId   int64
	PlaceId  int64
	Ratio    int8
	Feedback string
}
