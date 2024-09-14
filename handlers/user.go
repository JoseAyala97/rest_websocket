package handlers

import (
	"encoding/json"
	"net/http"
	"rest_websocket/models"
	"rest_websocket/repository"
	"rest_websocket/server"

	"github.com/segmentio/ksuid"
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
		//libreria que retorna de forma aleatoria o random un id unico
		id, err := ksuid.NewRandom()
		if err != nil {
			//se retorna error 500 porque si algo falla no es culpa del cliente, sino del servidor
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var user = models.User{
			Email:    request.Email,
			Password: request.Password,
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
