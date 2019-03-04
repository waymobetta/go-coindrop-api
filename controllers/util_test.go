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

	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())
	app.UseJWTAuthMiddleware(service, mw.Auth(auth.NewAuth(&auth.Config{
		CognitoRegion:     cognitoRegion,
		CognitoUserPoolID: cognitoUserPoolID,
	}), dbs))

	// Mount controllers
	/*
		healthcheckCtrlr := controllers.NewHealthcheckController(service, dbs)
		app.MountHealthcheckController(service, healthcheckCtrlr)

		userCtrlr := controllers.NewUserController(service, dbs)
		app.MountUserController(service, userCtrlr)
	*/

	walletCtrlr := controllers.NewWalletController(service, dbs)
	app.MountWalletController(service, walletCtrlr)

	tasksCtrlr := controllers.NewTasksController(service, dbs)
	app.MountTasksController(service, tasksCtrlr)
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
	//return "c1718f78-4869-4f2a-a96e-eba0bbdcdd21"
	return "8ea11ea0-567a-46be-a6c6-6c19adab372c"
}

// TODO make dynamic
func getAuthToken() string {
	return "eyJraWQiOiJlS3lvdytnb1wvXC9yWmtkbGFhRFNOM25jTTREd0xTdFhibks4TTB5b211aE09IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiI4ZWExMWVhMC01NjdhLTQ2YmUtYTZjNi02YzE5YWRhYjM3MmMiLCJldmVudF9pZCI6ImU0YTI4NDRlLTNlMTItMTFlOS05YWU3LTQzNjE3Y2Q0YzVkYSIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE1NTE2NTg2NDcsImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC51cy13ZXN0LTIuYW1hem9uYXdzLmNvbVwvdXMtd2VzdC0yX0wwVldGSEVueSIsImV4cCI6MTU1MTY2MjI0NywiaWF0IjoxNTUxNjU4NjQ3LCJqdGkiOiI4ZGU5MTMzMS1kODE5LTRmN2ItODc4Ni1lNTVlMTYyOWJhMjIiLCJjbGllbnRfaWQiOiI2ZjFzcGI2MzZwdG4wNzRvbjBwZGpnbms4bCIsInVzZXJuYW1lIjoiOGVhMTFlYTAtNTY3YS00NmJlLWE2YzYtNmMxOWFkYWIzNzJjIn0.VxhPCsAzcwq3_ojvIyVpgJQV2C2mz-3Tj1BBr-h2fwrxWeNje-NC95YM5RTNctlGGYJR-YEvkxfSwpNPWp2oWLtsrYIUKSjm5c7OfhxhW_nMlS4ECvdXY88rC-pPVRWUouI0JHytLqFJBRJO--QcZDHBdpKKsJCvQ4Fb8W_N5idkKkGvMjx7dmsdRNT7jAtzJDSKkilSRiJgDHirUtmQ2nqjNLRlkSo5Qq6FK8unbUghMTb_k8dNaP8nX2hVezANhSoQY-PmsADH4O9Ejsp-5Wp5ByFP7S915E2kMiVgA9x5lODevMjWJPwf6DA-vQ0q-ni7DBuedkIHibt8Wotjqg"
}
