## ts3-WebQuery

### Install

```bash
go get -u github.com/la02w/ts3-webquery/v2
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

import ts3 "github.com/la02w/ts3-webquery/v2"

func main() {
	c, _ := ts3.Login("http://127.0.0.1:10080", "BABc-CoPxT9lBZ5CB6gFSUdAIbbWZu1ZsSbUUj3")
	data := c.Exec("channelinfo").SetData(map[string]string{"cid": "1"}).GetData()
	SaveToJSONFile(data, "data.json")
}

func SaveToJSONFile(data interface{}, filename string) error {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonData, 0644)
}

```
