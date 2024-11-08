package ts3_webquery

import (
	"encoding/json"
)

func (c *Client) ClientList(data map[string]string) (*Response, error) {
	url := c.WebQuery + "/1/clientlist?api-key=" + c.APIKey
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
