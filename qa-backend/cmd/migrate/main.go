package main

import (
	"log"
	"github.com/MariaRuasCode/docanswers/internal/migrate"
)

func main() {
	if err := migrate.RunMigrations(); err != nil {
		log.Fatal("❌ Migrações falharam: ", err)
	}
}
