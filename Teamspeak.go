package ts3_webquery

import (
	"io"
	"net/http"
	"time"
)

type Client struct {
	WebQuery string
	APIKey   string
}

func Login(w string, a string) (*Client, error) {
	c := &Client{
		WebQuery: w,
		APIKey:   a,
	}
	return c, nil
}

func Get(url string) ([]byte, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
