package db

import (
	"database/sql"
	"encoding/json"

	"github.com/lib/pq"
	"github.com/waymobetta/go-coindrop-api/types"
)

// AddRedditUser adds the listing and associated data of a single user
func (db *DB) AddRedditUser(u *types.User) (*types.User, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return nil, err
	}

	// create SQL statement for db writes
	sqlStatement := `
		INSERT INTO
			coindrop_reddit
			(
				user_id,
				username,
				comment_karma,
				link_karma,
				subreddits,
				trophies,
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
				$7,
				$8
			)
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var accountsString sql.NullString

	// execute db write using unique user ID + associated data
	_, err = stmt.Exec(
		u.UserID,
		u.Social.Reddit.Username,
		u.Social.Reddit.LinkKarma,
		u.Social.Reddit.CommentKarma,
		&accountsString,
		pq.Array(u.Social.Reddit.Trophies),
		u.Social.Reddit.Verification.PostedVerificationCode,
		u.Social.Reddit.Verification.Verified,
	)
	if err != nil {
		// rollback transaction if error throw
		return nil, tx.Rollback()
	}

	byteArr := []byte(accountsString.String)

	err = json.Unmarshal(byteArr, &u.Social.Reddit.Subreddits)
	if err != nil {
		return nil, err
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaction if error thrown
		return nil, tx.Rollback()
	}

	return u, nil
}

// UpdateRedditUser updates the listing and associated data of a single user
func (db *DB) UpdateRedditUser(u *types.User) (*types.User, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return nil, err
	}

	// create SQL statement for db writes
	sqlStatement := `
		UPDATE
			coindrop_reddit
		SET
			user_id = $1,
			username,
			comment_karma,
			link_karma,
			subreddits,
			trophies,
			posted_verification_code,
			verified
		VALUES
			(
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7,
				$8
			)
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// execute db write using unique user ID + associated data
	_, err = stmt.Exec(
		u.UserID,
		u.Social.Reddit.Username,
		u.Social.Reddit.LinkKarma,
		u.Social.Reddit.CommentKarma,
		u.Social.Reddit.Subreddits,
		pq.Array(u.Social.Reddit.Trophies),
		u.Social.Reddit.Verification.PostedVerificationCode,
		u.Social.Reddit.Verification.Verified,
	)
	if err != nil {
		// rollback transaction if error throw
		return nil, tx.Rollback()
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaction if error thrown
		return nil, tx.Rollback()
	}

	return u, nil
}

// GetRedditUser ...
func (db *DB) GetRedditUser(u *types.User) (*types.User, error) {
	// create SQL statement for db writes
	sqlStatement := `
		SELECT
			coindrop_reddit.id,
			coindrop_reddit.username,
			coindrop_reddit.comment_karma,
			coindrop_reddit.link_karma,
			coindrop_reddit.subreddits,
			coindrop_reddit.trophies,
			coindrop_reddit.posted_verification_code,
			coindrop_reddit.confirmed_verification_code,
			coindrop_reddit.verified
		FROM
			coindrop_reddit
		WHERE
			coindrop_reddit.user_id = $1
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// initialize row object
	row := stmt.QueryRow(u.UserID)

	var accountsString sql.NullString
	var accountsMap map[string]int

	// iterate over row object to retrieve queried value
	err = row.Scan(
		&u.Social.Reddit.ID,
		&u.Social.Reddit.Username,
		&u.Social.Reddit.CommentKarma,
		&u.Social.Reddit.LinkKarma,
		&accountsString,
		pq.Array(&u.Social.Reddit.Trophies),
		&u.Social.Reddit.Verification.PostedVerificationCode,
		&u.Social.Reddit.Verification.ConfirmedVerificationCode,
		&u.Social.Reddit.Verification.Verified,
	)
	if err != nil {
		return nil, err
	}

	if accountsString.String == "" {
		u.Social.Reddit.Subreddits = accountsMap
		return u, nil
	}

	byteArr := []byte(accountsString.String)

	err = json.Unmarshal(byteArr, &accountsMap)
	if err != nil {
		return nil, err
	}

	u.Social.Reddit.Subreddits = accountsMap

	return u, nil
}

// RemoveRedditUser removes the listing and associated data of a single user
func (db *DB) RemoveRedditUser(u *types.User) (*types.User, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return nil, err
	}

	// create SQL statement for db writes
	sqlStatement := `
		DELETE FROM
			coindrop_reddit
		WHERE
			user_id = $1
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// execute db write using unique ID as the identifier
	_, err = stmt.Exec(u.UserID)
	if err != nil {
		// rollback transaction if error thrown
		return nil, tx.Rollback()
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaction if error thrown
		return nil, tx.Rollback()
	}

	return u, nil
}

