package db

import (
	"database/sql"

	"github.com/waymobetta/go-coindrop-api/types"
)

// GetTasks returns all available tasks
func (db *DB) GetTasks(tasks *types.Tasks) (*types.Tasks, error) {
	// create SQL statement for db query
	sqlStatement := `
	SELECT
		 id,
		 title,
		 type,
		 author,
		 description,
		 token_name,
		 token_allocation,
		 badge_id,
		 quiz_id
	FROM
		coindrop_tasks
	`

	// execute db query by passing in prepared SQL statement
	rows, err := db.client.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per user in db to hold user info
		task := types.Task{BadgeData: new(types.Badge)}

		err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Type,
			&task.Author,
			&task.Description,
			&task.Token,
			&task.TokenAllocation,
			&task.BadgeData.ID,
			&task.QuizID,
		)
		if err != nil {
			return nil, err
		}
		// append task object to slice of tasks
		tasks.Tasks = append(tasks.Tasks, task)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// AddTask adds the listing and associated data of a single task
func (db *DB) AddTask(t *types.Task) (*types.Task, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return nil, err
	}

	// create SQL statement for db writes
	sqlStatement := `
	INSERT INTO coindrop_tasks
		(
			title,
			type,
			author,
			description,
			token_name,
			token_allocation,
			badge_id,
			quiz_id
		)
	VALUES
		(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8
		)
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// execute db write using task title + associated data
	_, err = stmt.Exec(
		t.Title,
		t.Type,
		t.Author,
		t.Description,
		t.Token,
		t.TokenAllocation,
		t.BadgeData.ID,
		t.QuizID,
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

	return t, nil
}

// GetUserTasks returns all info for specific quiz
func (db *DB) GetUserTasks(t *types.TaskUser) ([]types.Task, string, error) {
	tasks := []types.Task{}
	var quizURL string

	sqlStatement := `
		SELECT
		coindrop_tasks.id,
		coindrop_tasks.title,
		coindrop_tasks.type,
		coindrop_tasks.author,
		coindrop_tasks.description,
		coindrop_tasks.token_name,
		coindrop_tasks.token_allocation,
		coindrop_badges.name,
		coindrop_badges.description,
		coindrop_badges.logo_url,
		coindrop_tasks.badge_id,
		coindrop_tasks.logo_url,
		coindrop_tasks.quiz_id,
		coindrop_quizzes.typeform_form_url,
		coindrop_user_tasks.completed
	FROM
		coindrop_tasks
	JOIN
		coindrop_user_tasks
	ON
		coindrop_tasks.id = coindrop_user_tasks.task_id
	JOIN
		coindrop_badges
	ON
		coindrop_tasks.badge_id = coindrop_badges.id
	JOIN
		coindrop_quizzes
	ON
		coindrop_quizzes.id = coindrop_tasks.quiz_id
	WHERE
		coindrop_user_tasks.user_id = $1
	`

	rows, err := db.client.Query(sqlStatement, t.UserID)
	if err != nil {
		return nil, "", err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per task in db to hold task info
		task := types.Task{BadgeData: new(types.Badge)}

		var (
			tokenName       sql.NullString
			taskDescription sql.NullString
			tokenAllocation sql.NullInt64
			badgeID         sql.NullString
			logoURL         sql.NullString
			quizID          sql.NullString
			typeformURL     sql.NullString
		)

		err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Type,
			&task.Author,
			&taskDescription,
			&tokenName,
			&tokenAllocation,
			&task.BadgeData.Name,
			&task.BadgeData.Description,
			&task.BadgeData.LogoURL,
			&badgeID,
			&logoURL,
			&quizID,
			&typeformURL,
			&task.Completed,
		)
		if err != nil {
			return nil, "", err
		}

		task.Description = taskDescription.String
		task.Token = tokenName.String
		task.TokenAllocation = int(tokenAllocation.Int64)
		task.BadgeData.ID = badgeID.String
		task.LogoURL = logoURL.String
		task.QuizID = quizID.String
		quizURL = typeformURL.String

		// append task object to slice of tasks
		tasks = append(tasks, task)
	}

	err = rows.Err()
	if err != nil {
		return nil, "", err
	}

	return tasks, quizURL, nil
}

// GetUserTask returns all info for specific quiz
func (db *DB) GetUserTask(t *types.TaskUser) (*types.Task, string, error) {
	sqlStatement := `
	SELECT
		coindrop_tasks.id,
		coindrop_tasks.title,
		coindrop_tasks.type,
		coindrop_tasks.author,
		coindrop_tasks.description,
		coindrop_tasks.token_name,
		coindrop_tasks.token_allocation,
		coindrop_badges.name,
		coindrop_badges.description,
		coindrop_badges.logo_url,
		coindrop_tasks.badge_id,
		coindrop_tasks.logo_url,
		coindrop_tasks.quiz_id,
		coindrop_quizzes.typeform_form_url,
		coindrop_user_tasks.completed
	FROM
		coindrop_tasks
	JOIN
		coindrop_user_tasks
	ON
		coindrop_tasks.id = coindrop_user_tasks.task_id
	JOIN
		coindrop_badges
	ON
		coindrop_tasks.badge_id = coindrop_badges.id
	JOIN
		coindrop_quizzes
	ON
		coindrop_quizzes.id = coindrop_tasks.quiz_id
	WHERE
		coindrop_user_tasks.user_id $1
	AND
		coindrop_user_tasks.task_id = $2
	LIMIT
		1
	`

	rows, err := db.client.Query(sqlStatement, t.UserID, t.TaskID)
	if err != nil {
		return nil, "", err
	}

	defer rows.Close()

	// initialize new struct per task in db to hold task info
	task := &types.Task{
		BadgeData: new(types.Badge),
	}
	var quizURL string

	// iterate over rows
	for rows.Next() {
		var (
			tokenName       sql.NullString
			taskDescription sql.NullString
			tokenAllocation sql.NullInt64
			badgeID         sql.NullString
			logoURL         sql.NullString
			quizID          sql.NullString
			typeformURL     sql.NullString
		)

		err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Type,
			&task.Author,
			&taskDescription,
			&tokenName,
			&tokenAllocation,
			&task.BadgeData.Name,
			&task.BadgeData.Description,
			&task.BadgeData.LogoURL,
			&badgeID,
			&logoURL,
			&quizID,
			&typeformURL,
			&task.Completed,
		)
		if err != nil {
			return nil, "", err
		}

		task.Description = taskDescription.String
		task.Token = tokenName.String
		task.TokenAllocation = int(tokenAllocation.Int64)
		task.BadgeData.ID = badgeID.String
		task.LogoURL = logoURL.String
		task.QuizID = quizID.String
		quizURL = typeformURL.String
	}

	err = rows.Err()
	if err != nil {
		return nil, "", err
	}

	return task, quizURL, nil
}

// AddUserTask adds the listing and associated task data of a specific user
func (db *DB) AddUserTask(u *types.UserTask) (*types.UserTask, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return nil, err
	}

	// create SQL statement for db writes
	sqlStatement := `
		INSERT INTO
			coindrop_user_tasks(
				user_id,
				task_id,
				completed
			)
		VALUES
			(
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

	// execute db write using unique user ID + associated data
	_, err = stmt.Exec(
		u.UserID,
		u.TaskID,
		u.Completed,
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

	return u, nil
}

// MarkUserTaskCompletedFromQuiz adds a task to the user's list of completed tasks
func (db *DB) MarkUserTaskCompletedFromQuiz(r *types.QuizResults) (*types.QuizResults, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return nil, err
	}

	// create SQL statement for db writes
	sqlStatement := `
		UPDATE
			coindrop_user_tasks
		SET
			completed = $1
		WHERE
			task_id = (
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
				)
				AND
					user_id = $3
`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// execute db write using unique user ID + associated data
	_, err = stmt.Exec(
		r.QuizTaken,
		r.TypeformFormID,
		r.UserID,
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

	return r, nil
}

// MarkUserTaskCompleted adds a task to the user's list of completed tasks
func (db *DB) MarkUserTaskCompleted(u *types.UserTask) (*types.UserTask, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return nil, err
	}

	// create SQL statement for db writes
	sqlStatement := `
		UPDATE
			coindrop_user_tasks
		SET
			completed = $1
		WHERE
			task_id = $2
		AND
			user_id = $3
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// execute db write using unique user ID + associated data
	_, err = stmt.Exec(
		u.Completed,
		u.TaskID,
		u.UserID,
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

	return u, nil
}
