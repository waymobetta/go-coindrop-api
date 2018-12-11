package router

import "net/http"

// Route is a struct object containing route info
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
