package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	go_homedir "github.com/mitchellh/go-homedir"
)

// 检查并获取基础文件夹路径
func getBaseDir() (string, error) {
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

// 检查用户输入的Json文件名指向的文件是否存在
func checkJsonExists(filename string) (string, error) {
	// 获取基础基础文件夹路径
	dir, err := getBaseDir()
	// 获取失败返回空路径和错误
	if err != nil {
		return "", err
	}

	// 检查输入文件名是否存在.json后缀
	// 不存在后缀则为其补充.json后缀
	if !strings.HasSuffix(filename, ".json") {
		filename = filename + ".json"
	}

	// 拼接文件路径
	file := filepath.Join(dir, filename)

	// 检查文件是否存在
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return file, fmt.Errorf("%s does not exist", file)
	}

	// 返回文件路径和 nil
	return file, nil
}

// 获得文件夹下所有文件的文件句柄
func getJsonFilePaths(dir string) ([]fs.DirEntry, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var jsonFiles []fs.DirEntry
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			jsonFiles = append(jsonFiles, file)
		}
	}

	return jsonFiles, nil
} 
