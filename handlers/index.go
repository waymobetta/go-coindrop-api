package handlers

import (
	"fmt"
	"net/http"
)

// TEST

// HandleIndex prints test to screen if successful
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "test")
}
