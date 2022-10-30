package middleware

import (
	"myapp/data"

	"github.com/moltenpl85/celeritas/v2"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
