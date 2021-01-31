package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	GenerateHTML(w, "hello", "layout", "public_navbar", "top")
}
