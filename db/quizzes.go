package db

import (
	"database/sql"

	"github.com/waymobetta/go-coindrop-api/types"
)

// GetQuiz returns all info for specific quiz
func (db *DB) GetQuiz(quizID string) (*types.Quiz, error) {
	// create SQL statement for db query
	sqlStatement := `
		SELECT
			title,
			quiz_url
		FROM
			coindrop_quizzes
		WHERE
			id = $1
		`

	// execute db query by passing in prepared SQL statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// initialize row object
	row := stmt.QueryRow(quizID)

	var title sql.NullString
	var quizURL sql.NullString

	// iterate over row object to retrieve queried value
	err = row.Scan(
		&title,
		&quizURL,
	)
	if err != nil {
		return nil, err
	}

	return &types.Quiz{
		ID:    quizID,
		Title: title.String,
		//QuizURL: quizURL.String,
	}, nil
}

// GetQuizzes returns all info for specific quiz
func (db *DB) GetQuizzes() ([]*types.Quiz, error) {
	// create SQL statement for db query
	sqlStatement := `
		SELECT
			id,
			title,
			quiz_url
		FROM
			coindrop_quizzes
		`

	// execute db query by passing in prepared SQL statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := db.client.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var quizzes []*types.Quiz

	for rows.Next() {
		var quizID sql.NullString
		var title sql.NullString
		var quizURL sql.NullString

		// iterate over row object to retrieve queried value
		err = rows.Scan(
			&quizID,
			&title,
			&quizURL,
		)
		if err != nil {
			return nil, err
		}

		quizzes = append(quizzes, &types.Quiz{
			ID:    quizID.String,
			Title: title.String,
			//QuizURL: quizURL.String,
		})
	}

	return quizzes, nil
}

// AddQuiz adds the listing and associated data of a single quiz
func (db *DB) AddQuiz(quiz *types.Quiz) (*types.Quiz, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return nil, err
	}

	// create SQL statement for db writes
	sqlStatement := `
		INSERT INTO
			coindrop_quizzes(title, quiz_url)
		VALUES ($1, $2)`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// execute db write using unique user ID + associated data
	_, err = stmt.Exec(
		quiz.Title,
		"",
	)
	if err != nil {
		// rollback transaction if error thrown
		tx.Rollback()
		return nil, err
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaciton if error thrown
		tx.Rollback()
		return nil, err
	}

	return quiz, nil
}

// AddQuizResults adds the quiz title and associated user results of a single quiz
func (db *DB) AddQuizResults(r *types.QuizResults) (*types.QuizResults, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return nil, err
	}

	// create SQL statement for db writes
	sqlStatement := `
		INSERT INTO
			coindrop_quiz_results(
				typeform_form_id,
				user_id,
				questions_correct,
				questions_incorrect,
				quiz_taken,
				quiz_id
			)
			VALUES 
			(
				$1,
				$2,
				$3,
				$4,
				$5, 
				(
					SELECT 
						id 
					FROM 
						coindrop_quizzes
					WHERE 
						typeform_form_id = $1
				)
			)
		`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// execute db write using unique user ID + associated data
	_, err = stmt.Exec(
		r.TypeformFormID,
		r.UserID,
		r.QuestionsCorrect,
		r.QuestionsIncorrect,
		r.QuizTaken,
	)
	if err != nil {
		// rollback transaction if error thrown
		tx.Rollback()
		return nil, err
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaciton if error thrown
		tx.Rollback()
		return nil, err
	}

	return r, err
}

// GetQuizResults returns all info for specific quiz
func (db *DB) GetQuizResults(quizID, userID string) (*types.QuizResults, error) {
	// create SQL statement for db query
	sqlStatement := `
		SELECT
			questions_correct,
			questions_incorrect,
			quiz_taken
		FROM
			coindrop_quiz_results
		WHERE
			quiz_id = $1
		AND
			user_id = $2
	`

	// execute db query by passing in prepared SQL statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// initialize row object
	row := stmt.QueryRow(quizID, userID)

	var questionsCorrect sql.NullInt64
	var questionsIncorrect sql.NullInt64
	var quizTaken sql.NullBool

	// iterate over row object to retrieve queried value
	err = row.Scan(
		&questionsCorrect,
		&questionsIncorrect,
		&quizTaken,
	)
	if err == sql.ErrNoRows {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &types.QuizResults{
		QuizID:             quizID,
		UserID:             userID,
		QuestionsCorrect:   int(questionsCorrect.Int64),
		QuestionsIncorrect: int(questionsIncorrect.Int64),
		QuizTaken:          quizTaken.Bool,
	}, nil
}

// GetAllQuizResults returns all quiz results
func (db *DB) GetAllQuizResults(userID string) ([]*types.QuizResults, error) {
	// create SQL statement for db query
	sqlStatement := `
		SELECT
			quiz_id,
			typeform_form_id,
			questions_correct,
			questions_incorrect,
			quiz_taken
		FROM
			coindrop_quiz_results
		WHERE
			user_id = $1`

	// initialize row object
	rows, err := db.client.Query(sqlStatement, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []*types.QuizResults

	for rows.Next() {
		var quizID sql.NullString
		var typeformID sql.NullString
		var questionsCorrect sql.NullInt64
		var questionsIncorrect sql.NullInt64
		var quizTaken sql.NullBool

		// iterate over row object to retrieve queried value
		err = rows.Scan(
			&quizID,
			&typeformID,
			&questionsCorrect,
			&questionsIncorrect,
			&quizTaken,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, &types.QuizResults{
			QuizID:             quizID.String,
			TypeformFormID:     typeformID.String,
			UserID:             userID,
			QuestionsCorrect:   int(questionsCorrect.Int64),
			QuestionsIncorrect: int(questionsIncorrect.Int64),
			QuizTaken:          quizTaken.Bool,
		})
	}

	return results, nil
}
