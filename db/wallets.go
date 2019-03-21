package db

import (
	"database/sql"

	"github.com/waymobetta/go-coindrop-api/types"
)

// WALLET

// UpdateWallet updates the wallet address of a single user
func (db *DB) UpdateWallet(userID, newWalletAddress, walletType string) (*types.Wallet, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return nil, err
	}

	// create SQL statement for db update

	// TODO:
	// fix error
	// pq: there is no unique or exclusion constraint matching the ON CONFLICT specification

	// sqlStatement := `
	// 	INSERT INTO
	// 		coindrop_wallets(address, user_id)
	// 	VALUES(
	// 		$1, $2
	// 	)
	// 	ON CONFLICT (user_id)
	// 	DO UPDATE
	// 	SET
	// 		address = $1
	// 	WHERE
	// 		coindrop_wallets.user_id = $2
	// 	AND
	// 		coindrop_wallets.type = $3
	// `

	sqlStatement := `
		UPDATE SET
			address = $1
		WHERE
			coindrop_wallets.user_id = $2
		AND
			coindrop_wallets.type = $3
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	wallet := &types.Wallet{
		Address: newWalletAddress,
		Type:    walletType,
	}

	// execute db write using unique ID as the identifier
	_, err = stmt.Exec(newWalletAddress, userID, walletType)
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

	return wallet, nil
}

// AddWallet adds a new wallet for the user
func (db *DB) AddWallet(userID, newWalletAddress, walletType string) (*types.Wallet, error) {
	// for simplicity, update the listing rather than updating single value
	tx, err := db.client.Begin()
	if err != nil {
		return nil, err
	}

	// create SQL statement for db insert

	sqlStatement := `
		INSERT INTO
			coindrop_wallets (
				address,
				user_id,
				type
			) VALUES (
				$1,
				$2,
				$3
			)
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	wallet := &types.Wallet{
		Address: newWalletAddress,
		Type:    walletType,
	}

	// execute db write using unique ID as the identifier
	_, err = stmt.Exec(
		newWalletAddress,
		userID,
		walletType,
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

	return wallet, nil
}

// GetWallet returns a user's wallet based on type
func (db *DB) GetWallet(userID, walletType string) (*types.Wallet, error) {
	// create SQL statement for db update
	sqlStatement := `
		SELECT
			address,
			type
		FROM
			coindrop_auth
		JOIN
			coindrop_wallets
		ON
			coindrop_auth.id = coindrop_wallets.user_id
		WHERE
			user_id = (
				SELECT 
					id
				FROM 
					coindrop_auth
				WHERE 
					cognito_auth_user_id = $1
			) AND
			coindrop_wallets.type = $2
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// initialize row object
	row := stmt.QueryRow(
		userID,
		walletType,
	)

	wallet := &types.Wallet{}
	var walletAddress sql.NullString

	// iterate over row object to retrieve queried value
	err = row.Scan(
		&walletAddress,
		&wallet.Type,
	)

	wallet.Address = walletAddress.String
	wallet.Type = walletType

	if err == sql.ErrNoRows {
		return wallet, nil
	}

	if err != nil {
		return nil, err
	}

	return wallet, nil
}

// GetWallets returns a user's wallets
func (db *DB) GetWallets(userID string) ([]types.Wallet, error) {
	wallets := []types.Wallet{}

	// create SQL statement for db update
	sqlStatement := `
		SELECT
			coindrop_wallets.address,
			coindrop_wallets.type
		FROM
			coindrop_wallets
		WHERE
			coindrop_wallets.user_id = (
				SELECT 
					id
				FROM 
					coindrop_auth
				WHERE 
					cognito_auth_user_id = $1
			)
	`

	// prepare statement
	rows, err := db.client.Query(
		sqlStatement,
		userID,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per wallet in db to hold wallet info
		wallet := types.Wallet{}

		err = rows.Scan(
			&wallet.Address,
			&wallet.Type,
		)
		if err != nil {
			return nil, err
		}

		// append wallet object to slice of wallets
		wallets = append(wallets, wallet)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return wallets, nil
}
