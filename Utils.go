package ts3_webquery

import (
	"encoding/json"
	"os"
)

var c, _ = Login("http://127.0.0.1:10080", "BABf9LIMFyscvC4X68S4WBDlGIRO8wR2XupSTni")

func SaveToJSONFile(data interface{}, filename string) error {
	// 将结构体序列化为JSON
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	// 将JSON数据写入文件
	return os.WriteFile(filename, jsonData, 0644)
}
