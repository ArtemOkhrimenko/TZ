package users

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func MakeUsersHandler(userSvc Service) http.Handler {
	x := mux.NewRouter()

	x.HandleFunc("/users", ListUsers(userSvc)).Methods(http.MethodGet)
	x.HandleFunc("/users/{id}", GetUser(userSvc)).Methods(http.MethodGet)
	x.HandleFunc("/users/{id}", DeleteUser(userSvc)).Methods(http.MethodDelete)
	x.HandleFunc("/users/{id}", UpsertUser(userSvc)).Methods(http.MethodPut)

	return x
}

func ListUsers(userSvc Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		users := userSvc.ListUsers()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(users)
		if err != nil {
			log.Println(err)
		}
		return
	}
}

func GetUser(userSvc Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)
		userID := v["id"]
		userIDint, err := strconv.Atoi(userID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode(err)
			if err != nil {
				log.Println(err)
			}
			return
		}


		users, err := userSvc.GetUser(userIDint)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			err := json.NewEncoder(w).Encode(err)
			if err != nil {
				log.Println(err)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			log.Println(err)
		}
		return
	}
}

func DeleteUser(userSvc Service)  func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		v := mux.Vars(r)
		userID := v["id"]
		userIDint, err := strconv.Atoi(userID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode(err)
			if err != nil {
				log.Println(err)
			}
			return
		}

		err = userSvc.DeleteUser(userIDint)
		if err != nil{
			w.WriteHeader(http.StatusUnprocessableEntity)
			err := json.NewEncoder(w).Encode(err)
			if err != nil {
				log.Println(err)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		}
}

func UpsertUser(userSvc Service)  func(w http.ResponseWriter, r *http.Request) {

	type request struct {
		Name     string	`json:"name"`
		Email    string	`json:"email"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		rec := &request{}
		if err := json.NewDecoder(r.Body).Decode(rec); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode(err)
			if err != nil {
				log.Println(err)
			}
			return
		}

		v := mux.Vars(r)
		userID := v["id"]
		userIDint, err := strconv.Atoi(userID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode(err)
			if err != nil {
				log.Println(err)
			}
			return
		}

		updatedUser := UserT{
			Name:     rec.Name,
			Email:    rec.Email,
		}

		user, err := userSvc.UpsertUser(userIDint, updatedUser)
		if err != nil{
			w.WriteHeader(http.StatusUnprocessableEntity)
			err := json.NewEncoder(w).Encode(err)
			if err != nil {
				log.Println(err)
			}
			return
		}


		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			log.Println(err)
		}
		return
	}
}