package server

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/config"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	chi         *chi.Mux
	config      *config.Config
	redisClient *redis.Client
	mongoClient *mongo.Database
}

func NewServer(cfg *config.Config, rdb *redis.Client, mdb *mongo.Database) *Server {
	return &Server{
		chi:         chi.NewRouter(),
		config:      cfg,
		redisClient: rdb,
		mongoClient: mdb,
	}
}

func (s *Server) Run() error {

	err := s.MapHandlers(s.chi)
	if err != nil {
		log.Fatalf("Error mapping handlers: %v", err)
		return err
	}

	log.Printf("Starting server on port %d", s.config.ServerConfig.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", s.config.ServerConfig.Port), s.chi)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
		return err
	}

	return nil
}
