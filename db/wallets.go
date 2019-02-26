package db

import "database/sql"

// select address from coindrop_auth2 join coindrop_wallets on coindrop_auth2.wallet_id = coindrop_wallets.id where cognito_auth_user_id = 'b8f8a28f-cbf4-4477-9732-3f0b1b886cb8'

// WALLET

// UpdateWallet updates the wallet address of a single user
func (db *DB) UpdateWallet(u *User) (*User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `
		UPDATE coindrop_wallets
		SET address = $1
		FROM coindrop_auth2
		WHERE coindrop_auth2.cognito_auth_user_id = $2;
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique ID as the identifier
	_, err = stmt.Exec(u.WalletAddress, u.AuthUserID)
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
