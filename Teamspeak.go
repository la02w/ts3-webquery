package ts3_webquery

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	Sid      string
	WebQuery string
	APIKey   string
	TimeOut  time.Duration
}

func Login(webquery string, apikey string) (*Client, error) {
	c := &Client{
		WebQuery: webquery,
		APIKey:   apikey,
	}
	c.Sid = "1"
	c.TimeOut = 10 * time.Second
	return c, nil
}

func Get(url string, timeout time.Duration) ([]byte, error) {
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func Post(url string, timeout time.Duration, body map[string]string) ([]byte, error) {
	client := http.Client{
		Timeout: timeout,
	}
	reader, _ := json.Marshal(body)
	resp, err := client.Post(url, "application/json", strings.NewReader(string(reader)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
