package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func UpdateIntgsById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("UPDATE intgs SET i_org=?,i_cms=?,i_status=?,i_app=?,app_url=?,comment=? WHERE i_id=?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	keyVal := make(map[string]int)
	json.Unmarshal(body, &keyVal)
	i_org := keyVal["i_org"]
	i_cms := keyVal["i_cms"]
	i_status := keyVal["i_status"]
	keyVal1 := make(map[string]string)
	json.Unmarshal([]byte(body), &keyVal1)
	app_url := keyVal1["app_url"]
	comment := keyVal1["comment"]
	i_app := keyVal1["i_app"]

	_, err = stmt.Exec(i_org, i_cms, i_status, i_app, app_url, comment, params["i_id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Congratulations! Data is updated successfully.... \n")
	resp := make(map[string]interface{})
	resp["message"] = "Success"
	resp["status"] = 200
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(jsonResp)
	return
}