// UpdateRedditInfo updates the listing and associated Reddit data of a single user
func (db *DB) UpdateRedditInfo(u *types.User) (*types.User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return nil, err
	}

	// create SQL statement for db update
	sqlStatement := `
		UPDATE
			coindrop_reddit
		SET
			comment_karma = $1,
			link_karma = $2,
			subreddits = $3,
			trophies = $4
		WHERE
			user_id = $5
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// execute db write using unique ID as the identifier
	_, err = stmt.Exec(
		u.Social.Reddit.CommentKarma,
		u.Social.Reddit.LinkKarma,
		u.Social.Reddit.Subreddits,
		pq.Array(u.Social.Reddit.Trophies),
		u.UserID,
	)
	if err != nil {
		// rollback transaction if error thrown
		return nil, tx.Rollback()
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaction if error thrown
		return nil, tx.Rollback()
	}

	return u, nil
}

// UpdateRedditKarmaInfo updates the listing and Reddit karma data of a single user
func (db *DB) UpdateRedditKarmaInfo(u *types.User) (*types.User, error) {
	tx, err := db.client.Begin()
	if err != nil {
		return nil, err
	}

	// create SQL statement for db update
	sqlStatement := `
		UPDATE
			coindrop_reddit
		SET
			comment_karma = $1,
			link_karma = $2
		WHERE
			user_id = $3
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// execute db write using unique ID as the identifier
	_, err = stmt.Exec(
		u.Social.Reddit.CommentKarma,
		u.Social.Reddit.LinkKarma,
		u.UserID,
	)
	if err != nil {
		// rollback transaction if error thrown
		return nil, tx.Rollback()
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaction if error thrown
		return nil, tx.Rollback()
	}

	return u, nil
}

// UpdateRedditSubInfo updates the listing and Reddit submission data of a single user
func (db *DB) UpdateRedditSubInfo(subMap, userID string) error {
	tx, err := db.client.Begin()
	if err != nil {
		return err
	}

	// create SQL statement for db update
	sqlStatement := `
		UPDATE
			coindrop_reddit
		SET
			subreddits = $1
		WHERE
			user_id = $2
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return err
	}

	defer stmt.Close()

	// execute db write using unique ID as the identifier
	_, err = stmt.Exec(
		subMap,
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

// UpdateRedditTrophyInfo updates the listing and Reddit trophy data of a single user
func (db *DB) UpdateRedditTrophyInfo(u *types.User) (*types.User, error) {
	tx, err := db.client.Begin()
	if err != nil {
		return nil, err
	}

	// create SQL statement for db update
	sqlStatement := `
		UPDATE
			coindrop_reddit
		SET
			trophies = $1
		WHERE
			user_id = $2
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// execute db write using unique ID as the identifier
	_, err = stmt.Exec(
		pq.Array(u.Social.Reddit.Trophies),
		u.UserID,
	)
	if err != nil {
		// rollback transaction if error thrown
		return nil, tx.Rollback()
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaction if error thrown
		return nil, tx.Rollback()
	}

	return u, nil
}

/// VERIFICATION

// UpdateRedditVerificationCode ...
func (db *DB) UpdateRedditVerificationCode(u *types.User) (*types.User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return nil, err
	}

	// create SQL statement for db update
	sqlStatement := `
		UPDATE
			coindrop_reddit
		SET
			posted_verification_code = $1,
			verified = $2
		WHERE
			user_id = $3
	`
	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// execute db write using unique reddit username as the identifier
	_, err = stmt.Exec(
		u.Social.Reddit.Verification.PostedVerificationCode,
		u.Social.Reddit.Verification.Verified,
		u.UserID,
	)
	if err != nil {
		// rollback transaction if error thrown
		return nil, tx.Rollback()
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaction if error thrown
		return nil, tx.Rollback()
	}

	return u, nil
}

// GetUserRedditVerification ...
func (db *DB) GetUserRedditVerification(u *types.User) (*types.User, error) {
	// create SQL statement for db writes
	sqlStatement := `
		SELECT
			coindrop_reddit.username,
			coindrop_reddit.posted_verification_code,
			coindrop_reddit.confirmed_verification_code,
			coindrop_reddit.verified
		FROM
			coindrop_reddit
		WHERE
			coindrop_reddit.user_id = $1
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// initialize row object
	row := stmt.QueryRow(u.UserID)

	// iterate over row object to retrieve queried value
	err = row.Scan(
		&u.Social.Reddit.Username,
		&u.Social.Reddit.Verification.PostedVerificationCode,
		&u.Social.Reddit.Verification.ConfirmedVerificationCode,
		&u.Social.Reddit.Verification.Verified,
	)
	if err != nil {
		return nil, err
	}

	return u, nil
}
