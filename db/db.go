package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/waymobetta/wmb"
)

func init() {
	wmb.Clear()

	args := os.Args

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)

	if len(args) > 1 {
		fmt.Println("[+] Connecting to localhost\n")
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", lHost, lPort, lUser, lPassword, lDbname, lSslmode)
	}

	_client, err := sql.Open("postgres", psqlInfo)
	Client = _client
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("api ready..\n")
}
