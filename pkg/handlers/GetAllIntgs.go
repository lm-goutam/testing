package handlers

import (
	"encoding/json"
	"net/http"

	model "github.com/lm-goutam/testing/pkg/model"

	_ "github.com/go-sql-driver/mysql"
)

// Access the Data from intgs table

func GetAllIntgs(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("Select * FROM intgs")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer result.Close()
	var i_id []model.IntegrationStatus
	for result.Next() {
		var intgs model.IntegrationStatus
		err := result.Scan(&intgs.IntegrationStatus_id, &intgs.IntegrationStatus_org, &intgs.IntegrationStatus_cms, &intgs.IntegrationStatus_status, &intgs.IntegrationStatus_app, &intgs.App_url, &intgs.Comment)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		i_id = append(i_id, intgs)

	}
	json.NewEncoder(w).Encode(i_id)
}
