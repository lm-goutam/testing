package handlers

import (
	"encoding/json"
	"net/http"

	model "github.com/lm-goutam/testing/pkg/model"

	_ "github.com/go-sql-driver/mysql"
)

// Access the Data from cms table

func GetAllCms(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("Select cms_id,cms_name FROM cms")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var cms_id []model.Cms
	for result.Next() {
		var cms model.Cms
		err := result.Scan(&cms.Cms_id, &cms.CmsInfo)
		if err != nil {
			panic(err.Error())
		}
		cms_id = append(cms_id, cms)

	}
	json.NewEncoder(w).Encode(cms_id)
}
