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

	tasksCtrlr := controllers.NewTasksController(service, dbs)
	app.MountTasksController(service, tasksCtrlr)
	svr := httptest.NewServer(service.Server.Handler)
	return svr
}

func setAuth(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+getAuthToken())
}

// TODO make dynamic
func getUserID() string {
	return "3d0a93ec-ae4d-4377-9efe-44f2b3c6a7f3"
}

// TODO make dynamic
func getAuthToken() string {
	return "eyJraWQiOiJlS3lvdytnb1wvXC9yWmtkbGFhRFNOM25jTTREd0xTdFhibks4TTB5b211aE09IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiIzZDBhOTNlYy1hZTRkLTQzNzctOWVmZS00NGYyYjNjNmE3ZjMiLCJldmVudF9pZCI6ImE0Y2U5NDFiLTM1OWQtMTFlOS05YjQ5LTVkMDA1MjcxMTYyYiIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE1NTA3Mjg2ODAsImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC51cy13ZXN0LTIuYW1hem9uYXdzLmNvbVwvdXMtd2VzdC0yX0wwVldGSEVueSIsImV4cCI6MTU1MDczMjI4MCwiaWF0IjoxNTUwNzI4NjgwLCJqdGkiOiJiMDgxNmRiNy03NWZkLTQzMDQtYTNmNy02M2I1ZmI2M2UyM2YiLCJjbGllbnRfaWQiOiI2ZjFzcGI2MzZwdG4wNzRvbjBwZGpnbms4bCIsInVzZXJuYW1lIjoiM2QwYTkzZWMtYWU0ZC00Mzc3LTllZmUtNDRmMmIzYzZhN2YzIn0.lJOGY91NcR3feVHaKL2SF2qHYNqnwCLyZhKHUmOoBy-H4CaHdgdyu4auiL-htblqU_ycUpRtcL06lnhlXYmyQv7-kPhxd-BzuxV0nzRHh-8JxR-is8JqTBQL9G0r4zOjDPEr0N9rRwZVAnW9ujzc658I_ruphgAMy73RG834prg8F2aNwgjm1xyM8Gvqlcpje5cHlqmvTOGelh4IiJvONWiO43cxm6w4hwY22qQgKJ55ptBwNRw3aoM0qL35ukw5MNhhdIDMQw2suoc5gjUohRJ_qbPvFNpmXRO_njWSjDQiv_LXoQULNks8zyfyJ8Pmgu5W24CELbhH3meG1WJVTg"
}
