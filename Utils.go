package ts3_webquery

import (
	"encoding/json"
	"os"
	"time"
)

var c, _ = Login("http://127.0.0.1:10080", "BABf9LIMFyscvC4X68S4WBDlGIRO8wR2XupSTni")

func SaveToJSONFile(data interface{}) error {
	// 将结构体序列化为JSON
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	// 将JSON数据写入文件
	return os.WriteFile("test.json", jsonData, 0644)
}

func CheckTimeOut(timeout time.Duration) time.Duration {
	if timeout.Nanoseconds() == 0 {
		return 10 * time.Second
	}
	return timeout
}

func CheckServerID(sid string) string {
	if sid == "" {
		return "1"
	}
	return sid
}
