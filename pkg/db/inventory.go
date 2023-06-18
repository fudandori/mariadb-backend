package db

import "database/sql"

type Place struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetPlaces(db *sql.DB) []Place {
	results, err := db.Query("SELECT id, name FROM place")
	if err != nil {
		panic(err.Error())
	}

	var output []Place
	for results.Next() {
		var place Place
		err := results.Scan(&place.Id, &place.Name)
		if err != nil {
			panic(err.Error())
		}
		output = append(output, place)
	}

	return output
}

func GetItems(db *sql.DB) []Item {
	results, err := db.Query("SELECT id, name FROM item")
	if err != nil {
		panic(err.Error())
	}

	var output []Item
	for results.Next() {
		var item Item
		err := results.Scan(&item.Id, &item.Name)
		if err != nil {
			panic(err.Error())
		}
		output = append(output, item)
	}

	return output
}
