package db

import (
	"database/sql"

	"github.com/waymobetta/go-coindrop-api/types"
)

// GetUserERC721s method returns all ERC721s tied to a specific user
func (db *DB) GetUserERC721s(userID string) ([]types.PublicBadge, error) {
	publicBadgeSlice := []types.PublicBadge{}

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

	rows, err := db.client.Query(sqlStatement, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per transaction in db to hold transaction info
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
		publicBadge.ERC721.ContractID = contractAddress.String
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

// // AddTransaction method adds a new transaction to the db
// func (db *DB) AddTransaction(trx *types.Transaction, resourceID string) error {
// 	// initialize statement write to database
// 	tx, err := db.client.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	// create SQL statement for db writes
// 	sqlStatement := `
// 		INSERT INTO
// 			coindrop_transactions
// 			(
// 				user_id,
// 				task_id,
// 				hash
// 			)
// 		VALUES
// 			(
// 				$1,
// 				(
// 				SELECT
// 					id
// 				FROM
// 					coindrop_tasks
// 				WHERE
// 					quiz_id = (
// 						SELECT
// 							id
// 						FROM
// 							coindrop_quizzes
// 						WHERE typeform_form_id = $2
// 					)
// 				),
// 				$3
// 			)
// 	`

// 	// prepare statement
// 	stmt, err := db.client.Prepare(sqlStatement)
// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	// execute db write using task title + associated data
// 	_, err = stmt.Exec(
// 		trx.UserID,
// 		resourceID,
// 		trx.Hash,
// 	)
// 	if err != nil {
// 		// rollback transaction if error thrown
// 		return tx.Rollback()
// 	}

// 	// commit db write
// 	err = tx.Commit()
// 	if err != nil {
// 		// rollback transaction if error thrown
// 		return tx.Rollback()
// 	}

// 	return nil
// }
