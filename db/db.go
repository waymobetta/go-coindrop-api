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
	// clear screen
	wmb.Clear()

	args := os.Args

	if len(args) < 2 {
		fmt.Println("usage: go-coindrop-api <localhost/staging/prod>")
		os.Exit(0)
	}

	var psqlInfo string

	switch {
	case args[1] == "localhost":
		fmt.Println("[*] Connecting to localhost")
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", lHost, lPort, lUser, lPassword, lDbname, lSslmode)
	case args[1] == "staging":
		fmt.Println("[*] Connecting to staging")
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	case args[1] == "prod":
		fmt.Println("[*] Connecting to prod")
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", pHost, pPort, pUser, pPassword, pDbname)
	}

	_client, err := sql.Open("postgres", psqlInfo)
	Client = _client
	if err != nil {
		return err
	}
	fmt.Println("api ready..")
	return nil
}
