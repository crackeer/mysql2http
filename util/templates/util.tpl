package util

import (
    "strings"
    "fmt"
)

func Fail(ctx *gin.Context, message string) {

}


func Success(ctx *gin.Context, data interface{})) {

}


// BuildQuery
//
//	@param query
//	@return string
//	@return []interface{}
func BuildQuery(query map[string]interface{}) (string, []interface{}) {
	queryConditions := []string{}
	params := []interface{}{}
	for key, val := range query {

		if strings.HasPrefix(key, "like@") || strings.HasPrefix(key, "plike@") {
		} else {
			params = append(params, val)
		}

		if !strings.Contains(key, "@") {
			queryConditions = append(queryConditions, fmt.Sprintf("%s in (?)", key))
			continue

		}
		parts := strings.Split(key, "@")
		if len(parts) < 2 {
			queryConditions = append(queryConditions, fmt.Sprintf("%s in (?)", key))
			continue
		}

		switch parts[0] {
		case "gt":
			queryConditions = append(queryConditions, fmt.Sprintf("%s > ?", parts[1]))
		case "gte":
			queryConditions = append(queryConditions, fmt.Sprintf("%s >= ?", parts[1]))
		case "lt":
			queryConditions = append(queryConditions, fmt.Sprintf("%s < ?", parts[1]))
		case "lte":
			queryConditions = append(queryConditions, fmt.Sprintf("%s <= ?", parts[1]))
		case "like":
			queryConditions = append(queryConditions, fmt.Sprintf("%s like '%%%v%%'", parts[1], val))
		case "plike":
			queryConditions = append(queryConditions, fmt.Sprintf("%s like '%v%%'", parts[1], val))
		default:
			queryConditions = append(queryConditions, fmt.Sprintf("%s in (?)", key))
		}
	}
	return strings.Join(queryConditions, " and "), params
} 


func Offset(page, pageSize int64) int64 {
    if page < 1 {
        return 0
    }

    return (page - 1) * pageSize
}

func TotalPage(total, pageSize int64) int64 {
    if total % pageSize > 0 {
        return total / pageSize + 1
    }
    return total / pageSize
}

