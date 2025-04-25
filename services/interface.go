package services

import (
	db "github.com/hamza-sharif/home-assessment-kai-cyber/database"
)

// Service initializes our database instance.
type Service struct {
	db db.Client
}

func NewService(db db.Client) *Service {
	m := &Service{
		db: db,
	}
	return m
}
