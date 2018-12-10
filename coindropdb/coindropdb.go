package coindropdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/waymobetta/wmb"
)

func init() {
	wmb.Clear()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	if host != "localhost" {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	}

	_client, err := sql.Open("postgres", psqlInfo)
	Client = _client
	if err != nil {
		log.Fatal(err)
	}
	// defer Client.Close()
	fmt.Println("api ready..\n")
}
