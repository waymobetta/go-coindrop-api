package gostackoverflow

import (
	"github.com/lib/pq"
	"github.com/waymobetta/go-coindrop-api/coindropdb"
)

// UpdateVerificationCode updates the verification code of a single user
func UpdateVerificationCode(s *StackOverflowData) (*StackOverflowData, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := coindropdb.Client.Begin()
	if err != nil {
		return s, err
	}

	// create SQL statement for db update
	sqlStatement := `UPDATE stackoverflowdb SET stored_verification_code = $1, posted_verification_code = $2, is_validated = $3 WHERE user_id = $4`

	// prepare statement
	stmt, err := coindropdb.Client.Prepare(sqlStatement)
	if err != nil {
		return s, err
	}

	defer stmt.Close()

	// execute db write using unique reddit username as the identifier
	_, err = stmt.Exec(s.VerificationData.StoredVerificationCode, s.VerificationData.PostedVerificationCode, s.VerificationData.IsVerified, s.UserID)
	if err != nil {
		// rollback transaction if error thrown
		tx.Rollback()
		return s, err
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaction if error thrown
		tx.Rollback()
		return s, err
	}

	return s, nil
}

/// STACK OVERFLOW

// AddStackUser adds the listing and associated data of a single user
func AddStackUser(s *StackOverflowData) (*StackOverflowData, error) {
	// initialize statement write to database
	tx, err := coindropdb.Client.Begin()
	if err != nil {
		return s, err
	}

	// create SQL statement for db writes
	sqlStatement := `INSERT INTO stackoverflowdb (exchange_account_id,user_id,display_name,accounts,posted_verification_code,stored_verification_code,is_validated) VALUES ($1,$2,$3,$4,$5,$6,$7)`

	// prepare statement
	stmt, err := coindropdb.Client.Prepare(sqlStatement)
	if err != nil {
		return s, err
	}

	defer stmt.Close()

	// execute db write using unique seller info hash to access data
	_, err = stmt.Exec(
		&s.ExchangeAccountID,
		&s.UserID,
		&s.DisplayName,
		pq.Array(&s.Accounts),
		&s.VerificationData.PostedVerificationCode,
		&s.VerificationData.StoredVerificationCode,
		&s.VerificationData.IsVerified,
	)
	if err != nil {
		// rollback transaction if error thrown
		tx.Rollback()
		return s, err
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaciton if error thrown
		tx.Rollback()
		return s, err
	}

	return s, err
}

// GetStackUser returns info for a single user
func GetStackUser(s *StackOverflowData) (*StackOverflowData, error) {
	// create SQL statement for db writes
	sqlStatement := `SELECT * FROM stackoverflowdb WHERE user_id = $1`

	// prepare statement
	stmt, err := coindropdb.Client.Prepare(sqlStatement)
	if err != nil {
		return s, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(s.UserID)

	// initialize new struct to hold user info
	var id int

	err = row.Scan(
		&id,
		&s.ExchangeAccountID,
		&s.UserID,
		&s.DisplayName,
		pq.Array(&s.Accounts),
		&s.VerificationData.PostedVerificationCode,
		&s.VerificationData.StoredVerificationCode,
		&s.VerificationData.IsVerified,
	)
	if err != nil {
		return s, err
	}

	return s, nil
}

// UpdateStackAboutInfo updates the listing and associated Reddit data of a single user
func UpdateStackAboutInfo(s *StackOverflowData) (*StackOverflowData, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := coindropdb.Client.Begin()
	if err != nil {
		return s, err
	}

	// create SQL statement for db update
	sqlStatement := `UPDATE stackoverflowdb SET exchange_account_id = $1, accounts = $2 WHERE user_id = $3`

	// prepare statement
	stmt, err := coindropdb.Client.Prepare(sqlStatement)
	if err != nil {
		return s, err
	}

	defer stmt.Close()

	// execute db write using unique reddit username as the identifier
	_, err = stmt.Exec(s.ExchangeAccountID, pq.Array(s.Accounts), s.UserID)
	if err != nil {
		// rollback transaction if error thrown
		tx.Rollback()
		return s, err
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaction if error thrown
		tx.Rollback()
		return s, err
	}

	return s, nil
}
