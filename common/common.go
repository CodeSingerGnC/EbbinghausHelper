package common

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	go_homedir "github.com/mitchellh/go-homedir"
)

// 检查并获取基础文件夹（~/.ebhelp）路径
func GetBaseDir() (string, error) {
	// 获取当前用户的 Home 目录
	homedir, err := go_homedir.Dir()
	if err != nil {
		return "", err
	}

	// ebhelp 使用的 Json 文件存放的文件夹
	basedir := ".ebhelp"

	// Json 文件存储文件夹
	dir := filepath.Join(homedir, basedir)

	// 分析该文件夹是否存在
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return "", fmt.Errorf("%s directory does not exist", dir)
	}

	// 返回基本文件夹
	return dir, nil
}

// 获得文件夹下所有 Json 文件的文件句柄
func GetJsonFiles(dir string) ([]fs.DirEntry, error) {
	// 寻找 dir 下所有文件夹
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// 切片 jsonFiles 存储 Json 文件的文件句柄
	var jsonFiles []fs.DirEntry
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			jsonFiles = append(jsonFiles, file)
		}
	}

	// 返回存储 Json 文件文件句柄的切片
	return jsonFiles, nil
} 
