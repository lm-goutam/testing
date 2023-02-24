package handlers

import "net/http"

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,GET,GET,GET,GET,OPTIONS,POST")
}
