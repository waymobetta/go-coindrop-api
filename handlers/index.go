package handlers

import (
	"fmt"
	"net/http"
)

// HandleIndex prints test to screen if successful
func (h *Handlers) HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "test")
}
