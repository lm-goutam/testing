package handlers

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	mm "github.com/lm-goutam/testing/pkg/model"
)

func GetAllApp(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("Select app_id,app_name FROM app")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var app_id []mm.Application
	for result.Next() {
		var app mm.Application
		err := result.Scan(&app.Application_id, &app.ApplicationInfo)
		if err != nil {
			panic(err.Error())
		}
		app_id = append(app_id, app)

	}
	json.NewEncoder(w).Encode(app_id)
}
