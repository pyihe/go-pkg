package files

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// NewPath 判断目录是否存在，如果不存在，则新建一个目录
func NewPath(targetPath string) error {
	if _, err := os.Stat(targetPath); err != nil {
		if !os.IsExist(err) {
			//创建目录
			if mErr := os.MkdirAll(targetPath, os.ModePerm); mErr != nil {
				return mErr
			}
		}
	}
	return nil
}

// LoadJSONFile 加载JSON文件
func LoadJSONFile(file string, dst interface{}) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, dst)
	return err
}
