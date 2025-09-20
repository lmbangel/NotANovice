package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-ping/ping"
	"github.com/lmbangel/_novice/internal/attempt"
	"github.com/lmbangel/_novice/internal/db"
	"github.com/lmbangel/_novice/internal/m_middleware"
	"github.com/lmbangel/_novice/internal/quiz"
	"github.com/lmbangel/_novice/internal/user"
	"github.com/lmbangel/_novice/pkg/agents"
	_ "modernc.org/sqlite"
)

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

func setupDatabase() *sql.DB {
	conn, err := sql.Open("sqlite", "./quiz.db")
	if err != nil {
		panic(err)
	}
	return conn
}

func main() {
	mux := chi.NewRouter()

	dbConn := setupDatabase()

	mux.Route("/v1", func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Use(middleware.RealIP)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)

		r.Group(func(r chi.Router) {
			Addr := "8.8.8.8:53"
			//Addr := "127.0.0.1:65534"
			Transport := "tcp"
			p, _ := ping.NewPinger(Addr)
			p.SetPrivileged(false)
			p.Count = 3
			p.Timeout = 5 * time.Second
			hRepo := m_middleware.NewHealthRepository(p, Addr, Transport)
			hService := m_middleware.NewHealthService(hRepo)
			h := &m_middleware.HealthHandler{HealthService: hService}
			r.Get("/health", h.CheckHealth)
		})

		r.Get("/quiz", HandleGetdailyQuiz)
		r.Post("/login", HandleLogin)

		r.Group(func(r chi.Router) {
			aRepo := attempt.NewSQLiteAttemptRepository(dbConn)
			aService := attempt.NewAttemptService(aRepo)
			h := &attempt.AttemptHandler{AttemptService: aService}

			r.Post("/attempts", h.HandleCreateNewAttempt)
			r.Get("/attempts", h.HandleGetAttempts)
			r.Get("/attempts/{id}", h.HandleGetAttemptByID)
			r.Get("/users/{id}/attempts", h.HandleGetAttemptByUserID)
		})

		r.Group(func(r chi.Router) {
			qRepo := quiz.NewSQLiteQuizRepository(dbConn)
			qService := quiz.NewQuizService(qRepo)
			h := &quiz.QuizHandler{QuizService: qService}
			r.Get("/quizes", h.HandleGetQuizes)
			r.Get("/quizes/{id}", h.HandleGetQuizByID)
		})

		r.Group(func(r chi.Router) {
			uRepo := user.NewSQLiteUserRepository(dbConn)
			uService := user.NewUserService(uRepo)
			h := &user.UserHandler{UserService: uService}

			r.Get("/users", h.GetUsers)
			r.Get("/users/{id}", h.GetUserByID)
			r.Post("/users", h.CreateUser)
			r.Put("/users", h.UpdateUser)
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
