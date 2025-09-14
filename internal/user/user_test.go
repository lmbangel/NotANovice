package user

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/lmbangel/_novice/internal/db"
)

func TestGetUser(t *testing.T) {

	t.Run("Get user by its id", func(t *testing.T) {
		Id := int64(1)
		got, _ := GetUserByID(Id)
		want := db.User{
			ID:    1,
			Email: "lmbangel@gmail.com",
		}
		if got.ID != want.ID || got.Email != want.Email {

			g, _ := json.Marshal(got)
			w, _ := json.Marshal(want)
			t.Errorf("Error: got %s, was expecting %s", g, w)
		}
	})
	t.Run("test get all users", func(t *testing.T) {
		got, _ := GetUsers()
		want := []db.User{}
		if reflect.TypeOf(got) != reflect.TypeOf(want) {
			t.Errorf("Error: got %T, was expecting %T", got, want)
		}

	})
}

func TestCreateUser(t *testing.T) {
	t.Run("Test create new user", func(t *testing.T) {

		user := db.User{
			Username: "TestUsername",
			Email:    "username@gmail.com",
		}

		got, err := CreateUser(user)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
		if got.ID == 0 {
			t.Errorf("Error: User Id not found, user not created ")
		}
	})
}
