package ts3_webquery

import "testing"

func TestGetData(t *testing.T) {
	var c, _ = Login("http://127.0.0.1:10080", "BABc-CoPxT9lBZ5CB6gFSUdAIbbWZu1ZsSbUUj3")
	data := c.Exec("channelinfo").SetData(map[string]string{"cid": "1"}).GetData()
	SaveToJSONFile(data, "data.json")
}
