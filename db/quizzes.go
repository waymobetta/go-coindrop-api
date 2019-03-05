package db

import (
	"encoding/json"

	"github.com/waymobetta/go-coindrop-api/types"
)

// GetQuiz returns all info for specific quiz
func (db *DB) GetQuiz(q *types.Quiz) (*types.Quiz, error) {
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

// AddQuiz adds the listing and associated data of a single quiz
func (db *DB) AddQuiz(q *types.Quiz) (*types.Quiz, error) {
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

// StoreQuizResults adds the quiz title and associated user results of a single quiz
func (db *DB) StoreQuizResults(q *types.QuizResults) (*types.QuizResults, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return q, err
	}

	// create SQL statement for db writes
	sqlStatement := `INSERT INTO coindrop_quiz_results (title, auth_user_id, questions_correct, questions_incorrect, has_taken_quiz) VALUES ($1,$2,$3,$4,$5)`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return q, err
	}

	defer stmt.Close()

	// execute db write using unique user ID + associated data
	_, err = stmt.Exec(
		q.Title,
		q.AuthUserID,
		q.QuestionsCorrect,
		q.QuestionsIncorrect,
		q.HasTakenQuiz,
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

// GetQuizResults returns all info for specific quiz
func (db *DB) GetQuizResults(q *types.QuizResults) (*types.QuizResults, error) {
	// create SQL statement for db query
	sqlStatement := `SELECT questions_correct, questions_incorrect, has_taken_quiz FROM coindrop_quiz_results WHERE title = $1 AND auth_user_id = $2`

	// execute db query by passing in prepared SQL statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return q, err
	}

	defer stmt.Close()

	// initialize row object
	row := stmt.QueryRow(q.Title, q.AuthUserID)

	// iterate over row object to retrieve queried value
	err = row.Scan(
		&q.QuestionsCorrect,
		&q.QuestionsIncorrect,
		&q.HasTakenQuiz,
	)
	if err != nil {
		return q, err
	}

	return q, nil
}

// GetAllQuizResults returns all info for specific quiz
func (db *DB) GetAllQuizResults(q *types.QuizResults, a *types.AllQuizResults) (*types.AllQuizResults, error) {
	// create SQL statement for db query
	sqlStatement := `SELECT id, title, questions_correct, questions_incorrect, has_taken_quiz FROM coindrop_quiz_results WHERE auth_user_id = $1`

	// execute db query by passing in prepared SQL statement
	rows, err := db.client.Query(sqlStatement, q.AuthUserID)
	if err != nil {
		return a, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per user in db to hold user info
		quizResults := types.QuizResults{}
		err = rows.Scan(
			&quizResults.ID,
			&quizResults.Title,
			&quizResults.QuestionsCorrect,
			&quizResults.QuestionsIncorrect,
			&quizResults.HasTakenQuiz,
		)
		if err != nil {
			return a, err
		}
		// append task object to slice of tasks
		a.QuizResults = append(a.QuizResults, quizResults)
	}
	err = rows.Err()
	if err != nil {
		return a, err
	}

	return a, nil
}
