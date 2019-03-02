package db

import (
	"github.com/lib/pq"
)

// AddStackUser adds the listing and associated data of a single user
func (db *DB) AddStackUser(u *User) (*User, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db writes
	sqlStatement := `INSERT INTO coindrop_stackoverflow (auth_user_id, exchange_account_id,user_id,display_name,accounts,posted_verification_code,stored_verification_code,is_verified) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique seller info hash to access data
	_, err = stmt.Exec(
		&u.AuthUserID,
		&u.StackOverflow.ExchangeAccountID,
		&u.StackOverflow.UserID,
		&u.StackOverflow.DisplayName,
		pq.Array(&u.StackOverflow.Accounts),
		&u.StackOverflow.Verification.PostedVerificationCode,
		&u.StackOverflow.Verification.ConfirmedVerificationCode,
		&u.StackOverflow.Verification.Verified,
	)
	if err != nil {
		// rollback transaction if error thrown
		tx.Rollback()
		return u, err
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaciton if error thrown
		tx.Rollback()
		return u, err
	}

	return u, err
}

// GetStackUser returns info for a single user
func (db *DB) GetStackUser(u *User) (*User, error) {
	// create SQL statement for db writes
	sqlStatement := `SELECT * FROM coindrop_stackoverflow WHERE auth_user_id = $1`

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
		&u.ID,
		&u.AuthUserID,
		&u.StackOverflow.ExchangeAccountID,
		&u.StackOverflow.UserID,
		&u.StackOverflow.DisplayName,
		pq.Array(&u.StackOverflow.Accounts),
		&u.StackOverflow.Verification.PostedVerificationCode,
		&u.StackOverflow.Verification.ConfirmedVerificationCode,
		&u.StackOverflow.Verification.Verified,
	)
	if err != nil {
		return u, err
	}

	return u, nil
}

// UpdateStackAboutInfo updates the listing and associated Reddit data of a single user
func (db *DB) UpdateStackAboutInfo(u *User) (*User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `UPDATE coindrop_stackoverflow SET exchange_account_id = $1, display_name = $2, accounts = $3 WHERE auth_user_id = $4`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique reddit username as the identifier
	_, err = stmt.Exec(u.StackOverflow.ExchangeAccountID, u.StackOverflow.DisplayName, pq.Array(u.StackOverflow.Accounts), u.AuthUserID)
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

// UpdateStackVerificationCode updates the verification code of a single user
func (db *DB) UpdateStackVerificationCode(u *User) (*User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `UPDATE coindrop_stackoverflow SET user_id = $1, stored_verification_code = $2, posted_verification_code = $3, is_verified = $4 WHERE auth_user_id = $5`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique reddit username as the identifier
	_, err = stmt.Exec(u.StackOverflow.UserID, u.StackOverflow.Verification.ConfirmedVerificationCode, u.StackOverflow.Verification.PostedVerificationCode, u.StackOverflow.Verification.Verified, u.AuthUserID)
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
