package router

import (
	"ms-user/internal/handler"
	"net/http"
)

func NewRouter(userHandler *handler.UserHandler) *http.ServeMux {
	mux := http.NewServeMux()

	// POST /users
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		case http.MethodGet:
			id := r.URL.Query().Get("id")
			if id != "" {
				userHandler.GetUserByID(w, r)
			} else {
				userHandler.GetAllUsers(w, r)
			}
		case http.MethodDelete:
			userHandler.DeleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
