package util

import (
	"gorm.io/gorm"
)

// Tables
//
//	@param db
//	@return []string
func Tables(db *gorm.DB) []string {
	list := []map[string]interface{}{}
	retData := []string{}
	if err := db.Raw("show tables").Scan(&list).Error; err != nil {
		return retData
	}
	for _, value := range list {
		for _, value := range value {
			retData = append(retData, value.(string))
		}
	}
	return retData
}

// TableField
type TableField struct {
	Field   string `gorm:"Field"`
	Type    string `gorm:"Type"`
	Null    string `gorm:"Null"`
	Key     string `gorm:"Key"`
	Default string `gorm:"Default"`
	Extra   string `gorm:"extra"`
}

// Desc
//
//	@param db
//	@param table
//	@return interface{}
//	@return error
func Desc(db *gorm.DB, table string) ([]TableField, error) {
	list := []TableField{}

	if err := db.Raw("desc " + table).Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}
