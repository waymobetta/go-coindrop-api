package db

import (
	"database/sql"

	"github.com/waymobetta/go-coindrop-api/types"
)

// GetUserERC721s method returns all ERC721s tied to a specific user
func (db *DB) GetUserERC721s(userID string) ([]*types.PublicBadge, error) {
	publicBadgeSlice := []*types.PublicBadge{}

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

	rows, err := db.client.Query(sqlStatement, userID)
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

// GetTaskAndBadgeBy721Id method returns the task details tied to the specific ERC721
func (db *DB) GetTaskAndBadgeBy721Id(tokenId string) (*types.ERC721Lookup, error) {
	sqlStatement := `
		SELECT
			coindrop_tasks.title,
			coindrop_tasks.type,
			coindrop_tasks.author,
			coindrop_tasks.description,
			coindrop_tasks.logo_url,
			coindrop_badges.name,
			coindrop_badges.description,
			coindrop_badges.logo_url,
			coindrop_badges.erc721_contract_address
		FROM
			coindrop_tasks
		JOIN
			coindrop_erc721s
		ON
			coindrop_erc721s.badge_id = coindrop_tasks.badge_id
		JOIN
			coindrop_badges
		ON
			coindrop_badges.id = coindrop_erc721s.badge_id
		WHERE
			coindrop_erc721s.token_id = $1
	`

	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(tokenId)

	erc721Lookup := &types.ERC721Lookup{
		Task: &types.Task{
			BadgeData: new(types.Badge),
		},
		ERC721: new(types.ERC721),
	}

	var taskTitle sql.NullString
	var taskType sql.NullString
	var taskAuthor sql.NullString
	var taskDescription sql.NullString
	var taskLogoURL sql.NullString
	var badgeName sql.NullString
	var badgeDescription sql.NullString
	var badgeLogoURL sql.NullString
	var badgeERC721ContractAddress sql.NullString

	err = row.Scan(
		&taskTitle,
		&taskType,
		&taskAuthor,
		&taskDescription,
		&taskLogoURL,
		&badgeName,
		&badgeDescription,
		&badgeLogoURL,
		&badgeERC721ContractAddress,
	)
	if err == sql.ErrNoRows {
		return erc721Lookup, nil
	}
	if err != nil {
		return nil, err
	}

	erc721Lookup.Task.Title = taskTitle.String
	erc721Lookup.Task.Type = taskType.String
	erc721Lookup.Task.Author = taskAuthor.String
	erc721Lookup.Task.Description = taskDescription.String
	erc721Lookup.Task.LogoURL = taskLogoURL.String
	erc721Lookup.Task.BadgeData.Name = badgeName.String
	erc721Lookup.Task.BadgeData.Description = badgeDescription.String
	erc721Lookup.Task.BadgeData.LogoURL = badgeLogoURL.String
	erc721Lookup.ERC721.TokenID = tokenId
	erc721Lookup.ERC721.ContractAddress = badgeERC721ContractAddress.String

	return erc721Lookup, nil
}

// AssignERC721ToUser method adds a new transaction to the db
func (db *DB) AssignERC721ToUser(
	tokenId,
	badgeId,
	userId string,
) error {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return err
	}

	// create SQL statement for db writes
	sqlStatement := `
		INSERT INTO
			coindrop_erc721s(
				token_id,
				badge_id,
				user_id
			)
		VALUES (
			$1,
			$2,
			$3
		)
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return err
	}

	defer stmt.Close()

	// execute db write using task title + associated data
	_, err = stmt.Exec(
		tokenId,
		badgeId,
		userId,
	)
	if err != nil {
		// rollback transaction if error thrown
		return tx.Rollback()
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaction if error thrown
		return tx.Rollback()
	}

	return nil
}
