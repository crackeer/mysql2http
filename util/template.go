package util

import (
	"embed"

	"github.com/flosch/pongo2/v4"
)

var (
	//go:embed templates/*.tpl
	templateFiles embed.FS
)

// Render
//
//	@param tplName
//	@param data
//	@param naked
//	@return []byte
//	@return error
func Render(tplName string, data map[string]interface{}) ([]byte, error) {
	bytes, err := templateFiles.ReadFile("templates/" + tplName)
	if err != nil {
		return nil, err
	}
	ponTpl, err := pongo2.FromBytes(bytes)
	if err != nil {
		return nil, err
	}
	return ponTpl.ExecuteBytes(pongo2.Context(data))
}
