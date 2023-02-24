package handlers

import (
	"encoding/json"
	"net/http"

	model "github.com/lm-goutam/testing/pkg/model"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllStat(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("Select stat_id,stat_name FROM status")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var stat_id []model.Status
	for result.Next() {
		var status model.Status
		err := result.Scan(&status.Status_id, &status.StatusInfo)
		if err != nil {
			panic(err.Error())
		}
		stat_id = append(stat_id, status)

	}
	json.NewEncoder(w).Encode(stat_id)
}
