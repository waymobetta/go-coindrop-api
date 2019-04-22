package db

import (
	"database/sql"

	"github.com/waymobetta/go-coindrop-api/types"
)

// GetBadgesByRedditUsername returns info for all users
func (db *DB) GetBadgesByRedditUsername(redditUsername string) ([]*types.Badge, error) {

	var badgeSlice []*types.Badge

	// create SQL statement for db query
	sqlStatement := `
		SELECT
			badge_id
		FROM
			coindrop_public
		WHERE
			user_id = (
				SELECT
					user_id
				FROM
					coindrop_reddit
				WHERE
					username = $1
			)
	`

	// execute db query by passing in prepared SQL statement
	rows, err := db.client.Query(
		sqlStatement,
		redditUsername,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per user in db to hold user info
		badge := new(types.Badge)

		var (
			// name        sql.NullString
			// description sql.NullString
			// logoURL     sql.NullString
			badgeId sql.NullString
		)

		err = rows.Scan(
			&badgeId,
		)
		if err != nil {
			return nil, err
		}

		// badge.Name = name.String
		// badge.Description = description.String
		// badge.LogoURL = logoURL.String
		badge.ID = badgeId.String

		// append badge struct to slice of badges
		badgeSlice = append(badgeSlice, badge)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return badgeSlice, nil
}
