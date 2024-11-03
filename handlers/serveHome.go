package handlers

import (
	"net/http"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}
