package search

import (
	"encoding/json"
	"os"
)

//注意这里文件路径的写法
const dataFile = "src/sample/data/data.json"

//Feed包含了要处理的种子的信息
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

//读取并解析种子数据文件
func RetrieveFeeds() ([]*Feed, error) {
	//打开文件
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	//预设文件在函数返回时关闭
	defer file.Close()

	//将文件解析到一个Feed类型切片中
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	//这里不检查错误，由调用者来做
	return feeds, err
}
