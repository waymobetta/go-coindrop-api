package db

// TASKS

// GetTasks returns all available tasks
func GetTasks(tasks *Tasks) (*Tasks, error) {
	// create SQL statement for db query
	sqlStatement := `SELECT * FROM coindrop_tasks`

	// execute db query by passing in prepared SQL statement
	rows, err := Client.Query(sqlStatement)
	if err != nil {
		return tasks, err
	}

	defer rows.Close()

	// iterate over rows
	for rows.Next() {
		// initialize new struct per user in db to hold user info
		task := Task{}
		err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Type,
			&task.Author,
			&task.Description,
			&task.Token,
			&task.TokenAllocation,
			&task.BadgeData.Name,
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
func AddTask(t *Task) (*Task, error) {
	// initialize statement write to database
	tx, err := Client.Begin()
	if err != nil {
		return t, err
	}

	// create SQL statement for db writes
	sqlStatement := `INSERT INTO coindrop_tasks (title, type, author, description, token_name, token_allocation, badge) VALUES ($1,$2,$3,$4,$5,$6,$7)`

	// prepare statement
	stmt, err := Client.Prepare(sqlStatement)
	if err != nil {
		return t, err
	}

	defer stmt.Close()

	// execute db write using unique user ID + associated data
	_, err = stmt.Exec(
		&t.Title,
		&t.Type,
		&t.Author,
		&t.Description,
		&t.Token,
		&t.TokenAllocation,
		&t.BadgeData.Name,
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