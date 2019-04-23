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
			coindrop_tasks.author,
			coindrop_user_tasks.erc721_id,
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
		publicBadge := &types.PublicBadge{
			ERC721: new(types.ERC721),
		}

		var (
			project           sql.NullString
			name              sql.NullString
			description       sql.NullString
			logoURL           sql.NullString
			erc721Id          sql.NullString
			erc721TokenId     sql.NullString
			erc721ContractId  sql.NullString
			erc721TotalMinted sql.NullInt64
		)

		err = rows.Scan(
			&project,
			&erc721Id,
			&name,
			&description,
			&logoURL,
		)
		if err != nil {
			return nil, err
		}

		publicBadge.Project = project.String
		publicBadge.Name = name.String
		publicBadge.Description = description.String
		publicBadge.LogoURL = logoURL.String
		publicBadge.ERC721.ID = erc721Id.String
		publicBadge.ERC721.TokenID = erc721TokenId.String
		publicBadge.ERC721.ContractID = erc721ContractId.String
		publicBadge.ERC721.TotalMinted = int(erc721TotalMinted.Int64)

		// append badge struct to slice of badges
		publicBadgeSlice = append(publicBadgeSlice, publicBadge)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return publicBadgeSlice, nil
}