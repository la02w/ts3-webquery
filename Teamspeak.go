package ts3_webquery

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Clinet struct {
	Url     string
	Key     string
	Sid     string
	Commond string
	TimeOut time.Duration
	Data    map[string]string
}

// 登录
func Login(url string, key string) *Clinet {
	c := &Clinet{
		Url: url,
		Key: key,
	}
	c.Sid = "1"
	c.TimeOut = 5 * time.Second
	c.Commond = "version"
	return c
}

// 测试连接
func (c *Clinet) Test() *Result {
	client := &http.Client{
		Timeout: c.TimeOut,
	}
	resp, err := client.Get(c.Url)
	if err != nil {
		return &Result{
			Code:    502,
			Message: "Error: Client.Timeout exceeded while awaiting headers",
		}
	}
	defer resp.Body.Close()
	response, _ := c.Exec("version").GetData()
	if response.Status.Code != 0 {
		return &response.Status
	}
	return nil
}

// 设置SID
func (c *Clinet) SetSid(sid string) *Clinet {
	c.Sid = sid
	return c
}

func (c *Clinet) SetTimeOut(time time.Duration) *Clinet {
	c.TimeOut = time
	return c
}

// 查询命令
func (c *Clinet) Exec(commond string) *Clinet {
	c.Commond = commond
	return c
}

// 设置查询配置
func (c *Clinet) SetData(data map[string]string) *Clinet {
	c.Data = data
	return c
}

// 创建HTTP请求获取数据
func (c *Clinet) GetData() (*Response, *Result) {
	reader, _ := json.Marshal(c.Data)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/%s", c.Url, c.Sid, c.Commond), strings.NewReader(string(reader)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, &Result{
			Code:    502,
			Message: err.Error(),
		}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.Key)
	client := &http.Client{
		Timeout: c.TimeOut,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, &Result{
			Code:    502,
			Message: err.Error(),
		}
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, &Result{
			Code:    502,
			Message: err.Error(),
		}
	}
	var response Response
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, &Result{
			Code:    502,
			Message: err.Error(),
		}
	}
	return &response, nil
}
