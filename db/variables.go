package db

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
	host     = os.Getenv("AWS_COINDROP_STAGING_HOST")
	port, _  = strconv.Atoi(os.Getenv("AWS_COINDROP_STAGING_PORT"))
	user     = os.Getenv("AWS_COINDROP_STAGING_USER")
	dbname   = os.Getenv("AWS_COINDROP_STAGING_DBNAME")
	password = os.Getenv("AWS_COINDROP_STAGING_PASS")
	// disable SSL for local testing
	sslmode = "disable"

	// environment variables for local db
	lHost     = os.Getenv("LOCAL_PG_HOST")
	lPort, _  = strconv.Atoi(os.Getenv("LOCAL_PG_PORT"))
	lUser     = os.Getenv("LOCAL_PG_USER")
	lDbname   = os.Getenv("LOCAL_PG_DBNAME")
	lPassword = os.Getenv("LOCAL_PG_PASS")
	// disable SSL for local testing
	lSslmode = "disable"
)
