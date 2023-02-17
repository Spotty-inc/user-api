package main

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"fmt"
)

func HealthCheck(w http.ResponseWriter, r *http.Request){
	log.Println("Endpoint Hit: Healthcheck")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status OK"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: CreateUser")
	var Newuser User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "wrong data. error: %s", err)
	}
	json.Unmarshal(reqBody, &Newuser)
	if err := Session.Query("INSERT INTO users(id, name, high_score) VALUES(?, ?, ?)",
		Newuser.ID, Newuser.Name, Newuser.High_score).Exec(); err != nil {
		log.Println("Error while inserting")
		log.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
	Conv, _ := json.MarshalIndent(Newuser, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
}

func GetSingleUser(w http.ResponseWriter, r *http.Request){
	log.Println("Endpoint Hit: returnSingleUser")
	UserID := mux.Vars(r)["id"]
	var users []User
	m := map[string]interface{}{}

	iter := Session.Query("SELECT * FROM users WHERE id=?", UserID).Iter()
	for iter.MapScan(m) {
		users = append(users, User{
			ID: m["id"].(int),
			Name: m["name"].(string),
			High_score: m["high_score"].(int),
		})
		m = map[string]interface{}{}
	}

	Conv, _ := json.MarshalIndent(users, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request){
	log.Println("Endpoint Hit: returnAllUsers")
	var users []User
	m := map[string]interface{}{}

	iter := Session.Query("SELECT * FROM users").Iter()
	for iter.MapScan(m) {
		users = append(users, User{
			ID: m["id"].(int),
			Name: m["name"].(string),
			High_score: m["high_score"].(int),
		})
		m = map[string]interface{}{}
	}

	Conv, _ := json.MarshalIndent(users, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
}

func CountAllUsers(w http.ResponseWriter, r *http.Request) {
	var Count string
	err := Session.Query("SELECT count(*) FROM users").Scan(&Count)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", Count)
}
