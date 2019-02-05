package db

import "github.com/lib/pq"

// AddRedditUser adds the listing and associated data of a single user
func (db *DB) AddRedditUser(u *User) (*User, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db writes
	sqlStatement := `INSERT INTO coindrop_reddit (auth_user_id, username, comment_karma, link_karma, subreddits, trophies, posted_verification_code, stored_verification_code, is_verified) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique user ID + associated data
	_, err = stmt.Exec(
		u.AuthUserID,
		u.RedditData.Username,
		u.RedditData.LinkKarma,
		u.RedditData.CommentKarma,
		pq.Array(u.RedditData.Subreddits),
		pq.Array(u.RedditData.Trophies),
		u.RedditData.VerificationData.PostedVerificationCode,
		u.RedditData.VerificationData.StoredVerificationCode,
		u.RedditData.VerificationData.IsVerified,
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

// GetUsers returns info for all users
func (db *DB) GetUsers(users *Users) (*Users, error) {
	// create SQL statement for db query
	sqlStatement := `SELECT * FROM coindrop_reddit,coindrop_stackoverflow`

	// execute db query by passing in prepared SQL statement
	rows, err := db.client.Query(sqlStatement)
	if err != nil {
		return users, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per user in db to hold user info
		user := User{}
		err = rows.Scan(
			// reddit
			&user.ID,
			&user.AuthUserID,
			&user.RedditData.Username,
			&user.RedditData.CommentKarma,
			&user.RedditData.LinkKarma,
			pq.Array(&user.RedditData.Subreddits),
			pq.Array(&user.RedditData.Trophies),
			&user.RedditData.VerificationData.PostedVerificationCode,
			&user.RedditData.VerificationData.StoredVerificationCode,
			&user.RedditData.VerificationData.IsVerified,
			// stack overflow
			&user.ID,
			&user.AuthUserID,
			&user.StackOverflowData.ExchangeAccountID,
			&user.StackOverflowData.UserID,
			&user.StackOverflowData.DisplayName,
			pq.Array(&user.StackOverflowData.Accounts),
			&user.StackOverflowData.VerificationData.PostedVerificationCode,
			&user.StackOverflowData.VerificationData.StoredVerificationCode,
			&user.StackOverflowData.VerificationData.IsVerified,
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

// GetRedditUser returns info for a single user
func (db *DB) GetRedditUser(u *User) (*User, error) {
	// create SQL statement for db writes
	sqlStatement := `SELECT * FROM coindrop_reddit WHERE auth_user_id = $1`

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
		&u.RedditData.Username,
		&u.RedditData.CommentKarma,
		&u.RedditData.LinkKarma,
		pq.Array(&u.RedditData.Subreddits),
		pq.Array(&u.RedditData.Trophies),
		&u.RedditData.VerificationData.PostedVerificationCode,
		&u.RedditData.VerificationData.StoredVerificationCode,
		&u.RedditData.VerificationData.IsVerified,
	)
	if err != nil {
		return u, err
	}

	return u, nil
}

// RemoveRedditUser removes the listing and associated data of a single user
func (db *DB) RemoveRedditUser(u *User) (*User, error) {
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
func (db *DB) UpdateRedditInfo(u *User) (*User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `UPDATE coindrop_reddit SET comment_karma = $1, link_karma = $2, subreddits = $3, trophies = $4 WHERE auth_user_id = $5`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique ID as the identifier
	_, err = stmt.Exec(u.RedditData.CommentKarma, u.RedditData.LinkKarma, pq.Array(u.RedditData.Subreddits), pq.Array(u.RedditData.Trophies), u.AuthUserID)
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

// UpdateRedditVerificationCode updates the verification data of a single user
func (db *DB) UpdateRedditVerificationCode(u *User) (*User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `UPDATE coindrop_reddit SET username = $1, stored_verification_code = $2, posted_verification_code = $3, is_verified = $4 WHERE auth_user_id = $5`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique reddit username as the identifier
	_, err = stmt.Exec(u.RedditData.Username, u.RedditData.VerificationData.StoredVerificationCode, u.RedditData.VerificationData.PostedVerificationCode, u.RedditData.VerificationData.IsVerified, u.AuthUserID)
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
