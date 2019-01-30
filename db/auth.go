package db

// AddUserID inserts an AWS cognito user ID to the coindrop_auth table
func (db *DB) AddUserID(u *User) (*User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `INSERT INTO coindrop_auth (auth_user_id) VALUES ($1)`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique ID as the identifier
	_, err = stmt.Exec(u.Info.AuthUserID)
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

// WALLET

// UpdateWallet updates the wallet address of a single user
func (db *DB) UpdateWallet(u *User) (*User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `UPDATE coindrop_auth SET wallet_address = $1 WHERE auth_user_id = $2`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique ID as the identifier
	_, err = stmt.Exec(u.Info.WalletAddress, u.Info.AuthUserID)
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
	sqlStatement := `SELECT wallet_address FROM coindrop_auth WHERE auth_user_id=$1`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// initialize row object
	row := stmt.QueryRow(u.Info.AuthUserID)

	// iterate over row object to retrieve queried value
	err = row.Scan(
		&u.Info.WalletAddress,
	)

	if err != nil {
		return u, err
	}

	return u, nil
}
