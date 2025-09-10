package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/lmbangel/_novice/internal/db"
	_ "github.com/mattn/go-sqlite3"
)

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "Up",
		"state":  "Healthy",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func HandleGetQuestions(w http.ResponseWriter, r *http.Request) {
	var qs db.Question
	conn, err := sql.Open("sqlite3", "./quiz.db")
	if err != nil {
		panic(err)
	}

	q := db.New(conn)
	qs, err = q.GetQuestion(context.Background(), 1)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Question not found.",
			})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(qs)
}

func HandleAnswersToQuestions(w http.ResponseWriter, r *http.Request) {
	var answer db.Attempt

	if err := json.NewDecoder(r.Body).Decode(&answer); err != nil {
		log.Fatalf("Error getting user answer: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(answer)
}

func HandleCreateNewUser(w http.ResponseWriter, r *http.Request) {
	var u db.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"Error": err.Error()})
	}

	conn, err := sql.Open("sqlite3", "./quiz.db")
	if err != nil {
		panic(err)
	}

	queries := db.New(conn)
	user, err := queries.CreateUser(context.Background(), db.CreateUserParams{
		Username: u.Username,
		Email:    u.Email,
	})

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]string{err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	conn, err := sql.Open("sqlite3", "./quiz.db")
	if err != nil {
		panic(err)
	}

	queries := db.New(conn)

	users, err := queries.GetUsers(context.Background())
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func HandleGetUserByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	conn, err := sql.Open("sqlite3", "./quiz.db")
	if err != nil {
		panic(err)
	}

	queries := db.New(conn)
	users, err := queries.GetUserByID(context.Background(), int64(id))
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "User not found.",
			})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var u db.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		panic(err)
	}
	conn, err := sql.Open("sqlite3", "./quiz.db")
	if err != nil {
		panic(err)
	}

	queries := db.New(conn)

	user, err := queries.GetUserByUsername(context.Background(), u.Username)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "User  not found, click sign up to create account.",
			})
			return
		}
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func main() {
	mux := chi.NewRouter()

	mux.Get("/health", HandleHealthCheck)
	mux.Route("/v1", func(r chi.Router) {
		r.Get("/questions", HandleGetQuestions)
		r.Post("/questions/{id}/answer", HandleAnswersToQuestions)

		r.Post("/login", HandleLogin)
		r.Group(func(r chi.Router) {
			r.Get("/users", HandleGetUsers)
			r.Get("/users/{id}", HandleGetUserByID)
			r.Post("/users", HandleCreateNewUser)
		})
	})

	port := fmt.Sprintf(":%d", 8000)
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	fmt.Printf("Server running on : %s", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("%s", err)
	}

}
