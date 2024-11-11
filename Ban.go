package ts3_webquery

import (
	"encoding/json"
	"fmt"
)

/*
banadd	| create a ban rule
banclient	| ban a client
bandel	| delete a ban rule
bandelall	| delete all ban rules
banlist	| list ban rules on a virtual server
*/

// banadd	| create a ban rule
//
//	'ip' : regex matching the client ip
//	'name' : regex matching the client nickname
//	'uid' : Uid of client
//	'mytsid' : myTS id of client
//	'time' : integer : time in seconds the ban will be active
//	'banreason' : text describing the reason
//	'lastnickname' :
//
//	must be set: 'ip', 'name', 'uid', or 'mytsid'.
func (c *Client) BanAdd(data map[string]string) (*Response, error) {
	url := fmt.Sprintf("%s/1/banadd?api-key=%s", c.WebQuery, c.APIKey)
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

// banclient	| ban a client
//
//	'-continueonerror' : if specified, will skip to the next clid if an error occurs (rather than stopping)
//	'clid' : integer : id of client to be banned
//	'time' : integer : time in seconds the ban will be active
//	'banreason' : text describing the reaso
func (c *Client) BanClient(data map[string]string) (*Response, error) {
	url := fmt.Sprintf("%s/1/banclient?api-key=%s", c.WebQuery, c.APIKey)
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

// bandel	| delete a ban rule
//
//	'banid' : integer : id of ban to be deleted
func (c *Client) BanDel(banid string) (*Response, error) {
	url := fmt.Sprintf("%s/1/bandel?api-key=%s", c.WebQuery, c.APIKey)
	data := map[string]string{
		"banid": banid,
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

// bandelall	| delete all ban rules
//
//	Deletes all active ban rules from the server
func (c *Client) BanDellAll() (*Status, error) {
	url := fmt.Sprintf("%s/1/bandelall?api-key=%s", c.WebQuery, c.APIKey)
	body, err := Get(url, c.TimeOut)
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

// banlist	| list ban rules on a virtual server
//
//	-count : also return the total number of entries
//	start : integer : skip the first `n` entries
//	duration : integer : only return `n` entries
func (c *Client) BanList(data map[string]string) (*Response, error) {
	url := fmt.Sprintf("%s/1/banlist?api-key=%s", c.WebQuery, c.APIKey)
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
