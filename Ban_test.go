package ts3_webquery

import (
	"testing"
)

/*
banadd	| create a ban rule
banclient	| ban a client
bandel	| delete a ban rule
bandelall	| delete all ban rules
banlist	| list ban rules on a virtual server
*/

func TestBanAdd(t *testing.T) {
	data := map[string]string{
		"ip": "1.2.3.4",
	}
	resp, err := c.BanAdd(data)
	if err != nil {
		t.Log(err)
	} else {
		SaveToJSONFile(resp)
	}
}

func TestBanClient(t *testing.T) {
	data := map[string]string{
		"clid": "8",
	}
	resp, err := c.BanClient(data)
	if err != nil {
		t.Log(err)
	} else {
		SaveToJSONFile(resp)
	}
}

func TestBanDel(t *testing.T) {
	resp, err := c.BanDel("2")
	if err != nil {
		t.Log(err)
	} else {
		SaveToJSONFile(resp)
	}
}

func TestBanDelAll(t *testing.T) {
	resp, err := c.BanDellAll()
	if err != nil {
		t.Log(err)
	} else {
		SaveToJSONFile(resp)
	}
}

func TestBanList(t *testing.T) {
	data := map[string]string{
		"-count": "",
	}
	resp, err := c.BanList(data)
	if err != nil {
		t.Log(err)
	} else {
		SaveToJSONFile(resp)
	}
}
