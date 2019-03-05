package db

import (
	"github.com/lib/pq"
	"github.com/waymobetta/go-coindrop-api/types"
)

// AddRedditUser adds the listing and associated data of a single user
func (db *DB) AddRedditUser(u *types.User2) (*types.User2, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db writes
	sqlStatement := `
	INSERT INTO
		coindrop_reddit2
		(
			user_id,
			username,
			comment_karma,
			link_karma,
			subreddits,
			trophies,
			posted_verification_code,
			confirmed_verification_code,
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
			$9
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
		u.Social.Reddit.Verification.ConfirmedVerificationCode,
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
			*
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
			&user.ID,
			&user.AuthUserID,
			&user.Reddit.Username,
			&user.Reddit.CommentKarma,
			&user.Reddit.LinkKarma,
			pq.Array(&user.Reddit.Subreddits),
			pq.Array(&user.Reddit.Trophies),
			&user.Reddit.Verification.PostedVerificationCode,
			&user.Reddit.Verification.ConfirmedVerificationCode,
			&user.Reddit.Verification.Verified,
			// stack overflow
			&user.ID,
			&user.AuthUserID,
			&user.StackOverflow.ExchangeAccountID,
			&user.StackOverflow.UserID,
			&user.StackOverflow.DisplayName,
			pq.Array(&user.StackOverflow.Accounts),
			&user.StackOverflow.Verification.PostedVerificationCode,
			&user.StackOverflow.Verification.ConfirmedVerificationCode,
			&user.StackOverflow.Verification.Verified,
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
func (db *DB) GetRedditUser(u *types.User2) (*types.User2, error) {
	// create SQL statement for db writes
	sqlStatement := `
		SELECT
			coindrop_reddit2.username,
			coindrop_reddit2.comment_karma,
			coindrop_reddit2.link_karma,
			coindrop_reddit2.subreddits,
			coindrop_reddit2.trophies,
			coindrop_reddit2.posted_verification_code,
			coindrop_reddit2.confirmed_verification_code,
			coindrop_reddit2.verified
		FROM
			coindrop_auth2
		JOIN
			coindrop_reddit2
		ON
			coindrop_auth2.id = coindrop_reddit2.user_id
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
	sqlStatement := `DELETE FROM coindrop_reddit WHERE auth_user_id = $1`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique ID as the identifier
	_, err = stmt.Exec(u.AuthUserID)
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

/// REDDIT

// UpdateRedditInfo updates the listing and associated Reddit data of a single user
func (db *DB) UpdateRedditInfo(u *types.User2) (*types.User2, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `
		UPDATE
			coindrop_reddit2
		SET
			comment_karma = $1,
			link_karma = $2,
			subreddits = $3,
			trophies = $4
		FROM
			coindrop_auth2
		WHERE
			coindrop_auth2.id = coindrop_reddit2.user_id AND
			coindrop_auth2.cognito_auth_user_id = $5
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
func (db *DB) UpdateRedditVerificationCode(u *types.User2) (*types.User2, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `
		UPDATE
			coindrop_reddit2
		SET
			confirmed_verification_code = $1,
			posted_verification_code = $2,
			verified = $3
		FROM
			coindrop_auth2
		WHERE
			coindrop_auth2.id = coindrop_reddit2.user_id AND
			coindrop_auth2.cognito_auth_user_id = $4
	`
	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique reddit username as the identifier
	_, err = stmt.Exec(
		u.Social.Reddit.Verification.ConfirmedVerificationCode,
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
