package database

import "kenalbatik-be/internal/domain"

func Migrate() {
	err := DB.Migrator().AutoMigrate(
		&domain.Tier{},
		&domain.Province{},
		&domain.Island{},
		&domain.Batik{},
		&domain.User{},
		&domain.Quiz{},
		&domain.UserAnswer{},
	)
	if err != nil {
		panic(err)
	}
}
