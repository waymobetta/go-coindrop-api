package coindropverification

import (
	"github.com/waymobetta/go-coindrop-api/coindropdb"
	"github.com/waymobetta/go-coindrop-api/goreddit"
)

// VERIFICATION

// UpdateRedditVerificationCode updates the 2FA data of a single user
func UpdateRedditVerificationCode(u *goreddit.User) (*goreddit.User, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := coindropdb.Client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db update
	sqlStatement := `UPDATE coindropdb SET stored_twofa_code = $1, posted_twofa_code = $2, is_validated = $3 WHERE reddit_username = $4`

	// prepare statement
	stmt, err := coindropdb.Client.Prepare(sqlStatement)
	if err != nil {
		return u, err
	}

	defer stmt.Close()

	// execute db write using unique reddit username as the identifier
	_, err = stmt.Exec(u.Info.TwoFAData.StoredTwoFACode, u.Info.TwoFAData.PostedTwoFACode, u.Info.TwoFAData.IsValidated, u.Info.RedditData.Username)
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
