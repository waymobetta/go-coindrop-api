package db

import (
	"database/sql"

	"github.com/waymobetta/go-coindrop-api/types"
)

// AddUserID inserts an AWS cognito user ID to the coindrop_auth table
func (db *DB) AddUserID(u *types.User) (*types.User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
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
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique ID as the identifier
	_, err = stmt.Exec(u.CognitoAuthUserID)
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

// GetUser gets user by ID
func (db *DB) GetUser(userID string) (*types.User, error) {
	sqlStatement := `
		SELECT
			coindrop_auth.id,
			coindrop_auth.cognito_auth_user_id,
			coindrop_wallets.id as wallet_id,
			coindrop_wallets.address
		FROM
			coindrop_auth
		FULL OUTER JOIN
			coindrop_wallets
		ON
			coindrop_wallets.user_id = coindrop_auth.id
		WHERE
			coindrop_auth.id = $1;
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
	user.Wallet = new(types.Wallet)
	var cognitoAuthUserID sql.NullString
	var walletID sql.NullString
	var walletAddress sql.NullString

	// iterate over row object to retrieve queried value
	err = row.Scan(
		&user.ID,
		&cognitoAuthUserID,
		&walletID,
		&walletAddress,
	)

	user.CognitoAuthUserID = cognitoAuthUserID.String
	user.Wallet.ID = walletID.String
	user.Wallet.Address = walletAddress.String

	if err == sql.ErrNoRows {
		return nil, nil
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
		return "", nil
	}

	if err != nil {
		return "", err
	}

	return userID, nil
}
