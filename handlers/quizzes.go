package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/waymobetta/go-coindrop-api/db"
	"github.com/waymobetta/go-coindrop-api/utils"
)

// ResultsPost handles posting quiz results of a specific quiz
func ResultsPost(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	// initialize new copy of Quiz struct to hold quiz info
	quizResults := new(db.Quiz)

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
	err = json.Unmarshal(body, quizResults)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// declare variable and assign to new copy of QuizResults struct
	storedQuizResults := new(db.QuizResults)

	// store quiz information from user in variable to store to db
	storedQuizResults = &db.QuizResults{
		Title:      quizResults.Title,
		AuthUserID: quizResults.AuthUserID,
	}

	// check to see if quiz has already been taken by user
	_, err = db.GetQuizResults(storedQuizResults)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	if storedQuizResults.HasTakenQuiz {
		response = utils.Message(false, "quiz results already submitted")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// TODO:
	// break up calculating quiz results into separate function

	var tried []string
	for i := range quizResults.QuizInfo.QuizData {
		tried = append(tried, quizResults.QuizInfo.QuizData[i].Answer)
	}

	// unmarshals results of a specific quiz
	_, err = db.GetQuiz(quizResults)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// declare empty slice to hold correct answer strings
	var actual []string

	// declare empty integer variable to hold number of correct answers
	var correctCounter int

	// declare empty integer variable to hold number of incorrect answers
	var incorrectCounter int

	// loop over raw quiz results (user's answers) and add
	for j := range quizResults.QuizInfo.QuizData {
		actual = append(actual, quizResults.QuizInfo.QuizData[j].Answer)
	}

	for index := range tried {
		if tried[index] != actual[index] {
			incorrectCounter++
		}
	}

	correctCounter = len(tried) - incorrectCounter

	// store quiz information from user in variable to store to db
	storedQuizResults = &db.QuizResults{
		QuestionsCorrect:   correctCounter,
		QuestionsIncorrect: incorrectCounter,
		HasTakenQuiz:       true,
	}

	// store user's quiz results in db
	_, err = db.StoreQuizResults(storedQuizResults)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, "success")
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)

	fmt.Printf("Successfully stored answers for: %s quiz from user: %s\n\n", quizResults.Title, quizResults.AuthUserID)
}

// ResultsGet handles queries to return all info results of a specific quiz
func ResultsGet(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	// initialize new copy of Quiz struct to hold quiz info
	quizResults := new(db.QuizResults)

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
	err = json.Unmarshal(body, quizResults)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// unmarshals results of a specific quiz
	_, err = db.GetQuizResults(quizResults)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	response = utils.Message(true, quizResults)
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)

	fmt.Printf("Successfully returned information for quiz: %s\n\n", quizResults.Title)
}

// QuizGet handles queries to return all info of a specific quiz
func QuizGet(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	// initialize new variable user of User struct
	quiz := new(db.Quiz)

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
	err = json.Unmarshal(body, quiz)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// unmarshals results of a specific quiz
	_, err = db.GetQuiz(quiz)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// locally overwrite quiz answers to make sure they aren't returned to front-end
	for i := range quiz.QuizInfo.QuizData {
		quiz.QuizInfo.QuizData[i].Answer = ""
	}

	response = utils.Message(true, quiz)
	w.WriteHeader(http.StatusOK)
	utils.Respond(w, response)

	fmt.Printf("Successfully returned information for quiz: %s\n\n", quiz.Title)
}

// QuizAdd adds a single quiz listing to db
func QuizAdd(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	// initialize new variable user of User struct
	quiz := new(db.Quiz)

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
	err = json.Unmarshal(body, quiz)
	if err != nil {
		response = utils.Message(false, err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-type", "application/json")
		utils.Respond(w, response)
		return
	}

	// add quiz data to db
	_, err = db.AddQuiz(quiz)
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

	fmt.Printf("Successfully added quiz: %s\n\n", quiz.Title)
}
