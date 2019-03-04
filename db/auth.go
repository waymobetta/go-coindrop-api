package db

import (
	"database/sql"
)

// AddUserID inserts an AWS cognito user ID to the coindrop_auth table
func (db *DB) AddUserID(u *User) (*User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `
		INSERT INTO coindrop_auth (auth_user_id)
		VALUES ($1)
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique ID as the identifier
	_, err = stmt.Exec(u.AuthUserID)
	if err != nil {
		// rollback transaction if error thrown
		tx.Rollback()
		return u, err
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaction if error thrown
		tx.Rollback()
		return u, err
	}

	return u, nil
}

// GetUserIDByCognitoUserID gets user ID matching the AWS Cognito auth user ID
func (db *DB) GetUserIDByCognitoUserID(cognitoUserID string) (string, error) {
	sqlStatement := `
		SELECT
			id
		FROM
			coindrop_auth2
		WHERE
			cognito_auth_user_id = $1;
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return "", err
	}

	defer stmt.Close()

	// initialize row object
	row := stmt.QueryRow(cognitoUserID)

	var userID string

	// iterate over row object to retrieve queried value
	err = row.Scan(
		&userID,
	)

	if err == sql.ErrNoRows {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	return userID, nil
}
