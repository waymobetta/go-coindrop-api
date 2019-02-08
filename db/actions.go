package db

import "encoding/json"

// GetAction returns all info for specific quiz
func (db *DB) GetAction(q *Quiz) (*Quiz, error) {
	// create SQL statement for db query
	sqlStatement := `SELECT * FROM coindrop_quizzes WHERE title = $1`

	// execute db query by passing in prepared SQL statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return q, err
	}

	defer stmt.Close()

	// initialize row object
	row := stmt.QueryRow(q.Title)

	// create temp string variable to store marshaled JSON
	var tempStr string

	// iterate over row object to retrieve queried value
	err = row.Scan(
		&q.ID,
		&q.Title,
		&tempStr,
	)
	if err != nil {
		return q, err
	}

	// Unmarshal JSON from temp string variable back into struct
	err = json.Unmarshal([]byte(tempStr), &q.QuizInfo.QuizData)
	if err != nil {
		return q, err
	}

	return q, nil
}

// AddAction adds the listing and associated data of a single quiz
func (db *DB) AddAction(q *Quiz) (*Quiz, error) {
	// marshal JSON for ease of storage
	byteArr, err := json.Marshal(&q.QuizInfo.QuizData)
	if err != nil {
		return q, err
	}

	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return q, err
	}

	// create SQL statement for db writes
	sqlStatement := `INSERT INTO coindrop_quizzes (title, quiz_data) VALUES ($1,$2)`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return q, err
	}

	defer stmt.Close()

	// execute db write using unique user ID + associated data
	_, err = stmt.Exec(
		&q.Title,
		// store marshaled JSON in db
		string(byteArr),
	)
	if err != nil {
		// rollback transaction if error thrown
		tx.Rollback()
		return q, err
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaciton if error thrown
		tx.Rollback()
		return q, err
	}

	return q, err
}