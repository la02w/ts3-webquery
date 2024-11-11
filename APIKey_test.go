package ts3_webquery

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

var c, _ = Login("http://127.0.0.1:10080", "BABf9LIMFyscvC4X68S4WBDlGIRO8wR2XupSTni", 5*time.Second)

func TestAPIKeyAdd(t *testing.T) {
	data := map[string]string{
		"scope": "read",
	}
	resp, err := c.APIKeyAdd(data)
	if err != nil {
		t.Log(err)
	} else {
		SaveToJSONFile(resp)
	}
}

func TestAPIKeyDels(t *testing.T) {
	resp, err := c.APIKeyDel("22")
	if err != nil {
		t.Log(err)
	} else {
		SaveToJSONFile(resp)
	}
}

func TestAPIKeyList(t *testing.T) {
	data := map[string]string{
		"-count": "",
	}
	resp, err := c.APIKeyList(data)
	if err != nil {
		t.Log(err)
	} else {
		SaveToJSONFile(resp)
	}
}

func SaveToJSONFile(data interface{}) error {
	// 将结构体序列化为JSON
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	// 将JSON数据写入文件
	return os.WriteFile("test.json", jsonData, 0644)
}
