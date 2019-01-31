package handlers

// TODO:
// create 2 table schemas:
// 1. to store actions (linked to tasks table)
// 2. to store results of action tasks

/*
// ActionsGet handles queries to return all stored tasks
func (h *Handlers) ActionGet(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	// initialize new variable tasks of Tasks struct
	tasks := new(db.Tasks)

	// return slice of structs of all task listings
	_, err := h.db.GetTasks(tasks)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// marshall into JSON
	_, err = json.Marshal(tasks)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, tasks)
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)

	fmt.Printf("Successfully returned information for %d tasks\n\n", len(tasks.Tasks))
}

// ActionAdd adds a single task listing to db
func (h *Handlers) ActionAdd(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	// initialize new variable user of User struct
	task := new(db.Task)

	// add limit for large payload protection
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}
	defer r.Body.Close()

	// unmarshal bytes into user struct
	err = json.Unmarshal(body, task)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// add task action listing to db
	_, err = h.db.AddAction(task)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, "success")
	w.WriteHeader(http.StatusCreated)
	utils.Respond(w, response)

	fmt.Printf("Successfully added task: %s\n\n", task.Title)
}
*/
