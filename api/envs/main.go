package envs

import (
	"github.com/hokauz/go-clean-api/api/pkg/db"
	"github.com/hokauz/go-clean-api/api/pkg/router"
)

// Settings -
type Settings struct {
	DB      *db.Environments
	Router  *router.Environments
	PathAPI string
	// JWT    *jwt.Environments
}

// GetInfo -
func GetInfo() *Settings {
	return &Settings{
		PathAPI: "api/v1",
		DB: &db.Environments{
			Cluster:  "",
			Database: "clean-api",
			User:     "cleanADMIN",
			Password: "asd123A5q3@#4@#45",
			Mode:     "dev",
			Host:     "localhost",
		},
		Router: &router.Environments{Port: ":3000"},
		// JWT: &jwt.Environments{
		// 	Secret:       "",
		// 	TimeDuration: 8, // horas
		// 	Timezone:     "America/Sao_Paulo",
		// },
	}
}
