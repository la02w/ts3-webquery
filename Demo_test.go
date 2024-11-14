package ts3_webquery

import (
	"log"
	"testing"
)

func TestDemo(t *testing.T) {
	// No need to import packages for local testing
	c, _ := Login("http://127.0.0.1:10080", "BABf9LIMFyscvC4X68S4WBDlGIRO8wR2XupSTni")
	// set properties
	// c.Sid = "1"
	// c.TimeOut = 5 * time.Second

	// base func
	resp, _ := c.ServerInfo()
	log.Println("serverinfo:", resp.Status.Code, resp.Status.Message)
	// Modify and Add Parameters
	data := map[string]string{
		"channel_name":     "channel_name",
		"channel_password": "abc",
		// ....
		// Indicates whether the channel is permanent or not
		// "channel_flag_permanent":"1",
	}
	resp, _ = c.ChannelCreate(data)
	log.Println("channelcreate:", resp.Status.Code, resp.Status.Message)
	// Use Cmd
	resp, _ = c.Cmd("version")
	for _, body := range resp.Body {
		log.Println("version:", body.Platform, body.Version, body.Build)
	}
	// Use ExecCmd
	data = map[string]string{
		"cldbid": "2",
	}
	resp, _ = c.ExecCmd("clientdbinfo", data)
	for _, body := range resp.Body {
		log.Println("clientdbinfo#nick_name:", body.ClientNickname)
	}
	SaveToJSONFile(resp)
}
