package ts3_webquery

import (
	"testing"
)

func TestAPIKeyAdd(t *testing.T) {
	data := map[string]string{
		"scope": "read",
	}
	resp, err := c.APIKeyAdd(data)
	if err != nil {
		t.Log(err)
	} else {
		SaveToJSONFile(resp, "test.json")
	}
}

func TestAPIKeyDels(t *testing.T) {
	resp, err := c.APIKeyDel("22")
	if err != nil {
		t.Log(err)
	} else {
		SaveToJSONFile(resp, "test.json")
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
		SaveToJSONFile(resp, "test.json")
	}
}
