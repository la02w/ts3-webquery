package ts3_webquery

import (
	"encoding/json"
	"fmt"
)

// APIKeyAdd 添加一个新的API密钥
//
//	'scope'={'manage'|'write'|'read'}
//	['lifetime'={days}]
//	['cldbid'={clientDBID}]
func (c *Client) APIKeyAdd(data map[string]string) (*Response, error) {
	url := fmt.Sprintf("%s/1/apikeyadd?api-key=%s", c.WebQuery, c.APIKey)
	body, err := Post(url, c.TimeOut, data)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// APIKeyDel 删除一个API密钥
//
//	'id'={APIKeyID}
func (c *Client) APIKeyDel(id string) (*Status, error) {
	url := fmt.Sprintf("%s/1/apikeydel?api-key=%s", c.WebQuery, c.APIKey)
	data := map[string]string{
		"id": id,
	}
	body, err := Post(url, c.TimeOut, data)
	if err != nil {
		return nil, err
	}
	var status Status
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, err
	}
	return &status, nil

}

// APIKeyList 列出所有API密钥
//
//	['cldbid'={clientDBID|'*'}]
//	['start'={offset}]
//	['duration'={limit}]
//	['-count']
func (c *Client) APIKeyList(data map[string]string) (*Response, error) {
	url := fmt.Sprintf("%s/1/apikeylist?api-key=%s", c.WebQuery, c.APIKey)
	body, err := Post(url, c.TimeOut, data)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil

}
