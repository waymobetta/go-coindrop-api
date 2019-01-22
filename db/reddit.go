package db

import "github.com/lib/pq"

// AddRedditUser adds the listing and associated data of a single user
func AddRedditUser(u *User) (*User, error) {
	// initialize statement write to database
	tx, err := Client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db writes
	sqlStatement := `INSERT INTO coindrop_reddit (auth_user_id, username, comment_karma, link_karma, subreddits, trophies, posted_verification_code, stored_verification_code, is_verified) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`

	// prepare statement
	stmt, err := Client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique user ID + associated data
	_, err = stmt.Exec(
		u.Info.AuthUserID,
		u.Info.RedditData.Username,
		u.Info.RedditData.LinkKarma,
		u.Info.RedditData.CommentKarma,
		pq.Array(u.Info.RedditData.Subreddits),
		pq.Array(u.Info.RedditData.Trophies),
		u.Info.RedditData.VerificationData.PostedVerificationCode,
		u.Info.RedditData.VerificationData.StoredVerificationCode,
		u.Info.RedditData.VerificationData.IsVerified,
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
func GetUsers(users *Users) (*Users, error) {
	// create SQL statement for db query
	sqlStatement := `SELECT * FROM coindrop_reddit,coindrop_stackoverflow`

	// execute db query by passing in prepared SQL statement
	rows, err := Client.Query(sqlStatement)
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
			&user.Info.ID,
			&user.Info.AuthUserID,
			&user.Info.RedditData.Username,
			&user.Info.RedditData.CommentKarma,
			&user.Info.RedditData.LinkKarma,
			pq.Array(&user.Info.RedditData.Subreddits),
			pq.Array(&user.Info.RedditData.Trophies),
			&user.Info.RedditData.VerificationData.PostedVerificationCode,
			&user.Info.RedditData.VerificationData.StoredVerificationCode,
			&user.Info.RedditData.VerificationData.IsVerified,
			// stack overflow
			&user.Info.ID,
			&user.Info.AuthUserID,
			&user.Info.StackOverflowData.ExchangeAccountID,
			&user.Info.StackOverflowData.UserID,
			&user.Info.StackOverflowData.DisplayName,
			pq.Array(&user.Info.StackOverflowData.Accounts),
			&user.Info.StackOverflowData.VerificationData.PostedVerificationCode,
			&user.Info.StackOverflowData.VerificationData.StoredVerificationCode,
			&user.Info.StackOverflowData.VerificationData.IsVerified,
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
func GetRedditUser(u *User) (*User, error) {
	// create SQL statement for db writes
	sqlStatement := `SELECT * FROM coindrop_reddit WHERE auth_user_id = $1`

	// prepare statement
	stmt, err := Client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// initialize row object
	row := stmt.QueryRow(u.Info.AuthUserID)

	// iterate over row object to retrieve queried value
	err = row.Scan(
		&u.Info.ID,
		&u.Info.AuthUserID,
		&u.Info.RedditData.Username,
		&u.Info.RedditData.CommentKarma,
		&u.Info.RedditData.LinkKarma,
		pq.Array(&u.Info.RedditData.Subreddits),
		pq.Array(&u.Info.RedditData.Trophies),
		&u.Info.RedditData.VerificationData.PostedVerificationCode,
		&u.Info.RedditData.VerificationData.StoredVerificationCode,
		&u.Info.RedditData.VerificationData.IsVerified,
	)
	if err != nil {
		return u, err
	}

	return u, nil
}

// RemoveRedditUser removes the listing and associated data of a single user
func RemoveRedditUser(u *User) (*User, error) {
	// initialize statement write to database
	tx, err := Client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db writes
	sqlStatement := `DELETE FROM coindrop_reddit WHERE auth_user_id = $1`

	// prepare statement
	stmt, err := Client.Prepare(sqlStatement)
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
		// rollback transaciton if error thrown
		tx.Rollback()
		return u, err
	}

	return u, nil
}

/// REDDIT

// UpdateRedditInfo updates the listing and associated Reddit data of a single user
func UpdateRedditInfo(u *User) (*User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := Client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `UPDATE coindrop_reddit SET comment_karma = $1, link_karma = $2, subreddits = $3, trophies = $4 WHERE auth_user_id = $5`

	// prepare statement
	stmt, err := Client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique ID as the identifier
	_, err = stmt.Exec(u.Info.RedditData.CommentKarma, u.Info.RedditData.LinkKarma, pq.Array(u.Info.RedditData.Subreddits), pq.Array(u.Info.RedditData.Trophies), u.Info.AuthUserID)
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
func UpdateRedditVerificationCode(u *User) (*User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := Client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `UPDATE coindrop_reddit SET username = $1, stored_verification_code = $2, posted_verification_code = $3, is_verified = $4 WHERE auth_user_id = $5`

	// prepare statement
	stmt, err := Client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique reddit username as the identifier
	_, err = stmt.Exec(u.Info.RedditData.Username, u.Info.RedditData.VerificationData.StoredVerificationCode, u.Info.RedditData.VerificationData.PostedVerificationCode, u.Info.RedditData.VerificationData.IsVerified, u.Info.AuthUserID)
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