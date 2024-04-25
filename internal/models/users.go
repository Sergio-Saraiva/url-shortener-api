package models

type User struct {
	ID       string `json:"id" bson:"id,omitempty"`
	Username string `json:"username" bson:"username,omitempty"`
	Password string `json:"password" bson:"password,omitempty"`
	Email    string `json:"email" bson:"email,omitempty"`
	Type     string `json:"type" bson:"type,omitempty"`
}

type CreateUserRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type CreateUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInReponse struct {
	Token string `json:"token"`
}
