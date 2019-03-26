package db

import (
	"github.com/waymobetta/go-coindrop-api/types"
)

// GetBadges method returns all badges available
func (db *DB) GetBadges() ([]types.Badge, error) {
	badges := []types.Badge{}

	sqlStatement := `
		SELECT 
			coindrop_badges.id,
			coindrop_badges.name,
			coindrop_badges.description
		FROM
			coindrop_badges
	`

	rows, err := db.client.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per badge in db to hold badge info
		badge := new(types.Badge)

		err = rows.Scan(
			&badge.ID,
			&badge.Name,
			&badge.Description,
		)
		if err != nil {
			return nil, err
		}

		// append badge object to slice of badges
		badges = append(badges, badge)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return badges, nil
}

// GetUserBadges method returns all badges tied to a specific user
func (db *DB) GetUserBadges(userID string) ([]types.Badge, error) {
	badges := []types.Badge{}

	sqlStatement := `
		SELECT
			coindrop_badges.id,
			coindrop_badges.name,
			coindrop_badges.description
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
			coindrop_user_tasks.user_id = $1
	`

	rows, err := db.client.Query(sqlStatement, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per badge in db to hold badge info
		badge := new(types.Badge)

		err = rows.Scan(
			&badge.ID,
			&badge.Name,
			&badge.Description,
		)
		if err != nil {
			return nil, err
		}
		// append badge object to slice of badges
		badges = append(badges, badge)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return badges, nil
}

// AddBadge method adds a new badge to the library
func (db *DB) AddBadge(badge *types.Badge) error {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return err
	}

	// create SQL statement for db writes
	sqlStatement := `
		INSERT INTO 
			coindrop_badges
			(
				name,
				description
			)
		VALUES
			(
				$1,
				$2	
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
		badge.Name,
		badge.Description,
	)
	if err != nil {
		// rollback transaction if error thrown
		return tx.Rollback()
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaciton if error thrown
		return tx.Rollback()
	}

	return nil
}
