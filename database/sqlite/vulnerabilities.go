package sqlite

import (
	"fmt"
	"github.com/hamza-sharif/home-assessment-kai-cyber/models"
)

func (cli *client) AddVulnerabilities(vul []*models.Vulnerability) error {
	result := cli.conn.Create(&vul)
	if result.Error != nil {
		fmt.Println("Insert error:", result.Error)
		return result.Error
	}

	return nil
}

func (cli *client) FindVulnerabilities(filter interface{}) ([]*models.Vulnerability, error) {
	var vals []*models.Vulnerability
	result := cli.conn.Find(&vals, filter)

	return vals, result.Error
}
