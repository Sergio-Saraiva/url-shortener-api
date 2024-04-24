package users

import (
	"net/http"
)

type UsersDelivery interface {
	CreateUser() http.HandlerFunc
	SignIn() http.HandlerFunc
}
