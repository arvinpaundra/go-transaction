package util

import (
	"errors"
	"net/http"
	"strings"
)

var (
	ErrUrlUndefined    = errors.New("url must not empty")
	ErrMethodUndefined = errors.New("method undefined")
)

type Curl struct {
	Method  string
	URL     string
	Headers http.Header
	Body    string
	Err     error
}

func (c *Curl) To(url string) *Curl {
	if url == "" {
		return &Curl{Err: ErrMethodUndefined}
	}
	c.URL = url
	return c
}

func (c *Curl) Do() (*http.Response, error) {
	req, err := http.NewRequest(c.Method, c.URL, strings.NewReader(c.Body))
	if err != nil {
		return nil, err
	}
	req.Header = c.Headers

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return res, err
}

func (c *Curl) SetHeader(name, value string) *Curl {
	if c.Headers == nil {
		c.Headers = make(http.Header)
	}
	c.Headers.Set(name, value)
	return c
}

func (c *Curl) SetBody(body string) *Curl {
	c.Body = body
	return c
}

func (c *Curl) Get() *Curl {
	c.Method = "GET"
	return c
}

func (c *Curl) Post() *Curl {
	c.Method = "POST"
	return c
}

func (c *Curl) Put() *Curl {
	c.Method = "PUT"
	return c
}

func (c *Curl) Delete() *Curl {
	c.Method = "DELETE"
	return c
}
