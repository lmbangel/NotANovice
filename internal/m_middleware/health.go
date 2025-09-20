package m_middleware

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"time"

	"github.com/go-ping/ping"
)

type HealthRepository interface {
	CheckHealth(ctx context.Context) (*Health, error)
}

type Health struct {
	Status string    `json:"status"`
	Error  string    `json:"error"`
	State  string    `json:"state"`
	Time   time.Time `json:"time"`
}

type healthRepository struct {
	pinger    *ping.Pinger
	Address   string
	Transport string
}

func NewHealthRepository(pinger *ping.Pinger, Address string, Transport string) HealthRepository {
	return &healthRepository{
		pinger:    pinger,
		Address:   Address,
		Transport: Transport,
	}
}

func (r *healthRepository) CheckHealth(ctx context.Context) (*Health, error) {

	timeout := 3 * time.Second
	conn, err := net.DialTimeout(r.Transport, r.Address, timeout)
	status := "up"
	errMsg := ""
	if err != nil {
		status = "down"
		errMsg = err.Error()
	} else {
		conn.Close()
	}

	h := &Health{
		Status: status,
		State:  "OK",
		Time:   time.Now(),
		Error:  errMsg,
	}
	return h, nil
}

type HealthService struct {
	repo HealthRepository
}

func NewHealthService(repo HealthRepository) *HealthService {
	return &HealthService{repo: repo}
}

func (s *HealthService) CheckHealth(ctx context.Context) (*Health, error) {

	h, err := s.repo.CheckHealth(ctx)
	if err != nil {
		return nil, err
	}
	return h, nil
}

type HealthHandler struct {
	HealthService *HealthService
}

func (h *HealthHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {

	health, err := h.HealthService.CheckHealth(context.Background())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"Message": err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(health)
}
