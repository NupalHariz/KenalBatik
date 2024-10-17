package database

import "kenalbatik-be/internal/domain"

func Migrate() {
	err := DB.Migrator().AutoMigrate(
		&domain.Island{},
		&domain.Province{},
		&domain.Batik{},
		&domain.User{},
		&domain.Tier{},
		&domain.Quiz{},
		&domain.UserAnswer{},
	)
	if err != nil {
		panic(err)
	}
}
