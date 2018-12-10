package coindropdb

import (
	"database/sql"
	"os"
	"strconv"
)

var (
	// TODO:
	// remove global variable usage

	// Client is a pointer to the sql DB object
	Client *sql.DB

	// environment variables for db
	host     = os.Getenv("LOCAL_PG_HOST")
	port, _  = strconv.Atoi(os.Getenv("LOCAL_PG_PORT"))
	user     = os.Getenv("LOCAL_PG_USER")
	dbname   = os.Getenv("LOCAL_PG_DBNAME")
	password = os.Getenv("LOCAL_PG_PASS")
	// disable SSL for local testing
	sslmode = "disable"
)
