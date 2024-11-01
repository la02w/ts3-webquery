package ts3_webquery

import (
	"io"
	"log"
	"net/http"
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
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error fetching URL:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
