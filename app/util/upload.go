package util

import (
	"gin-user-center/app/config"
	"io"
	"mime/multipart"
	"os"
	"path"
)

/**
 * 保存到文件
 */
func SaveToFile(file *multipart.FileHeader) (fileName string, err error) {
	// 获取后缀名
	ext := path.Ext(file.Filename)
	fileName = Uuid() + ext
	src, err := file.Open()
	if err != nil {
		return fileName, err
	}
	defer src.Close()
	//创建 dst 文件
	out, err := os.Create(config.Conf.File.DirName + fileName)
	if err != nil {
		return fileName, err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	if err != nil {
		return fileName, err
	}
	return fileName, nil
}
