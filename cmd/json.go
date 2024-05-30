package cmd

import (
	"encoding/json"
	"time"
	// "time"
)

type Item struct {
	No			 int 		`json:"no"`
	Name 		 string 	`json:"name"`
	Website 	 string 	`json:"website"`
	ThisTime	 time.Time 	`json:"thistime"`
	Scheduled    time.Time 	`json:"scheduled"`
	Times 		 int 		`json:"times"`
	Extra		 string 	`json:"extra"`
}

type Data struct {
	ReviewScheme []int 		`json:"reviewscheme"`
	Items		 []Item 	`json:"items"`
	Initialized  bool 		`json:"initialized"`
	Frequency	 int 		`json:"Frequency"`
}

// 初始化 Data
func NewData() Data {
	return Data{
		ReviewScheme: []int{3, 7, 21, 30},
		Items: []Item{},
		Initialized: false,
	}
}

// 将 Data 结构体对象转换为 JSON 字符串
func (d *Data) Marshal() (string, error) {
	jsonData, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// 将 JSON 字符串转换为 Data 结构体对象
func (d *Data) Unmarshal(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), d)
}

// 新增 Item
func (d *Data) AddItem(newItem Item) {
	d.Items = append(d.Items, newItem)
}
