package ts3

import (
	"encoding/json"
	"log"
)

// Displays detailed configuration information about the selected virtual server including unique ID, number of clients online, configuration, etc. For detailed information, see Virtual Server Properties.
//
// 显示关于所选虚拟服务器的详细配置信息，包括唯一 ID、在线客户端数量、配置等。有关详细信息，请参阅虚拟服务器属性。
func (c *Client) ServerInfo() (*ServerInfo, error) {
	url := c.WebQuery + "/1/serverinfo?api-key=" + c.APIKey
	body, err := Get(url)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}
	var response ServerInfo
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return nil, err
	}
	return &response, nil
}
