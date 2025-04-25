package services

import (
	"github.com/hamza-sharif/home-assessment-kai-cyber/models"
)

type Wrapper struct {
	Filters map[string]interface{} `json:"filters"`
}

func (m *Service) QueryVul(filter interface{}) ([]*models.Vulnerability, error) {
	return m.db.FindVulnerabilities(filter)
}
