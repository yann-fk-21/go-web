package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/yann-fk-21/todo-platform/types"
)

func TestUserServiceHandlers(t *testing.T) {
	s := &mockUserStore{}
	handler := NewHandler(s)

	t.Run("Echec lorsque le body est vide", func(t *testing.T){
		payload := types.RegisterUserPayload{
			LastName: "Yann",
			FirstName: "James",
			Email: "",
			Password: "",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.registerHandler)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected status %v got %v", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("Should create a ressource", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			LastName: "Yann",
			FirstName: "James",
			Email: "jam@gmail.com",
			Password: "loulou",
		}

		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.registerHandler)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("Expected status %v got %v", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockUserStore struct {}

func (m *mockUserStore)GetUserByEmail(email string)(*types.User, error) {
	return nil, nil
}

func (m *mockUserStore)CreateUser(u types.User) error {
	return nil
}

func (m *mockUserStore)GetUserByID(ID int) (*types.User, error) {
	return nil, nil
}