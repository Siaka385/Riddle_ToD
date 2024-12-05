package controllers

import (
	"net/http"

	"socials/views"
)

func Index(w http.ResponseWriter, r *http.Request) {
	views.Render(w, r, "index.html", nil)
}
