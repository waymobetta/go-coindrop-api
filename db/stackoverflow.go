package db

import (
	"github.com/lib/pq"
	"github.com/waymobetta/go-coindrop-api/types"
)

// AddStackUser adds the listing and associated data of a single user
func (db *DB) AddStackUser(u *types.User) (*types.User, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db writes
	sqlStatement := `
		INSERT INTO 
			coindrop_stackoverflow 
			(
				user_id,
				exchange_account_id,
				stack_user_id,
				display_name,
				accounts,
				posted_verification_code,
				verified
				)
		VALUES 
			(
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7
			)
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique seller info hash to access data
	_, err = stmt.Exec(
		&u.UserID,
		&u.Social.StackOverflow.ExchangeAccountID,
		&u.Social.StackOverflow.StackUserID,
		&u.Social.StackOverflow.DisplayName,
		&u.Social.StackOverflow.Accounts,
		&u.Social.StackOverflow.Verification.PostedVerificationCode,
		&u.Social.StackOverflow.Verification.Verified,
	)
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

	return u, err
}

// GetStackUser returns info for a single user
func (db *DB) GetStackUser(u *types.User) (*types.User, error) {
	// create SQL statement for db writes
	sqlStatement := `
		SELECT
			id,
			exchange_account_id,
			stack_user_id,
			display_name,
			accounts,
			posted_verification_code,
			confirmed_verification_code,
			verified
		FROM
			coindrop_stackoverflow
		WHERE
			user_id = $1
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// initialize row object
	row := stmt.QueryRow(u.UserID)

	// iterate over row object to retrieve queried value
	err = row.Scan(
		&u.Social.StackOverflow.ID,
		&u.Social.StackOverflow.ExchangeAccountID,
		&u.Social.StackOverflow.StackUserID,
		&u.Social.StackOverflow.DisplayName,
		pq.Array(&u.Social.StackOverflow.Accounts),
		&u.Social.StackOverflow.Verification.PostedVerificationCode,
		&u.Social.StackOverflow.Verification.ConfirmedVerificationCode,
		&u.Social.StackOverflow.Verification.Verified,
	)
	if err != nil {
		return u, err
	}

	return u, nil
}

// UpdateStackAboutInfo updates the listing and associated Reddit data of a single user
func (db *DB) UpdateStackAboutInfo(u *types.User) (*types.User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `
		UPDATE 
			coindrop_stackoverflow
		SET 
			exchange_account_id = $1, 
			display_name = $2, 
			accounts = $3
		WHERE
			user_id = $4
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique reddit username as the identifier
	_, err = stmt.Exec(
		u.Social.StackOverflow.ExchangeAccountID,
		u.Social.StackOverflow.DisplayName,
		u.Social.StackOverflow.Accounts,
		u.UserID,
	)
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

// UpdateStackProfileInfo updates the listing and associated Reddit data of a single user
func (db *DB) UpdateStackProfileInfo(u *types.User) (*types.User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `
		UPDATE 
			coindrop_stackoverflow
		SET 
			exchange_account_id = $1,
			display_name = $2
		WHERE
			user_id = $4
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique reddit username as the identifier
	_, err = stmt.Exec(
		u.Social.StackOverflow.ExchangeAccountID,
		u.Social.StackOverflow.DisplayName,
		u.UserID,
	)
	if err != nil {
		// rollback transaction if error thrown
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

// UpdateStackCommunityInfo updates the listing and associated Reddit data of a single user
func (db *DB) UpdateStackCommunityInfo(communityMap, userID string) error {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return err
	}

	// create SQL statement for db update
	sqlStatement := `
		UPDATE 
			coindrop_stackoverflow
		SET 
			accounts = $1
		WHERE
			user_id = $2
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return err
	}

	defer stmt.Close()

	// execute db write using unique reddit username as the identifier
	_, err = stmt.Exec(
		communityMap,
		userID,
	)
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

// UpdateStackVerificationCode updates the verification code of a single user
func (db *DB) UpdateStackVerificationCode(u *types.User) (*types.User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `
		UPDATE 
			coindrop_stackoverflow 
		SET 
			posted_verification_code = $1, 
			verified = $2
		WHERE
			user_id = $3
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique reddit username as the identifier
	_, err = stmt.Exec(
		u.Social.StackOverflow.Verification.PostedVerificationCode,
		u.Social.StackOverflow.Verification.Verified,
		u.UserID,
	)
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

// GetUserStackOverflowVerification ...
func (db *DB) GetUserStackOverfloVerification(u *types.User) (*types.User, error) {
	// create SQL statement for db writes
	sqlStatement := `
		SELECT
			stack_user_id,
			posted_verification_code,
			confirmed_verification_code,
			verified
		FROM
			coindrop_stackoverflow
		WHERE
			user_id = $1
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// initialize row object
	row := stmt.QueryRow(u.UserID)

	// iterate over row object to retrieve queried value
	err = row.Scan(
		&u.Social.StackOverflow.StackUserID,
		&u.Social.StackOverflow.Verification.PostedVerificationCode,
		&u.Social.StackOverflow.Verification.ConfirmedVerificationCode,
		&u.Social.StackOverflow.Verification.Verified,
	)
	if err != nil {
		return u, err
	}

	return u, nil
}
