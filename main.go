package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-ping/ping"
	"github.com/lmbangel/_novice/internal/attempt"
	"github.com/lmbangel/_novice/internal/leaderboard"
	"github.com/lmbangel/_novice/internal/m_middleware"
	"github.com/lmbangel/_novice/internal/question"
	"github.com/lmbangel/_novice/internal/quiz"
	"github.com/lmbangel/_novice/internal/user"
	_ "modernc.org/sqlite"
)

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

		// CORS middleware
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

				if r.Method == "OPTIONS" {
					w.WriteHeader(http.StatusOK)
					return
				}

				next.ServeHTTP(w, r)
			})
		})

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

		r.Group(func(r chi.Router) {
			aRepo := attempt.NewSQLiteAttemptRepository(dbConn)
				qRepo := quiz.NewSQLiteQuizRepository(dbConn)
			qService := quiz.NewQuizService(qRepo)
			aService := attempt.NewAttemptService(aRepo, qService)
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
			lRepo := leaderboard.NewSQLiteLeaderBoardRepository(dbConn)
			lService := leaderboard.NewLeaderBoardService(lRepo)
			h := &leaderboard.LeaderBoardHandler{LeaderBoardService: lService}
			r.Get("/leaderboard", h.HandleGetLeaderBoard)
			r.Get("/leaderboard/{id}", h.HandleGetLeaderBoardUserID)
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

		r.Group(func(r chi.Router) {
			qRepo := question.NewSQLiteQuestionRepository(dbConn)
			qService := question.NewQuestionService(qRepo)
			h := &question.QuestionHandler{QuestionService: qService}
			r.Post("/question", h.GenerateQuestion)
			r.Get("/questions", h.GetQuestions)
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
