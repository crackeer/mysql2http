package util

import (
	"strings"

	"github.com/gookit/goutil/strutil"
)

// GenerateModel
//
//	@param database
//	@param table
//	@param fields
//	@return []byte
//	@return error
func GenerateModel(database, table string, fields []TableField) ([]byte, error) {
	return Render("model.tpl", map[string]interface{}{
		"database":          database,
		"table":             table,
		"table_struct_name": strutil.UpperFirst(strutil.Camel(table)),
		"fields":            convertFields(fields),
	})
}

var typeTypeMapping map[string]string = map[string]string{
	"int":      "int64",
	"varchar":  "string",
	"datetime": "time.Time",
}

func getDataType(sqlType string) string {
	for key, value := range typeTypeMapping {
		if strings.Contains(sqlType, key) {
			return value
		}
	}
	return "string"
}

func convertFields(fields []TableField) []map[string]interface{} {
	retData := []map[string]interface{}{}
	for _, item := range fields {
		retData = append(retData, map[string]interface{}{
			"field": item.Field,
			"type":  getDataType(item.Type),
			"name":  strutil.UpperFirst(strutil.Camel(item.Field)),
		})
	}
	return retData
}
