package db

import (
	"github.com/waymobetta/go-coindrop-api/types"
)

// GetEligibleRedditUsersAcrossSingleSub returns info for all users
func (db *DB) GetEligibleRedditUsersAcrossSingleSub(sub string, threshold int) ([]types.User, error) {

	users := []types.User{}

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
		return users, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per user in db to hold user info
		user := types.User{}
		err = rows.Scan(
			user.UserID,
		)
		if err != nil {
			return users, err
		}
		// append user object to slice of users
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return users, err
	}

	return users, nil
}

// GetEligibleRedditUsersAcrossMultipleSubs returns info for all users
func (db *DB) GetEligibleRedditUsersAcrossMultipleSubs(sub1, sub2 string, threshold int) ([]types.User, error) {

	users := []types.User{}

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
		return users, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per user in db to hold user info
		user := types.User{}
		err = rows.Scan(
			user.UserID,
		)
		if err != nil {
			return users, err
		}
		// append user object to slice of users
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return users, err
	}

	return users, nil
}
