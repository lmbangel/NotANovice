package attempt

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/lmbangel/_novice/internal/db"
)

func TestCreateAttempt(t *testing.T) {
	t.Run("Test createing an attempt", func(t *testing.T) {
		att := db.Attempt{
			UserID: 1,
			QuizID: 3,
			Answer: "A",
		}
		got, err := CreateAttempt(att)

		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		if got.ID == 0 {
			t.Errorf("Error: attempt not created, ID not set")
		}
	})
}

func TestGetAttempts(t *testing.T) {
	t.Run("Test getting attempts", func(t *testing.T) {
		got, err := GetAttempts()
		want := &[]db.Attempt{}

		if err != nil {
			t.Errorf("Error: Could not create an attempt;  %s", err.Error())
		}

		if reflect.TypeOf(got) != reflect.TypeOf(want) {
			t.Errorf("Error: wanted a slice of db.Attempt, i.e []db.Attempt. Received %s", reflect.TypeOf(got))
		}
	})
	t.Run("Test getting attempt by ID", func(t *testing.T) {
		got, err := GetAttemptByID(int64(1))
		want := &db.Attempt{
			ID: 1,
		}

		if err != nil {
			t.Errorf("Error: Could not create an attempt;  %s", err.Error())
		}

		if got.ID != want.ID {
			g, _ := json.Marshal(got)
			w, _ := json.Marshal(want)
			t.Errorf("Error: got %s, was expecting %s", g, w)
		}
	})
}
