package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)



// Posting sending the Data to database intgs table

func PostAllIntgs(w http.ResponseWriter, r *http.Request) {

	//EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	stmt, err := db.Prepare("INSERT INTO `intgs`(`i_org`,`i_cms`,`i_status`,`i_app`,`app_url`,`comment`) VALUES(?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
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

	_, err = stmt.Exec(i_org, i_cms, i_status, i_app, app_url, comment)
	if err != nil {
		panic(err.Error())
	}
	//fmt.Fprintf(w, "Congratulations! Data is added successfully.... \n")
	resp := make(map[string]interface{})
	resp["message"] = "Success"
	resp["status"] = 200
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}
