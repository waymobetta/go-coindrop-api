package db

import (
	"database/sql"
	"encoding/json"

	"github.com/waymobetta/go-coindrop-api/types"
)

// Subreddits is a list of approved subreddits used for targeting
var Subreddits = []string{
	"ethereum",
	"bitcoin",
	"ethtrader",
	"adchain",
	"cryptocurrency",
	"ethdev",
	"BATProject",
	"makerdao",
	"0xproject",
	"cryptoeconomics",
	"consensys",
	"ethereumclassic",
	"omise_go",
	"cardano",
	"dogecoin",
	"litecoin",
	"ripple",
	"monero",
	"augur",
	"maidsafe",
	"decentraland",
	"district0x",
	"spankchain",
	"joincolony",
	"everex",
	"maecenas",
	"storj",
	"IPFS",
	"loomnetwork",
	"dashpay",
	"bitcoincash",
	"eos",
	"binance",
	"stellar",
	"tezos",
	"iota",
	"zilliqa",
	"dfinity",
	"cosmosnetwork",
	"chainlink",
}

// GetEligibleRedditUsersAcrossSingleSub returns info for all users
func (db *DB) GetEligibleRedditUsersAcrossSingleSub(sub string, threshold int) ([]string, error) {

	var usersSlice []string

	// create SQL statement for db query
	// pulls a list of all user_id who are above or equal to the eligibility threshold of a single subreddits
	sqlStatement := `
		SELECT
			user_id
		FROM
			coindrop_reddit
		WHERE
			CAST (
			subreddits ->> $1 AS INTEGER
			) >= $2
	`

	// execute db query by passing in prepared SQL statement
	rows, err := db.client.Query(
		sqlStatement,
		sub,
		threshold,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per user in db to hold user info

		var user sql.NullString

		err = rows.Scan(
			&user,
		)
		if err != nil {
			return nil, err
		}
		// append user string to slice of users
		usersSlice = append(usersSlice, user.String)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return usersSlice, nil
}

// GetEligibleRedditUsersAcrossMultipleSubs returns info for all users
func (db *DB) GetEligibleRedditUsersAcrossMultipleSubs(sub1, sub2 string, threshold int) ([]string, error) {

	var usersSlice []string

	// create SQL statement for db query
	// pulls a list of all user_id who are above or equal to the eligibility threshold of multiple subreddits
	sqlStatement := `
		SELECT
			user_id
		FROM
			coindrop_reddit
		WHERE
			CAST (
			subreddits ->> $1 AS INTEGER
			) >= $3
		AND
			CAST (
			subreddits ->> $2 AS INTEGER
			) >= $3
	`

	// execute db query by passing in prepared SQL statement
	rows, err := db.client.Query(
		sqlStatement,
		sub1,
		sub2,
		threshold,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {

		var user sql.NullString

		err = rows.Scan(
			&user,
		)
		if err != nil {
			return nil, err
		}
		// append user ID to slice of users
		usersSlice = append(usersSlice, user.String)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return usersSlice, nil
}

// GetRedditUsersAndSubs returns IDs and subreddits of all reddit users (anonymous)
func (db *DB) GetRedditUsersAndSubs() ([]types.User, error) {
	usersSlice := []types.User{}

	// create SQL statement for db query
	// pulls a list of all user_id who are above or equal to the eligibility threshold of multiple subreddits
	sqlStatement := `
		SELECT
			user_id,
			subreddits
		FROM
			coindrop_reddit
	`

	// execute db query by passing in prepared SQL statement
	rows, err := db.client.Query(
		sqlStatement,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		user := types.User{}

		var userID sql.NullString
		var accountsString sql.NullString
		var accountsMap map[string]int

		err = rows.Scan(
			&user.UserID,
			&accountsString,
		)
		if err != nil {
			return nil, err
		}

		user.UserID = userID.String

		if accountsString.String == "" {
			user.Social.Reddit.Subreddits = accountsMap
		}

		byteArr := []byte(accountsString.String)

		err = json.Unmarshal(byteArr, &accountsMap)
		if err != nil {
			return nil, err
		}

		user.Social.Reddit.Subreddits = accountsMap

		// append user ID to slice of users
		usersSlice = append(usersSlice, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return usersSlice, nil
}
