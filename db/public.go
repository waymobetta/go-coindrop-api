package db

import (
	"database/sql"

	"github.com/waymobetta/go-coindrop-api/types"
)

// GetBadgesByRedditUsername returns info for all users
func (db *DB) GetBadgesByRedditUsername(redditUsername string) ([]*types.PublicBadge, error) {

	var publicBadgeSlice []*types.PublicBadge

	sqlStatement := `
		SELECT
			coindrop_badges.name,
			coindrop_badges.description,
			coindrop_badges.logo_url,
			coindrop_tasks.author
		FROM
			coindrop_erc721s
		JOIN
			coindrop_tasks
		ON
			coindrop_tasks.badge_id = coindrop_erc721s.badge_id
		JOIN
			coindrop_badges
		ON
			coindrop_erc721s.badge_id = coindrop_badges.id
		WHERE
			coindrop_erc721s.user_id = (
				SELECT
					coindrop_reddit.user_id
				FROM
					coindrop_reddit
				WHERE
					coindrop_reddit.username = $1
			)
	`

	rows, err := db.client.Query(sqlStatement, redditUsername)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per transaction in db to hold transaction info
		publicBadge := new(types.PublicBadge)

		var (
			badgeName        sql.NullString
			badgeDescription sql.NullString
			badgeLogoURL     sql.NullString
			badgeProject     sql.NullString
		)

		err = rows.Scan(
			&badgeName,
			&badgeDescription,
			&badgeLogoURL,
			&badgeProject,
		)
		if err != nil {
			return nil, err
		}

		publicBadge.Name = badgeName.String
		publicBadge.Description = badgeDescription.String
		publicBadge.LogoURL = badgeLogoURL.String
		publicBadge.Project = badgeProject.String

		// append publicBadge object to slice of publicBadgeSlice
		publicBadgeSlice = append(publicBadgeSlice, publicBadge)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return publicBadgeSlice, nil
}
