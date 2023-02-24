package handlers

import (
	"encoding/json"
	"net/http"

	model "github.com/lm-goutam/testing/pkg/model"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllOrg(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("Select org_id,org_name FROM org")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer result.Close()
	var org_id []model.Organization
	for result.Next() {
		var org model.Organization
		err := result.Scan(&org.Organization_id, &org.OrganizationInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		org_id = append(org_id, org)

	}
	json.NewEncoder(w).Encode(org_id)
}
