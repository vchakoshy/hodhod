package api

import "fmt"

type AuthenticationDummy struct{}

func (AuthenticationDummy) Login(username string, password string) (AuthData, error) {
	if username == "demo" && password == "demo" {
		return AuthData{UserID: 1, Token: "1234"}, nil
	}
	return AuthData{}, fmt.Errorf("user not found")
}
