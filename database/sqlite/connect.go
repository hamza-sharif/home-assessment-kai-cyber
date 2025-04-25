package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"

	"github.com/pkg/errors"

	"github.com/hamza-sharif/home-assessment-kai-cyber/database"
	"github.com/hamza-sharif/home-assessment-kai-cyber/models"
)

type client struct {
	conn *gorm.DB
}

func NewClient() (database.Client, error) {
	conn, err := gorm.Open(sqlite.Open("data/mydb.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	if err = conn.AutoMigrate(&models.Vulnerability{}); err != nil {
		return nil, errors.Wrap(err, "failed to create tables")
	}

	return &client{conn: conn}, err
}
