count, err := dbmap.SelectInt("select count(*) from userplacesmtm")
	log.Println("count userPlacesMTM:", count)