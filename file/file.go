package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetFileExist(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		// path/to/whatever does not exist
		fmt.Printf("没有找到文件：%s\r\n", filepath)
		return false
	}
	return true
}

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth string, suffix []string) (files []string, err error) {
	files = make([]string, 0, 30)
	for i := 0; i < len(suffix); i++ {
		suffix[i] = strings.ToUpper(suffix[i]) //忽略后缀匹配的大小写
	}
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录

		if fi.IsDir() { // 忽略目录
			return nil
		}
		for _, v := range suffix {
			if strings.HasSuffix(strings.ToUpper(fi.Name()), v) {
				files = append(files, filename)
				break
			}
		}
		return nil
	})
	return files, err
}
