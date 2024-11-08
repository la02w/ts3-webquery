package ts3_webquery

import (
	"encoding/json"
	"fmt"
)

// Displays a list of channels created on a virtual server including their ID, order, name, etc. The output can be modified using several command options.
//
//	显示在虚拟服务器上创建的频道列表，包括它们的 ID、排序、名称等。输出可以通过使用多个命令选项进行修改。
func (c *Client) ChannelList(data map[string]string) (*Response, error) {
	url := c.WebQuery + "/1/channellist?api-key=" + c.APIKey
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

// Displays detailed configuration information about a channel including ID, topic, description, etc. For detailed information, see Channel Properties.
//
// 显示关于一个频道的详细配置信息，包括 ID、主题、描述等。有关详细信息，请参阅频道属性。
func (c *Client) ChannelInfo(cid string) (*Response, error) {
	url := c.WebQuery + "/1/channelinfo?api-key=" + c.APIKey
	data := map[string]string{
		"cid": cid,
	}
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

// Displays a list of channels matching a given name pattern.
//
// 显示与给定名称模式匹配的频道列表。
func (c *Client) ChannelFind(channelName string) (any, error) {
	return nil, nil
}

// Moves a channel to a new parent channel with the ID cpid. If order is specified, the channel will be sorted right under the channel with the specified ID. If order is set to 0, the channel will be sorted right below the new parent.
//
// 将一个频道移动到一个具有 ID 为 cpid 的新父频道。如果指定了排序，该频道将会被排序到具有指定 ID 的频道的正下方。如果排序设置为 0，该频道将会被排序到新父频道的正下方。
func (c *Client) ChannelMove(cid string, cpid string, order string) (any, error) {
	return nil, nil
}

// Creates a new channel using the given properties and displays its ID. Note that this command accepts multiple properties which means that you're able to specifiy all settings of the new channel at once. For detailed information, see Channel Properties.
//
// 创建一个新频道，使用给定的属性并显示其 ID。请注意，此命令接受多个属性，这意味着你能够一次性指定新频道的所有设置。有关详细信息，请参阅频道属性。
func (c *Client) ChannelCreate(data map[string]string) (*Response, error) {
	url := fmt.Sprintf("%s/1/channelcreate?api-key=%s", c.WebQuery, c.APIKey)
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

// Deletes an existing channel by ID. If force is set to 1, the channel will be deleted even if there are clients within. The clients will be kicked to the default channel with an appropriate reason message.
//
// 删除一个现有频道通过其 ID。如果force设置为 1，即使频道内有客户端，该频道也将被删除。客户端将被踢到默认频道，并伴有适当的理由消息。
func (c *Client) ChannelDelete(cid string, force bool) (any, error) {
	return nil, nil
}

// Changes a channels configuration using given properties. Note that this command accepts multiple properties which means that you're able to change all settings of the channel specified with cid at once. For detailed information, see Channel Properties.
//
// 更改频道的配置，使用给定的属性。请注意，此命令接受多个属性，这意味着你能够一次性更改指定 cid 频道的所有设置。有关详细信息，请参阅频道属性。
func (c *Client) ChannelEdit(data map[string]string) (*Status, error) {
	url := fmt.Sprintf("%s/1/channeledit?api-key=%s", c.WebQuery, c.APIKey)
	body, err := Post(url, c.TimeOut, data)
	if err != nil {
		return nil, err
	}
	var response Status
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
