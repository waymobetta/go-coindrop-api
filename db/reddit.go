package db

import (
	"github.com/lib/pq"
	"github.com/waymobetta/go-coindrop-api/types"
)

// AddRedditUser adds the listing and associated data of a single user
func (db *DB) AddRedditUser(u *types.User) (*types.User, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
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
			$8,
		)
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique user ID + associated data
	_, err = stmt.Exec(
		u.UserID,
		u.Social.Reddit.Username,
		u.Social.Reddit.LinkKarma,
		u.Social.Reddit.CommentKarma,
		pq.Array(u.Social.Reddit.Subreddits),
		pq.Array(u.Social.Reddit.Trophies),
		u.Social.Reddit.Verification.PostedVerificationCode,
		u.Social.Reddit.Verification.Verified,
	)
	if err != nil {
		// rollback transaction if error throw
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

// GetUsers returns info for all users
func (db *DB) GetUsers(users *types.Users) (*types.Users, error) {
	// create SQL statement for db query
	sqlStatement := `
		SELECT
			coindrop_reddit.user_id,
			coindrop_reddit.username,
			coindrop_reddit.comment_karma,
			coindrop_reddit.link_karma,
			coindrop_reddit.subreddits,
			coindrop_reddit.trophies,
			coindrop_reddit.verified,
			coindrop_stackoverflow.user_id,
			coindrop_stackoverflow.exchange_account_id,
			coindrop_stackoverflow.stack_user_id,
			coindrop_stackoverflow.display_name,
			coindrop_stackoverflow.accounts,
			coindrop_stackoverflow.verified
		FROM
			coindrop_reddit,
			coindrop_stackoverflow
	`

	// execute db query by passing in prepared SQL statement
	rows, err := db.client.Query(sqlStatement)
	if err != nil {
		return users, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per user in db to hold user info
		user := types.User{}
		err = rows.Scan(
			// reddit
			&user.Social.Reddit.Username,
			&user.Social.Reddit.CommentKarma,
			&user.Social.Reddit.LinkKarma,
			pq.Array(&user.Social.Reddit.Subreddits),
			pq.Array(&user.Social.Reddit.Trophies),
			&user.Social.Reddit.Verification.Verified,
			// stack overflow
			&user.Social.StackOverflow.UserID,
			&user.Social.StackOverflow.ExchangeAccountID,
			&user.Social.StackOverflow.DisplayName,
			pq.Array(&user.Social.StackOverflow.Accounts),
			&user.Social.StackOverflow.Verification.Verified,
		)
		if err != nil {
			return users, err
		}
		// append user object to slice of users
		users.Users = append(users.Users, user)
	}

	err = rows.Err()
	if err != nil {
		return users, err
	}

	return users, nil
}

// GetRedditUser ...
func (db *DB) GetRedditUser(u *types.User) (*types.User, error) {
	// create SQL statement for db writes
	sqlStatement := `
		SELECT
			coindrop_reddit.username,
			coindrop_reddit.comment_karma,
			coindrop_reddit.link_karma,
			coindrop_reddit.subreddits,
			coindrop_reddit.trophies,
			coindrop_reddit.posted_verification_code,
			coindrop_reddit.confirmed_verification_code,
			coindrop_reddit.verified
		FROM
			coindrop_auth
		JOIN
			coindrop_reddit
		ON
			coindrop_auth.id = coindrop_reddit.user_id
		WHERE
			cognito_auth_user_id = $1
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// initialize row object
	row := stmt.QueryRow(u.CognitoAuthUserID)

	// iterate over row object to retrieve queried value
	err = row.Scan(
		&u.Social.Reddit.Username,
		&u.Social.Reddit.CommentKarma,
		&u.Social.Reddit.LinkKarma,
		pq.Array(&u.Social.Reddit.Subreddits),
		pq.Array(&u.Social.Reddit.Trophies),
		&u.Social.Reddit.Verification.PostedVerificationCode,
		&u.Social.Reddit.Verification.ConfirmedVerificationCode,
		&u.Social.Reddit.Verification.Verified,
	)
	if err != nil {
		return u, err
	}

	return u, nil
}

// RemoveRedditUser removes the listing and associated data of a single user
func (db *DB) RemoveRedditUser(u *types.User) (*types.User, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db writes
	sqlStatement := `
	DELETE FROM
		coindrop_reddit
	WHERE
		auth_user_id = $1
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
		// rollback transaciton if error thrown
		tx.Rollback()
		return u, err
	}

	return u, nil
}

// UpdateRedditInfo updates the listing and associated Reddit data of a single user
func (db *DB) UpdateRedditInfo(u *types.User) (*types.User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
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
		FROM
			coindrop_auth
		WHERE
			coindrop_auth.id = coindrop_reddit.user_id AND
			coindrop_auth.cognito_auth_user_id = $5
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique ID as the identifier
	_, err = stmt.Exec(
		u.Social.Reddit.CommentKarma,
		u.Social.Reddit.LinkKarma,
		pq.Array(u.Social.Reddit.Subreddits),
		pq.Array(u.Social.Reddit.Trophies),
		u.CognitoAuthUserID,
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

/// VERIFICATION

// UpdateRedditVerificationCode ...
func (db *DB) UpdateRedditVerificationCode(u *types.User) (*types.User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `
		UPDATE
			coindrop_reddit
		SET
			posted_verification_code = $1,
			verified = $2
		FROM
			coindrop_auth
		WHERE
			coindrop_auth.id = coindrop_reddit.user_id AND
			coindrop_auth.cognito_auth_user_id = $3
	`
	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique reddit username as the identifier
	_, err = stmt.Exec(
		u.Social.Reddit.Verification.PostedVerificationCode,
		u.Social.Reddit.Verification.Verified,
		u.CognitoAuthUserID,
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
