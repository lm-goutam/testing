package handlers

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func DeleteIntgsById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("DELETE FROM intgs WHERE i_id=?;", params["i_id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer result.Close()
}
