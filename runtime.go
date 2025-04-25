package kai_cyber_home_assessment

import (
	"github.com/pkg/errors"

	"github.com/hamza-sharif/home-assessment-kai-cyber/database/sqlite"
	"github.com/hamza-sharif/home-assessment-kai-cyber/services"
)

type Runtime struct {
	Svc *services.Service
}

func NewRuntime() (*Runtime, error) {
	db, err := sqlite.NewClient()
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect with database")
	}

	return &Runtime{Svc: services.NewService(db)}, nil
}
