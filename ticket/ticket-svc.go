package ticket

import (
	"ticket-service/config"
)

type service struct {
	cnf *config.Config
	cache Cache
}

func NewService(cnf *config.Config, cache Cache) Service {
	return &service{
		cnf: cnf,
		cache: cache,
	}
}
