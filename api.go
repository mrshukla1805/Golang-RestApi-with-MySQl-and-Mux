package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = ""

type Project struct {
	gorm.Model
	ProjectName  string `json:"project"`
	Technologies string `json:"tech"`
	Description  string `json:"description"`
}

func ApiCall() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("DB Connected")
	}
	DB.AutoMigrate(&Project{})
}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var projects []Project
	DB.Find(&projects)
	json.NewEncoder(w).Encode(projects)
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var project Project
	DB.First(&project, params["id"])
	json.NewEncoder(w).Encode(project)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var project Project
	json.NewDecoder(r.Body).Decode(&project)
	DB.Create(&project)
	json.NewEncoder(w).Encode(project)
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var project Project
	DB.First(&project, params["id"])
	json.NewDecoder(r.Body).Decode(&project)
	DB.Save(&project)
	json.NewEncoder(w).Encode(project)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var project Project
	DB.Delete(&project, params["id"])
	json.NewEncoder(w).Encode("Project Deleted")
}
