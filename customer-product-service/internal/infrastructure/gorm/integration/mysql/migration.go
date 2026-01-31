package mysql

import (
	gormMysqlMigrations "ptxyz/customer-product-service/internal/infrastructure/gorm/integration/mysql/migrations"

	"gorm.io/gorm"
)	

type Migration interface {
    ID() string
    Up(db *gorm.DB) error
    Down(db *gorm.DB) error
}

// list all models
func Models() []Migration{
	return []Migration{
		gormMysqlMigrations.CreateUsersTable{},
	}
}
