package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	//port to use
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

type Server interface {
	//Config's server
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

// Method return config
func (b *Broker) Config() *Config {
	return b.config
}

// Constructor
func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		//*Broker, error
		return nil, errors.New("port is required")
	}

	if config.JWTSecret == "" {
		//*Broker, error
		return nil, errors.New("secret is required")
	}

	if config.DatabaseUrl == "" {
		//*Broker, error
		return nil, errors.New("database is required")
	}

	// Instance of broker
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}

	return broker, nil
}

// Method Start broker
func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	//Method instance new routers
	b.router = mux.NewRouter()
	binder(b, b.router)
	log.Println("Starting server on port", b.config.Port)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Println("error starting server:", err)
	} else {
		log.Fatalf("server stopped")
	}
}
