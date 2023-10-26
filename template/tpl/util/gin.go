package util

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Fail(ctx *gin.Context, message string) {
	ctx.PureJSON(http.StatusOK, map[string]interface{}{
		"error": message,
		"data":  nil,
	})
	ctx.Abort()
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.PureJSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  data,
	})
	ctx.Abort()
}

// QuerySetting
type QuerySetting struct {
	OrderBy string   `json:"order_by"`
	Limit   int64    `json:"limit"`
	Fields  []string `json:"fields"`
}

type QueryRequestSetting struct {
	Setting *QuerySetting `json:"_setting"`
}

// ParseQueryRequest
//
//	@param ctx
//	@return map[string]interface{}
//	@return *QuerySetting
//	@return error
func ParseQueryRequest(ctx *gin.Context) (map[string]interface{}, *QuerySetting, error) {
	var (
		queryMap     map[string]interface{} = make(map[string]interface{})
		querySetting *QueryRequestSetting   = &QueryRequestSetting{}
	)
	bytes, err := ctx.GetRawData()
	if err != nil {
		return nil, nil, err
	}
	if err := json.Unmarshal(bytes, &queryMap); err != nil {
		return nil, nil, err
	}

	if err := json.Unmarshal(bytes, querySetting); err != nil {
		return nil, nil, err
	}

	delete(queryMap, "_setting")
	if querySetting.Setting.Limit < 1 {
		querySetting.Setting.Limit = 200
	}
	return queryMap, querySetting.Setting, nil
}

// QuerySetting
type ModifySetting struct {
	Where map[string]interface{} `json:"where"`
}

type ModifyRequestSetting struct {
	Setting *ModifySetting `json:"_setting"`
}

// ParseModifyRequest
//
//	@param ctx
//	@return map[string]interface{}
//	@return *ModifySetting
//	@return error
func ParseModifyRequest(ctx *gin.Context) (map[string]interface{}, *ModifySetting, error) {
	var (
		queryMap      map[string]interface{} = make(map[string]interface{})
		modifySetting *ModifyRequestSetting  = &ModifyRequestSetting{}
	)

	bytes, err := ctx.GetRawData()
	if err != nil {
		return nil, nil, err
	}
	if err := json.Unmarshal(bytes, &queryMap); err != nil {
		return nil, nil, err
	}

	if err := json.Unmarshal(bytes, modifySetting); err != nil {
		return nil, nil, err
	}

	delete(queryMap, "_setting")

	return queryMap, modifySetting.Setting, nil
}

// ParseDeleteRequest
//
//	@param ctx
//	@return map[string]interface{}
//	@return *ModifySetting
//	@return error
func ParseDeleteRequest(ctx *gin.Context) (map[string]interface{}, error) {
	var (
		queryMap map[string]interface{} = make(map[string]interface{})
	)
	if err := ctx.ShouldBindJSON(&queryMap); err != nil {
		return nil, err
	}

	return queryMap, nil
}

// ParseCreateRequest
//
//	@param ctx
//	@return map[string]interface{}
//	@return error
func ParseCreateRequest(ctx *gin.Context, dest interface{}) error {
	return ctx.ShouldBindJSON(&dest)
}
