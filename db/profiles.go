package db

import (
	"database/sql"

	"github.com/waymobetta/go-coindrop-api/types"
)

// UpsertProfile upserts user profile
func (db *DB) UpsertProfile(p *types.Profile) (*types.Profile, error) {
	tx, err := db.client.Begin()
	if err != nil {
		return nil, err
	}

	sqlStatement := `
		INSERT INTO
			coindrop_profiles
			(
				user_id,
				name,
				username
			)
		VALUES
			(
				$1, $2, $3
			)
		ON CONFLICT(user_id)
		DO UPDATE SET
			name = $2,
			username = $3;
	`

	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		p.UserID,
		p.Name,
		p.Username,
	)
	if err != nil {
		return nil, tx.Rollback()
	}

	err = tx.Commit()
	if err != nil {
		return nil, tx.Rollback()
	}

	return p, nil
}

// UpdateProfile update user profile
func (db *DB) UpdateProfile(p *types.Profile) (*types.Profile, error) {
	tx, err := db.client.Begin()
	if err != nil {
		return nil, err
	}

	sqlStatement := `
		UPDATE
			coindrop_profiles
		SET
			name = $1,
			username = $2
		WHERE
			user_id = $3
	`

	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		p.Name,
		p.Username,
		p.UserID,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return p, nil
}

// GetProfile returns user profile
func (db *DB) GetProfile(userID string) (*types.Profile, error) {
	sqlStatement := `
		SELECT
			name,
			username
		FROM
			coindrop_profiles
		WHERE
			user_id = $1
		`

	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(userID)

	var name sql.NullString
	var username sql.NullString

	err = row.Scan(
		&name,
		&username,
	)
	if err == sql.ErrNoRows {
		return &types.Profile{
			Name:     "",
			Username: "",
		}, nil
	}
	if err != nil {
		return nil, err
	}

	return &types.Profile{
		Name:     name.String,
		Username: username.String,
	}, nil
}
