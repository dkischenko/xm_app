package models

type User struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	PasswordHash string `json:"passwordHash"`
}

type UserRequest struct {
	Name     string `json:"name" validate:"required,alpha"`
	Password string `json:"password" validate:"required"`
}

type UserCreateResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserLoginResponse struct {
	Hash string `json:"hash"`
}
