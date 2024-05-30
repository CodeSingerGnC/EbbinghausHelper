package cmd

import (
	"fmt"
	"os"
	"time"
)

// 创建一个新的Json文件
func new(args []string) {
	// 获取文件名
	filename := args[0]

	// 检查文件是否已经存在
	filepath, err := checkJsonExists(filename)

	// 如果文件已经存在，打印错误消息并退出程序
	if err == nil {
		fmt.Printf("%s already exists\n", filepath)
		os.Exit(1)
	}

	// 尝试创建新的文件，如果文件已经存在，返回一个错误
	f, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 关闭文件
	defer f.Close()

	data := NewData()
	// TODO：需要更好的方案进行替代
	newItem := Item{
		No: 999,
		Name: "1.两数之和",
		Website: "https://leetcode.cn/problems/two-sum/description/",
		Scheduled: time.Time{},
		Times: 0,
		Extra: "",
	}

	data.AddItem(newItem)
	data.Frequency = 10

	jsonData, _ := data.Marshal()
	f.WriteString(jsonData)

	f.Sync()
}