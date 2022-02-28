package api

type Authentication interface {
	Login(username string, password string) (error, AuthData)
}

type AuthData struct {
	UserID uint64
	Token  string
}
