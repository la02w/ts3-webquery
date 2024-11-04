package ts3_webquery

import (
	"encoding/json"
)

func (c *Client) ClientList() (*ClientList, error) {
	url := c.WebQuery + "/1/clientlist?api-key=" + c.APIKey
	body, err := Get(url)
	if err != nil {
		return nil, err
	}
	var response ClientList
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
