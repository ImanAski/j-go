package models

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
}
