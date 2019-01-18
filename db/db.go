package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/waymobetta/wmb"
)

// New initializes a new db connection; synonymous with init
func New() error {
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
		return err
	}
	fmt.Println("api ready..\n")
	return nil
}
