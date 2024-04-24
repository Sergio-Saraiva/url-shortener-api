package http

import (
	"encoding/json"
	"log"
	"net/http"
	"url-shortener/internal/models"
	"url-shortener/internal/users"
)

type httpUserHandler struct {
	useCase users.UsersUseCases
}

// CreateUser implements users.UsersDelivery.
func (h *httpUserHandler) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createUserReq models.CreateUserRequest
		err := json.NewDecoder(r.Body).Decode(&createUserReq)
		if err != nil {
			log.Printf("Error decoding request: %v", err)
			http.Error(w, "Error decoding request", http.StatusBadRequest)
			return
		}

		if createUserReq.Password != createUserReq.PasswordConfirmation {
			log.Println("Passwords do not match")
			http.Error(w, "Passwords do not match", http.StatusBadRequest)
			return
		}

		userExists, err := h.useCase.GetUser(r.Context(), createUserReq.Email)
		if err != nil {
			log.Printf("Error getting user: %v", err)
			http.Error(w, "Error getting user", http.StatusInternalServerError)
			return
		}

		if userExists != nil {
			log.Println("User already exists")
			http.Error(w, "User already exists", http.StatusBadRequest)
			return
		}

		userCreated, err := h.useCase.CreateUser(r.Context(), &models.User{
			Email:    createUserReq.Email,
			Password: createUserReq.Password,
		})

		if err != nil {
			log.Printf("Error creating user: %v", err)
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		var userCreatedRes models.CreateUserResponse

		userCreatedRes.Username = userCreated.Username
		userCreatedRes.Email = userCreated.Email
		userCreatedRes.ID = userCreated.ID

		log.Println("User created")
		json.NewEncoder(w).Encode(userCreatedRes)
		w.WriteHeader(http.StatusOK)
	}
}

// SignIn implements users.UsersDelivery.
func (h *httpUserHandler) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Signing in")

		var signInReq models.SignInRequest
		err := json.NewDecoder(r.Body).Decode(&signInReq)
		if err != nil {
			log.Printf("Error decoding request: %v", err)
			http.Error(w, "Error decoding request", http.StatusBadRequest)
			return
		}

		user, err := h.useCase.GetUser(r.Context(), signInReq.Email)
		if err != nil {
			log.Printf("Error getting user: %v", err)
			http.Error(w, "Error getting user", http.StatusInternalServerError)
			return
		}

		if user == nil {
			log.Println("User does not exist")
			http.Error(w, "User does not exist", http.StatusBadRequest)
			return
		}

		token, err := h.useCase.SignIn(r.Context(), user, &signInReq)
		if err != nil {
			if err.Error() == "invalid password" {
				log.Println("Invalid password")
				http.Error(w, "Invalid password", http.StatusBadRequest)
				return
			}

			log.Printf("Error signing in: %v", err)
			http.Error(w, "Error signing in", http.StatusInternalServerError)
			return
		}

		var signInRes models.SignInReponse
		signInRes.Token = token
		json.NewEncoder(w).Encode(signInRes)
		w.WriteHeader(http.StatusOK)
	}
}

func NewHttpUserHandler(useCase users.UsersUseCases) users.UsersDelivery {
	return &httpUserHandler{
		useCase: useCase,
	}
}
