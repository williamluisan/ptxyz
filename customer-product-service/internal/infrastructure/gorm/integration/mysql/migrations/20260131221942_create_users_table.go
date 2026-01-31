package migrations

import (
	gormModel "ptxyz/customer-product-service/internal/infrastructure/gorm/model"

	"gorm.io/gorm"
)
	

type CreateUsersTable struct {}

func (m CreateUsersTable) ID() string {
	return "20260131221942_create_users_table"
}

func (m CreateUsersTable) Up(db *gorm.DB) error {
	return db.Migrator().CreateTable(&gormModel.UserModel{})
}

func (m CreateUsersTable) Down(db *gorm.DB) error {
	var mdl gormModel.UserModel

	return db.Migrator().DropTable(mdl.TableName())
}