package config

import (
	"dnsServer/models"
)

func Start() {
	models.DomainName{}.Migrate()
}
