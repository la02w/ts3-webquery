## ts3-WebQuery

### Install

```bash
go get -u github.com/la02w/ts3-webquery
```

### Install Teamspeak HTTP Server

> Create Query WhiteList

```bash
mkdir -p  $HOME/.config/ts3server
# all ipv4 ipv6
cat > $HOME/.config/ts3server/query_ip_allowlist.txt << EOF
0.0.0.0/0
::/0
EOF
```

```bash
docker run -p 9987:9987/udp -p 10080:10080 \
-e TS3SERVER_LICENSE=accept \
-e TS3SERVER_QUERY_PROTOCOLS=raw,http \
-v $HOME/.config/ts3server/query_ip_allowlist.txt:/var/ts3server/query_ip_allowlist.txt \
--restart=always \
--name teamspeak teamspeak
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
    c, _ := ts3.Login("http://127.0.0.1:10080", "BABf9LIMFyscvC4X68S4WBDlGIRO8wR2XupSTni")
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
}
```

```log
2024/11/06 16:07:20 serverinfo: 0 ok
2024/11/06 16:07:20 channelcreate: 0 ok
2024/11/06 16:07:20 version: Linux 3.13.7 1655727713
2024/11/06 16:07:20 clientdbinfo#nick_name: ServerQuery Guest
```
