package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-api/app"
	"github.com/waymobetta/go-coindrop-api/auth"
	"github.com/waymobetta/go-coindrop-api/controllers"
	"github.com/waymobetta/go-coindrop-api/db"
	mw "github.com/waymobetta/go-coindrop-api/middleware"
)

// createServer returns a http server for using in tests
func createServer() *httptest.Server {
	service := goa.New("coindrop")

	cognitoRegion := os.Getenv("AWS_COINDROP_COGNITO_REGION")
	cognitoUserPoolID := os.Getenv("AWS_COINDROP_COGNITO_USER_POOL_ID")

	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	app.UseJWTAuthMiddleware(service, mw.Auth(auth.NewAuth(&auth.Config{
		CognitoRegion:     cognitoRegion,
		CognitoUserPoolID: cognitoUserPoolID,
	})))

	dbport, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	dbs := db.NewDB(&db.Config{
		Host:    os.Getenv("POSTGRES_HOST"),
		Port:    dbport,
		User:    os.Getenv("POSTGRES_USER"),
		Pass:    os.Getenv("POSTGRES_PASS"),
		Dbname:  os.Getenv("POSTGRES_DBNAME"),
		SSLMode: os.Getenv("POSTGRES_SSL_MODE"),
	})

	// Mount controllers
	/*
		healthcheckCtrlr := controllers.NewHealthcheckController(service, dbs)
		app.MountHealthcheckController(service, healthcheckCtrlr)

		userCtrlr := controllers.NewUserController(service, dbs)
		app.MountUserController(service, userCtrlr)
	*/

	walletCtrlr := controllers.NewWalletController(service, dbs)
	app.MountWalletController(service, walletCtrlr)

	/*
		tasksCtrlr := controllers.NewTasksController(service, dbs)
		app.MountTasksController(service, tasksCtrlr)
	*/
	/*

		resultsCtrlr := controllers.NewResultsController(service, dbs)
		app.MountResultsController(service, resultsCtrlr)

		quizCtrlr := controllers.NewQuizController(service, dbs)
		app.MountQuizController(service, quizCtrlr)
	*/

	svr := httptest.NewServer(service.Server.Handler)
	return svr
}

func setAuth(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+getAuthToken())
}

// TODO make dynamic
func getUserID() string {
	return "c1718f78-4869-4f2a-a96e-eba0bbdcdd21"
}

// TODO make dynamic
func getAuthToken() string {
	return "eyJraWQiOiJlS3lvdytnb1wvXC9yWmtkbGFhRFNOM25jTTREd0xTdFhibks4TTB5b211aE09IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiJjMTcxOGY3OC00ODY5LTRmMmEtYTk2ZS1lYmEwYmJkY2RkMjEiLCJldmVudF9pZCI6IjA3YzY5ODViLTM5NjAtMTFlOS05ODVjLTI1MmU5YTE3NTE3NiIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE1NTExNDIwMjIsImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC51cy13ZXN0LTIuYW1hem9uYXdzLmNvbVwvdXMtd2VzdC0yX0wwVldGSEVueSIsImV4cCI6MTU1MTE0NTYyMiwiaWF0IjoxNTUxMTQyMDIyLCJqdGkiOiI3ZDg1OGQwNi1iZjMwLTQ3ODEtYThmNC0yMTE3ZTMyM2E3NzciLCJjbGllbnRfaWQiOiI2ZjFzcGI2MzZwdG4wNzRvbjBwZGpnbms4bCIsInVzZXJuYW1lIjoiYzE3MThmNzgtNDg2OS00ZjJhLWE5NmUtZWJhMGJiZGNkZDIxIn0.DykZC1UKAvky2aPrtmXkk7sL0J7IZOpT2P4nIoTznBr5wGn9HVPFnklg0HArXHozhYmJoL3MScX_NYN5JmibE1Hes1wDnj7xFyDvIt3FzAjMfeWTaURmJTfyuaPUO8XpUBonDB4NsqnY7Q5OwUDQ2ICKlri1sZ2_7NPUREHnTk1hpHkoJUaPtt3F-Skjk1BGb1cfHYw_VvcchfH9zWtD5tmnfebPm4PbMvOtqpBAkdDx1eowbrfL-8q_m2NggwLXlCk2YBWna3n5nCKiVBhx3nQ6Wzy0adpF2P7GeIit9YRxZ4neJbgoQz__1nQJoIUm110RNwA_AvoWUCxQ8QoyYQ"
}
