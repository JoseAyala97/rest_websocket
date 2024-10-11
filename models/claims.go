package models

import "github.com/golang-jwt/jwt"

// data codificada en el token
type AppClaims struct {
	UserId             string `json:"userId"`
	jwt.StandardClaims        // se incluye para que el token sea valido
	//composicion sobre herencia
}
