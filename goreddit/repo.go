package goreddit

import (
	"github.com/lib/pq"
	"github.com/waymobetta/go-coindrop-api/coindropdb"
)

// AddUser adds the listing and associated data of a single user
func AddUser(u *User) (*User, error) {
	// initialize statement write to database
	tx, err := coindropdb.Client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db writes
	sqlStatement := `INSERT INTO coindropdb (reddit_username, wallet_address, comment_karma, link_karma, subreddits, trophies, posted_twofa_code, stored_twofa_code, is_validated) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`

	// prepare statement
	stmt, err := coindropdb.Client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique seller info hash to access data
	_, err = stmt.Exec(
		u.Info.RedditData.Username,
		u.Info.WalletAddress,
		u.Info.RedditData.LinkKarma,
		u.Info.RedditData.CommentKarma,
		pq.Array(u.Info.RedditData.Subreddits),
		pq.Array(u.Info.RedditData.Trophies),
		u.Info.TwoFAData.PostedTwoFACode,
		u.Info.TwoFAData.StoredTwoFACode,
		u.Info.TwoFAData.IsValidated,
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
	sqlStatement := `SELECT * FROM coindropdb`

	// execute db query by passing in prepared SQL statement
	rows, err := coindropdb.Client.Query(sqlStatement)
	if err != nil {
		return users, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per user in db to hold user info
		var id int
		user := User{}

		err = rows.Scan(
			&id,
			&user.Info.RedditData.Username,
			&user.Info.WalletAddress,
			&user.Info.RedditData.CommentKarma,
			&user.Info.RedditData.LinkKarma,
			pq.Array(&user.Info.RedditData.Subreddits),
			pq.Array(&user.Info.RedditData.Trophies),
			&user.Info.TwoFAData.PostedTwoFACode,
			&user.Info.TwoFAData.StoredTwoFACode,
			&user.Info.TwoFAData.IsValidated,
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

// GetUser returns info for a single user
func GetUser(u *User) (*User, error) {
	// create SQL statement for db writes
	sqlStatement := `SELECT * FROM coindropdb WHERE reddit_username = $1`

	// prepare statement
	stmt, err := coindropdb.Client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(u.Info.RedditData.Username)

	// initialize new struct to hold user info
	var id int

	err = row.Scan(
		&id,
		&u.Info.RedditData.Username,
		&u.Info.WalletAddress,
		&u.Info.RedditData.CommentKarma,
		&u.Info.RedditData.LinkKarma,
		pq.Array(&u.Info.RedditData.Subreddits),
		pq.Array(&u.Info.RedditData.Trophies),
		&u.Info.TwoFAData.PostedTwoFACode,
		&u.Info.TwoFAData.StoredTwoFACode,
		&u.Info.TwoFAData.IsValidated,
	)
	if err != nil {
		return u, err
	}

	return u, nil
}

// RemoveUser removes the listing and associated data of a single user
func RemoveUser(u *User) (*User, error) {
	// initialize statement write to database
	tx, err := coindropdb.Client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db writes
	sqlStatement := `DELETE FROM coindropdb WHERE reddit_username = $1`

	// prepare statement
	stmt, err := coindropdb.Client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique reddit username as the identifier
	_, err = stmt.Exec(u.Info.RedditData.Username)
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

// UpdateWallet updates the wallet address of a single user
func UpdateWallet(u *User) (*User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := coindropdb.Client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `UPDATE coindropdb SET wallet_address = $1 WHERE reddit_username = $2`

	// prepare statement
	stmt, err := coindropdb.Client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique reddit username as the identifier
	_, err = stmt.Exec(u.Info.WalletAddress, u.Info.RedditData.Username)
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

// REDDIT

// UpdateRedditInfo updates the listing and associated Reddit data of a single user
func UpdateRedditInfo(u *User) (*User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := coindropdb.Client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `UPDATE coindropdb SET comment_karma = $1, link_karma = $2, subreddits = $3, trophies = $4 WHERE reddit_username = $5`

	// prepare statement
	stmt, err := coindropdb.Client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique reddit username as the identifier
	_, err = stmt.Exec(u.Info.RedditData.CommentKarma, u.Info.RedditData.LinkKarma, pq.Array(u.Info.RedditData.Subreddits), pq.Array(u.Info.RedditData.Trophies), u.Info.RedditData.Username)
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
