package db

import (
	"database/sql"

	"github.com/waymobetta/go-coindrop-api/types"
)

// GetBadgesByRedditUsername returns info for all users
func (db *DB) GetBadgesByRedditUsername(redditUsername string) ([]*types.PublicBadge, error) {

	var publicBadgeSlice []*types.PublicBadge

	// create SQL statement for db query
	sqlStatement := `
		SELECT
			coindrop_badges.name,
			coindrop_badges.description,
			coindrop_badges.logo_url
		FROM
			coindrop_badges
		JOIN
			coindrop_tasks
		ON
			coindrop_tasks.badge_id = coindrop_badges.id
		JOIN
			coindrop_user_tasks
		ON
			coindrop_user_tasks.task_id = coindrop_tasks.id
		WHERE
			coindrop_user_tasks.completed = true
		AND
			coindrop_user_tasks.user_id = (
				SELECT
					coindrop_reddit.user_id
				FROM
					coindrop_reddit
				WHERE
					coindrop_reddit.username = $1
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
		publicBadge := new(types.PublicBadge)

		var (
			name        sql.NullString
			description sql.NullString
			logoURL     sql.NullString
			// badgeId sql.NullString
		)

		err = rows.Scan(
			&name,
			&description,
			&logoURL,
			// &badgeId,
		)
		if err != nil {
			return nil, err
		}

		publicBadge.Name = name.String
		publicBadge.Description = description.String
		publicBadge.LogoURL = logoURL.String
		// publicBadge.ID = badgeId.String

		// append badge struct to slice of badges
		publicBadgeSlice = append(publicBadgeSlice, publicBadge)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return publicBadgeSlice, nil
}
