package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lmbangel/_novice/cmd/handlers"
	"github.com/lmbangel/_novice/internal/db"
	"github.com/lmbangel/_novice/pkg/agents"
	_ "modernc.org/sqlite"
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

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var u db.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		panic(err)
	}
	conn, err := sql.Open("sqlite", "./quiz.db")
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

func HandleGetdailyQuiz(w http.ResponseWriter, r *http.Request) {
	sysPrompt := `You are a discipleship leader (a shepherd).
					You want your student to know the foundation and most important things of the faith of Christianity as commented on by the writer of the book of Hebrews (Hebrews 5:12-14).
					Therefore, to test daily if your sheep/disciples are spending time learning and growing in the faith by particularly reading their Bible; have a quiz daily that is used to test the discipled for the day.
					It's a straightforward biblical question every day. Either part of a scripture, or question on context of verse, or question about naming a verse that talks about a specific topic in a specific kind of context.
					For example, How many times is Jesus recorded to be crying in the Bible.`

	userPrompt := "What is the quiz of the day ?"

	o := agents.Ollama{
		Url: "http://localhost:11434/api/generate",
		Request: &agents.Request{
			Model:  "llama3.2:latest",
			Prompt: fmt.Sprintf("%s %s", sysPrompt, userPrompt),
		},
	}
	response := o.Prompt()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]string{response.Response})
}

func main() {
	mux := chi.NewRouter()

	mux.Get("/health", HandleHealthCheck)

	mux.Route("/v1", func(r chi.Router) {
		r.Get("/quiz", HandleGetdailyQuiz)
		r.Post("/login", HandleLogin)

		r.Group(func(r chi.Router) {
			r.Get("/attempts", handlers.HandleGetAttempts)
			r.Get("/attempts/{id}", handlers.HandleGetAttemptByID)
		})

		r.Group(func(r chi.Router) {
			r.Get("/users", handlers.HandleGetUsers)
			r.Get("/users/{id}", handlers.HandleGetUserByID)
			r.Post("/users", handlers.HandleCreateNewUser)
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
