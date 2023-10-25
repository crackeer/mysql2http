package template

import (
	"embed"

	"github.com/flosch/pongo2/v4"
)

var (
	//go:embed tpl/* tpl/*/*
	templateFiles embed.FS
)

const FS_PREFIX = "tpl"

// Render
//
//	@param tplName
//	@param data
//	@param naked
//	@return []byte
//	@return error
func Render(tplName string, data map[string]interface{}) ([]byte, error) {
	bytes, err := templateFiles.ReadFile(FS_PREFIX + "/" + tplName)
	if err != nil {
		return nil, err
	}
	ponTpl, err := pongo2.FromBytes(bytes)
	if err != nil {
		return nil, err
	}
	return ponTpl.ExecuteBytes(pongo2.Context(data))
}

// ReadAllFileList
//
//	@return []string
func ReadAllFileList() []string {
	retData := []string{}
	queue := []string{""}

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]

		dir := FS_PREFIX + "/" + first
		if len(first) < 1 {
			dir = FS_PREFIX
		}

		list, err := templateFiles.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, item := range list {
			path := first + "/" + item.Name()
			if len(first) < 1 {
				path = item.Name()
			}
			if item.IsDir() {
				queue = append(queue, path)
				continue
			} else {
				retData = append(retData, path)
			}
		}
	}
	return retData
}

// Read
//
//	@param filePath
//	@return []byte
//	@return error
func Read(filePath string) ([]byte, error) {
	return templateFiles.ReadFile(FS_PREFIX + "/" + filePath)
}
