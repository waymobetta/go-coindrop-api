package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"

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
	healthcheckCtrlr := controllers.NewHealthcheckController(service, dbs)
	app.MountHealthcheckController(service, healthcheckCtrlr)

	userCtrlr := controllers.NewUserController(service, dbs)
	app.MountUserController(service, userCtrlr)

	walletCtrlr := controllers.NewWalletController(service, dbs)
	app.MountWalletController(service, walletCtrlr)

	tasksCtrlr := controllers.NewTasksController(service, dbs)
	app.MountTasksController(service, tasksCtrlr)

	resultsCtrlr := controllers.NewResultsController(service, dbs)
	app.MountResultsController(service, resultsCtrlr)

	quizCtrlr := controllers.NewQuizController(service, dbs)
	app.MountQuizController(service, quizCtrlr)

	svr := httptest.NewServer(service.Server.Handler)
	return svr
}

func setAuth(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+getAuthToken())
}

// TODO make dynamic
func getUserID() string {
	//return "c1718f78-4869-4f2a-a96e-eba0bbdcdd21"
	//return "8ea11ea0-567a-46be-a6c6-6c19adab372c"
	return "efb93f03-9802-45b5-9ba0-5c23e68c9d68"
}

// TODO make dynamic
func getAuthToken() string {
	if _, err := os.Stat(".access_token"); os.IsNotExist(err) {
		return ""
	}

	data, err := ioutil.ReadFile(".access_token")
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(data))
}
