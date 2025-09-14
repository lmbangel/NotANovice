package attempt

import (
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
