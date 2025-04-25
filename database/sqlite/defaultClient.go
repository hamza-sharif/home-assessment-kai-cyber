package sqlite

import (
	"gorm.io/gorm"

	"github.com/hamza-sharif/home-assessment-kai-cyber/database"
	"github.com/hamza-sharif/home-assessment-kai-cyber/models"
)

type defaultClient struct {
	conn *gorm.DB
}

func (cli *defaultClient) AddVulnerabilities(vul []*models.Vulnerability) error {

	return nil
}

func (cli *defaultClient) FindVulnerabilities(filter interface{}) ([]*models.Vulnerability, error) {
	return nil, nil
}

func NewDefaultClient() (database.Client, error) {

	return &defaultClient{conn: nil}, nil
}
