package ts3_webquery

import (
	"encoding/json"
	"log"
)

func (c *Client) ClientList() (*ClientList, error) {
	url := c.WebQuery + "/1/clientlist?api-key=" + c.APIKey
	body, err := Get(url)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}
	var response ClientList
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return nil, err
	}
	return &response, nil
}
