package ts3_webquery

import "encoding/json"

func (c *Client) Cmd(cmd string) (*Response, error) {
	url := c.WebQuery + "/" + CheckServerID(c.Sid) + "/" + cmd + "?api-key=" + c.APIKey
	body, err := Get(url, CheckTimeOut(c.TimeOut))
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
func (c *Client) ExecCmd(cmd string, data map[string]string) (*Response, error) {
	url := c.WebQuery + "/" + CheckServerID(c.Sid) + "/" + cmd + "?api-key=" + c.APIKey
	body, err := Post(url, CheckTimeOut(c.TimeOut), data)
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
