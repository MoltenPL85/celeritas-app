package middleware

import (
	"myapp/data"

	"github.com/moltenpl85/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
