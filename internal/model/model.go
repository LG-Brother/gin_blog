package model

import (
	"fmt"
	"gin_blog/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Model struct {
	ID        uint32 `gorm:"primary_key" json:"id"`
	CreatedOn uint32 `json:"created_on"`
	CreatedBy string `json:"created_by"`
	UpdatedOn uint32 `json:"update_on"`
	UpdatedBy string `json:"updated_by"`
	DeletedOn uint32 `json:"deleted_on"`
	IsDel     uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	sqDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil
}
