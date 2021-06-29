package files

import "os"

// MakeNewPath 判断目录是否存在，如果不存在，则新建一个目录
func MakeNewPath(targetPath string) error {
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
