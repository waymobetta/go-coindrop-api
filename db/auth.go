package db

import (
	"database/sql"

	"github.com/waymobetta/go-coindrop-api/types"
)

// AddUserID inserts an AWS cognito user ID to the coindrop_auth table
func (db *DB) AddUserID(cognitoUserID string) error {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return err
	}

	// create SQL statement for db update
	sqlStatement := `
		INSERT INTO 
			coindrop_auth
			(
				cognito_auth_user_id
			)
		VALUES 
			(
				$1
			)
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return err
	}

	defer stmt.Close()

	// execute db write using unique ID as the identifier
	_, err = stmt.Exec(cognitoUserID)
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

// GetUser gets user by ID
func (db *DB) GetUser(userID string) (*types.User, error) {
	sqlStatement := `
		SELECT
			coindrop_auth.id
		FROM
			coindrop_auth
		WHERE
			coindrop_auth.cognito_auth_user_id = $1;
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// initialize row object
	row := stmt.QueryRow(userID)

	user := new(types.User)
	var cognitoAuthUserID sql.NullString

	// iterate over row object to retrieve queried value
	err = row.Scan(
		&user.ID,
	)

	user.CognitoAuthUserID = cognitoAuthUserID.String

	if err == sql.ErrNoRows {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserIDByCognitoUserID gets user ID matching the AWS Cognito auth user ID
func (db *DB) GetUserIDByCognitoUserID(cognitoUserID string) (string, error) {
	sqlStatement := `
		SELECT
			id
		FROM
			coindrop_auth
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
		return "", err
	}

	if err != nil {
		return "", err
	}

	return userID, nil
}
