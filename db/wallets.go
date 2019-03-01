package db

import (
	"database/sql"
	"fmt"
)

// WALLET

// UpdateWallet updates the wallet address of a single user
func (db *DB) UpdateWallet(u *User) (*User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update

	// TODO:
	// INSERT (IF NOT EXISTS) + JOIN
	sqlStatement := `
		UPDATE
			coindrop_wallets
		SET 
			address = $1
		FROM
			coindrop_auth2
		WHERE
			coindrop_auth2.cognito_auth_user_id = $2
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique ID as the identifier
	res, err := stmt.Exec(u.WalletAddress, u.AuthUserID)
	if err != nil {
		// rollback transaction if error thrown
		id, err := res.RowsAffected()
		if err != nil {
			return u, err
		}

		if id != 1 {
			fmt.Println("ERROR")
		}

		return u, tx.Rollback()
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaction if error thrown
		return u, tx.Rollback()
	}

	return u, nil
}

// GetWallet updates the wallet address of a single user
func (db *DB) GetWallet(u *User) (*User, error) {
	// create SQL statement for db update

	sqlStatement := `
		SELECT address 
		FROM coindrop_auth2 
		JOIN coindrop_wallets 
		ON coindrop_auth2.wallet_id = coindrop_wallets.id
		WHERE cognito_auth_user_id = $1
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// initialize row object
	row := stmt.QueryRow(u.AuthUserID)

	// iterate over row object to retrieve queried value
	err = row.Scan(
		&u.WalletAddress,
	)

	if err == sql.ErrNoRows {
		return u, nil
	}

	if err != nil {
		return u, err
	}

	return u, nil
}
