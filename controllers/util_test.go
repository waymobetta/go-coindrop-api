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
	return "c1718f78-4869-4f2a-a96e-eba0bbdcdd21"
}

// TODO make dynamic
func getAuthToken() string {
	return "eyJraWQiOiJlS3lvdytnb1wvXC9yWmtkbGFhRFNOM25jTTREd0xTdFhibks4TTB5b211aE09IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiJjMTcxOGY3OC00ODY5LTRmMmEtYTk2ZS1lYmEwYmJkY2RkMjEiLCJldmVudF9pZCI6ImUwNmM3YjE1LTM1OGMtMTFlOS05ZjE1LWJiMmFhZTEzNjVmMyIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE1NTA3MjE0NzgsImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC51cy13ZXN0LTIuYW1hem9uYXdzLmNvbVwvdXMtd2VzdC0yX0wwVldGSEVueSIsImV4cCI6MTU1MDcyNTA3OCwiaWF0IjoxNTUwNzIxNDc5LCJqdGkiOiIxOWNjY2FhYi00MGFkLTQ5NTMtOTk3Ny1mNmUwZDMxOTc4ODQiLCJjbGllbnRfaWQiOiI2ZjFzcGI2MzZwdG4wNzRvbjBwZGpnbms4bCIsInVzZXJuYW1lIjoiYzE3MThmNzgtNDg2OS00ZjJhLWE5NmUtZWJhMGJiZGNkZDIxIn0.HCJVeK0yl4LYkipuuII8eyElrWByHHzyIUpQDdxk1EtsOcYjQekn3rVT8lhj9766CbVrIeLjBGAMAuA7Pt4iy6yGbgl8w0L4yNyYj89Cg6ltyAQeTYjV0PEOlfwnDVsPfkODgLwYhiOX4f-7SOnxY7-v-wYQGdARsT68XRe8QYGXkx56WEz2BkvPhQ6gGPf0tWb-5h0W6YI7lowtWNNI7dpz5usNfVrw7_bO0x0HFRgPQFn3vLHiUaTSxwQQmEXHy_gozqHBQwY9l4oV4fAsUFCOO3DEqmLQWq4Uo48IkbV6_TVT7OtH2QaSkuTgkEjWISxIEFCNhtAuTVHvn1lHSg"
}
