## ts3-WebQuery

### Install

```bash
go get -u github.com/la02w/ts3-webquery
```

### Use

```go
package main

import (
	"log"
	"time"

	ts3 "github.com/la02w/ts3-webquery"
)

func main() {
	c, _ := ts3.Login("http://127.0.0.1:10080", "BABJGI7p-2EHIr4EB......", 5*time.Second)
    // base func
	resp, _ := c.ServerInfo()
	log.Println("serverinfo:",resp.Status.Code,resp.Status.Message)
	log.Println()
    // Modify and Add Parameters
	data := map[string]string{
		"channel_name":     "channel_name",
		"channel_password": "abc",
        // ....
	}
	resp, _ = c.ChannelCreate(data)
	log.Println("channelcreate:",resp.Status.Code,resp.Status.Message)
    // Use Cmd
    resp, _ = c.Cmd("version")
	for _, body := range *resp.Body {
		log.Println("version:",*body.Platform, *body.Version, *body.Build)
	}
    // Use ExecCmd
	data = map[string]string{
		"cldbid": "2",
	}
	resp, _ = c.ExecCmd("clientdbinfo", data)
	for _, body := range *resp.Body {
		log.Println("clientdbinfo#nick_name:",*body.ClientNickname)
	}
}

```

```log
2024/11/06 16:07:20 serverinfo: 0 ok
2024/11/06 16:07:20 channelcreate: 0 ok
2024/11/06 16:07:20 version: Linux 3.13.7 1655727713
2024/11/06 16:07:20 clientdbinfo#nick_name: ServerQuery Guest
```
