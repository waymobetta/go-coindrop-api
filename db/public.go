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
			coindrop_tasks.author,
			coindrop_badges.description,
			coindrop_badges.logo_url,
			coindrop_badges.erc721_contract_address,
			coindrop_erc721s.token_id
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
			badgeName        sql.NullString
			badgeProject     sql.NullString
			badgeDescription sql.NullString
			badgeLogoURL     sql.NullString
			contractAddress  sql.NullString
			tokenId          sql.NullString
		)

		err = rows.Scan(
			&badgeName,
			&badgeProject,
			&badgeDescription,
			&badgeLogoURL,
			&contractAddress,
			&tokenId,
		)
		if err != nil {
			return nil, err
		}

		publicBadge.Name = badgeName.String
		publicBadge.Project = badgeProject.String
		publicBadge.Description = badgeDescription.String
		publicBadge.LogoURL = badgeLogoURL.String
		publicBadge.ERC721.ContractAddress = contractAddress.String
		publicBadge.ERC721.TokenID = tokenId.String

		// append publicBadge object to slice of publicBadgeSlice
		publicBadgeSlice = append(publicBadgeSlice, publicBadge)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return publicBadgeSlice, nil
}
