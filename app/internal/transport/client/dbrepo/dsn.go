package dbrepo

import (
	"fmt"
	"github.com/laterius/service_architecture_hw3/app/internal/domain"
)

func Dsn(cfg domain.Db) string {
	//user:pass@tcp(127.0.0.1:3306)/dbname
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)

	if cfg.Extras != "" {
		connString += " " + cfg.Extras
	}
	return connString
}
