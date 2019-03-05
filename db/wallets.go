package db

import (
	"database/sql"

	"github.com/waymobetta/go-coindrop-api/types"
)

// WALLET

// UpdateWallet updates the wallet address of a single user
func (db *DB) UpdateWallet(userID, newWalletAddress string) (*types.Wallet, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return nil, err
	}

	// create SQL statement for db update

	sqlStatement := `
		INSERT INTO
			coindrop_wallets(address, user_id)
		VALUES(
			$1, $2
		)
		ON CONFLICT (user_id)
		DO UPDATE
		SET
			address = $1
		WHERE
			coindrop_wallets.user_id = $2
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	wallet := &types.Wallet{
		Address: newWalletAddress,
	}

	// execute db write using unique ID as the identifier
	_, err = stmt.Exec(newWalletAddress, userID)
	if err != nil {
		// rollback transaction if error thrown
		tx.Rollback()
		return nil, err
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaction if error thrown
		tx.Rollback()
		return nil, err
	}

	return wallet, nil
}

// GetWallet updates the wallet address of a single user
func (db *DB) GetWallet(userID string) (*types.Wallet, error) {
	// create SQL statement for db update
	sqlStatement := `
		SELECT
			address
		FROM
			coindrop_wallets
		WHERE
			user_id = $1
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// initialize row object
	row := stmt.QueryRow(userID)

	wallet := &types.Wallet{}
	var walletAddress sql.NullString

	// iterate over row object to retrieve queried value
	err = row.Scan(
		&walletAddress,
	)

	wallet.Address = walletAddress.String

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return wallet, nil
}
