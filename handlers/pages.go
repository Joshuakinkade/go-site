package handlers

import "net/http"

func ShowHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/json")
	w.Write([]byte("{\"message\": \"Hello, World!\"}"))
}

func ShowPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/json")
	w.Write([]byte("{\"message\": \"Posts\"}"))
}
