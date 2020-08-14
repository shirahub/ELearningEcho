package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Username string `json:"username" form:"username" query:"username" name:"username"`
	Password string `json:"password" form:"password" query:"password" name:"password"`
}

type UserData struct {
	UserId   int
	Username string
	Password string
	Fullname string
}

type Claims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

type M map[string]interface{}
