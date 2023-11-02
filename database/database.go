package database

import (
	"gql_jobportal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	db.Migrator().DropTable((&models.Job{}))
	err = db.Migrator().AutoMigrate(&models.User{}, &models.Company{}, &models.Job{})
	if err != nil {
		return nil, err
	}

	return db, nil

}
