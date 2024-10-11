package handlers

import (
	"encoding/json"
	"net/http"
	"rest_websocket/models"
	"rest_websocket/repository"
	"rest_websocket/server"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	HASH_COST = 8
)

// lo que se solicitara para que usuario sea capaz de registrarse en la aplicaicon
type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// se retornara
type SignUpResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Quien estara encargado de manejar la logica de negocio
func SingUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// se crea un objeto de SingUpRequest
		var request = SignUpRequest{}
		//enviando el cuerpo de la peticion -- codificar en elr equest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			//se retorna 400 porque seguramente loe nviado en el body este mal
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), HASH_COST)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//libreria que retorna de forma aleatoria o random un id unico
		id, err := ksuid.NewRandom()
		if err != nil {
			//se retorna error 500 porque si algo falla no es culpa del cliente, sino del servidor
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var user = models.User{
			Email:    request.Email,
			Password: string(hashedPassword),
			ID:       id.String(),
		}
		//se envia contexto del request que se esta procesando
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			// error que sera del lado del servidor
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SignUpResponse{
			Id:    user.ID,
			Email: user.Email,
		})
	}
}
func LoginHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignUpRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := repository.GetUserByEmail(r.Context(), request.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if user == nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		//CompareHashAndPassword primero se envia lo que esta almacenado en db
		//y despues lo que se esta enviando en la peticion
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		claims := models.AppClaims{
			UserId: user.ID,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(2 * time.Hour * 24).Unix(),
			},
		}
		//algoritmo de firmado
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		//generando el token - enviado como respuesta en la peticion
		tokenString, err := token.SignedString([]byte(s.Config().JWTSecret))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//se envia el token
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(LoginResponse{
			Token: tokenString,
		})
	}
}
