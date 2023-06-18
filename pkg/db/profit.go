package db

import "database/sql"

type Earn struct {
	Day      string  `json:"day"`
	Profit   int     `json:"profit"`
	Exchange float32 `json:"exchange"`
	Invest   int     `json:"invest"`
}

type Client struct {
	Day  string `json:"day"`
	Name string `json:"name"`
}

func GetEarnings(db *sql.DB) []Earn {
	results, err := db.Query("SELECT day, profit, exchange, invest FROM earn")
	if err != nil {
		panic(err.Error())
	}

	var output []Earn
	for results.Next() {
		var earn Earn
		err := results.Scan(&earn.Day, &earn.Profit, &earn.Exchange, &earn.Invest)
		if err != nil {
			panic(err.Error())
		}
		output = append(output, earn)
	}

	return output
}

func GetClients(db *sql.DB) []Client {
	results, err := db.Query("SELECT day, name FROM membership")
	if err != nil {
		panic(err.Error())
	}

	var output []Client
	for results.Next() {
		var client Client
		err := results.Scan(&client.Day, &client.Name)
		if err != nil {
			panic(err.Error())
		}
		output = append(output, client)
	}

	return output
}
