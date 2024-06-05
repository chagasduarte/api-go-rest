package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chagasduarte/go-rest-api/database"
	"github.com/chagasduarte/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Word")
}

func TodasContas(w http.ResponseWriter, r *http.Request) {
	var contas []models.Conta
	database.DB.Find(&contas)
	json.NewEncoder(w).Encode(contas)
}

func Conta(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var conta models.Conta
	database.DB.First(&conta, id)
	json.NewEncoder(w).Encode(conta)
}

func NovaConta(w http.ResponseWriter, r *http.Request) {
	var novaConta models.Conta
	json.NewDecoder(r.Body).Decode(&novaConta)
	database.DB.Create(&novaConta)
	json.NewEncoder(w).Encode(novaConta)
}

func DeletaConta(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var conta models.Conta
	database.DB.Delete(&conta, id)
	json.NewEncoder(w).Encode(conta)
}

func EditaConta(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var conta models.Conta
	database.DB.First(&conta, id)
	json.NewDecoder(r.Body).Decode(&conta)
	database.DB.Save(&conta)
	json.NewEncoder(w).Encode(conta)
}
