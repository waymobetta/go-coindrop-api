package db

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq" // required
)

// DB ...
type DB struct {
	client *sql.DB
}

// Config ...
type Config struct {
	Host    string
	Port    int
	User    string
	Pass    string
	Dbname  string
	SSLMode string
}

// NewDB initializes a new database connection
func NewDB(config *Config) *DB {
	sslMode := "enable"
	if config.SSLMode != "" {
		sslMode = config.SSLMode
	}

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Pass, config.Dbname, sslMode)

	log.Printf("[db] connecting to database: %s:%v/%s", config.Host, config.Port, config.Dbname)

	client, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}
}
