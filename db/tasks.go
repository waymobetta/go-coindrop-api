package db

import (
	"database/sql"
	"fmt"
)

// GetTasks returns all available tasks
func (db *DB) GetTasks(tasks *Tasks) (*Tasks, error) {
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
		 badge_id
	FROM
		coindrop_tasks2
	`

	// execute db query by passing in prepared SQL statement
	rows, err := db.client.Query(sqlStatement)
	if err != nil {
		return tasks, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per user in db to hold user info
		task := Task{BadgeData: new(Badge)}

		err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Type,
			&task.Author,
			&task.Description,
			&task.Token,
			&task.TokenAllocation,
			&task.BadgeData.ID,
		)
		if err != nil {
			return tasks, err
		}
		// append task object to slice of tasks
		tasks.Tasks = append(tasks.Tasks, task)
	}
	err = rows.Err()
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

// AddTask adds the listing and associated data of a single task
func (db *DB) AddTask(t *Task) (*Task, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return t, err
	}

	// create SQL statement for db writes
	sqlStatement := `
	INSERT INTO coindrop_tasks2
		(
			title,
			type,
			author,
			description,
			token_name,
			token_allocation,
			badge_id
			)
	VALUES
		(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		)
	`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		return t, err
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
	)
	if err != nil {
		// rollback transaction if error thrown
		tx.Rollback()
		return t, err
	}

	// commit db write
	err = tx.Commit()
	if err != nil {
		// rollback transaciton if error thrown
		tx.Rollback()
		return t, err
	}

	return t, err
}

// GetUserTasks returns all info for specific quiz
func (db *DB) GetUserTasks(t *TaskUser) ([]Task, error) {
	tasks := []Task{}

	sqlStatement := `
	SELECT
		coindrop_tasks2.id,
		coindrop_tasks2.title,
		coindrop_tasks2.type,
		coindrop_tasks2.author,
		coindrop_tasks2.description,
		coindrop_tasks2.token_name,
		coindrop_tasks2.token_allocation,
		coindrop_tasks2.badge_id
	FROM
		coindrop_tasks2
	JOIN
		coindrop_user_tasks2
	ON
		coindrop_tasks2.id = coindrop_user_tasks2.task_id
	WHERE
		coindrop_user_tasks2.user_id = $1
	`

	rows, err := db.client.Query(sqlStatement, t.UserID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per task in db to hold task info
		task := Task{BadgeData: new(Badge)}

		var (
			tokenName       sql.NullString
			taskDescription sql.NullString
			tokenAllocation sql.NullInt64
			badgeID         sql.NullString
		)

		err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Type,
			&task.Author,
			&taskDescription,
			&tokenName,
			&tokenAllocation,
			&badgeID,
		)
		if err != nil {
			return nil, err
		}

		task.Description = taskDescription.String
		task.Token = tokenName.String
		task.TokenAllocation = int(tokenAllocation.Int64)
		task.BadgeData.ID = badgeID.String

		// append task object to slice of tasks
		tasks = append(tasks, task)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetUserTask returns all info for specific quiz
func (db *DB) GetUserTask(t *TaskUser) (*Task, error) {
	sqlStatement := `
	SELECT
		coindrop_tasks2.id,
		coindrop_tasks2.title,
		coindrop_tasks2.type,
		coindrop_tasks2.author,
		coindrop_tasks2.description,
		coindrop_tasks2.token_name,
		coindrop_tasks2.token_allocation,
		coindrop_tasks2.badge_id
	FROM
		coindrop_tasks2
	JOIN
		coindrop_user_tasks2
	ON
		coindrop_tasks2.id = coindrop_user_tasks2.task_id
	WHERE
		coindrop_user_tasks2.user_id = $1
	AND
		coindrop_user_tasks2.task_id = $2
	LIMIT
		1
	`

	rows, err := db.client.Query(sqlStatement, t.UserID, t.TaskID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// initialize new struct per task in db to hold task info
	task := &Task{BadgeData: new(Badge)}

	// iterate over rows
	for rows.Next() {
		var (
			tokenName       sql.NullString
			taskDescription sql.NullString
			tokenAllocation sql.NullInt64
			badgeID         sql.NullString
		)

		err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Type,
			&task.Author,
			&taskDescription,
			&tokenName,
			&tokenAllocation,
			&badgeID,
		)
		if err != nil {
			return nil, err
		}

		task.Description = taskDescription.String
		task.Token = tokenName.String
		task.TokenAllocation = int(tokenAllocation.Int64)
		task.BadgeData.ID = badgeID.String
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return task, nil
}

// AddUserTask adds the listing and associated task data of a specific user
func (db *DB) AddUserTask(u *UserTask2) (*UserTask2, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db writes
	sqlStatement := `
		INSERT INTO
			coindrop_user_tasks2(
				user_id,
				task_id,
				completed
			)
			VALUES ($1,$2,$3)`

	// prepare statement
	stmt, err := db.client.Prepare(sqlStatement)
	if err != nil {
		fmt.Println("HII")
		return u, err
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

// MarkUserTaskCompleted adds a task to the user's list of completed tasks
func (db *DB) MarkUserTaskCompleted(u *UserTask2) (*UserTask2, error) {
	// initialize statement write to database
	tx, err := db.client.Begin()
	if err != nil {
		return u, err
	}

	// create SQL statement for db writes
	sqlStatement := `
		UPDATE
			coindrop_user_tasks2
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
		return u, err
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
