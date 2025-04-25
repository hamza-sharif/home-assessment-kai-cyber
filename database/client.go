package database

import "github.com/hamza-sharif/home-assessment-kai-cyber/models"

type Client interface {
	AddVulnerabilities(vul []*models.Vulnerability) error
	FindVulnerabilities(filter interface{}) ([]*models.Vulnerability, error)
}
