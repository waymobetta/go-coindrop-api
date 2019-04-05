package db

import (
	"database/sql"

	"github.com/waymobetta/go-coindrop-api/types"
)

// GetTransactions method returns all transactions recorded
func (db *DB) GetTransactions() ([]types.Transaction, error) {
	transactions := []types.Transaction{}

	sqlStatement := `
		SELECT 
			id,
			user_id,
			task_id,
			hash
		FROM
			coindrop_transactions
	`

	rows, err := db.client.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per transaction in db to hold transaction info
		transaction := types.Transaction{}

		err = rows.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.TaskID,
			&transaction.Hash,
		)
		if err != nil {
			return nil, err
		}
		// append transaction object to slice of transactions
		transactions = append(transactions, transaction)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

// GetTransaction method returns transaction based off of typeform form ID recorded
func (db *DB) GetTransactionByFormID(formID string) (*types.Transaction, error) {
	sqlStatement := `
		SELECT 
			coindrop_transactions.id,
			coindrop_transactions.user_id,
			coindrop_transactions.task_id,
			coindrop_transactions.hash
		FROM
			coindrop_transactions
		WHERE
			coindrop_transactions.task_id = (
				SELECT
					id
				FROM
					coindrop_tasks
				WHERE
					quiz_id = (
						SELECT
							id
						FROM
							coindrop_quizzes
						WHERE typeform_form_id = $1
					)
			)
	`

	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(formID)

	transaction := new(types.Transaction)
	var id sql.NullString
	var userId sql.NullString
	var taskId sql.NullString
	var hash sql.NullString

	err = row.Scan(
		&id,
		&userId,
		&taskId,
		&hash,
	)

	transaction.ID = id.String
	transaction.UserID = userId.String
	transaction.TaskID = taskId.String
	transaction.Hash = hash.String

	if err == sql.ErrNoRows {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return transaction, nil
}

// GetUserTransaction method returns all transactions tied to a specific user
func (db *DB) GetUserTransactions(userID string) ([]types.Transaction, error) {
	transactions := []types.Transaction{}

	sqlStatement := `
		SELECT
			id,
			task_id,
			hash
		FROM
			coindrop_transactions
		WHERE
			user_id = $1
	`

	rows, err := db.client.Query(sqlStatement, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per transaction in db to hold transaction info
		transaction := types.Transaction{}

		err = rows.Scan(
			&transaction.ID,
			&transaction.TaskID,
			&transaction.Hash,
		)
		if err != nil {
			return nil, err
		}
		// append transaction object to slice of transactions
		transactions = append(transactions, transaction)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

// AddTransaction method adds a new transaction to the db
func (db *DB) AddTransaction(trx *types.Transaction, resourceID string) error {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return err
	}

	// create SQL statement for db writes
	sqlStatement := `
		INSERT INTO 
			coindrop_transactions
			(
				user_id,
				task_id,
				hash
			)
		VALUES
			(
				$1,
				(
				SELECT
					id
				FROM
					coindrop_tasks
				WHERE
					quiz_id = (
						SELECT
							id
						FROM
							coindrop_quizzes
						WHERE typeform_form_id = $2
					)
				),
				$3
			)
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return err
	}

	defer stmt.Close()

	// execute db write using task title + associated data
	_, err = stmt.Exec(
		trx.UserID,
		resourceID,
		trx.Hash,
	)
	if err != nil {
		// rollback transaction if error thrown
		return tx.Rollback()
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaciton if error thrown
		return tx.Rollback()
	}

	return nil
}
