package https

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/go-pkg/serialize"
)

var (
	ErrInvalidUrl     = errors.New("url must start with 'http'")
	ErrInvalidEncoder = errors.New("invalid encoder")
)

// Get 发起http get请求
func Get(client *http.Client, url string) ([]byte, error) {
	if url == "" || !strings.HasPrefix(url, "http") {
		return nil, ErrInvalidUrl
	}
	if client == nil {
		client = http.DefaultClient
	}

	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

// GetWithObj 发起get请求，并将结果反序列化到指定obj中
func GetWithObj(client *http.Client, url string, encoder serialize.Codec, obj interface{}) error {
	data, err := Get(client, url)
	if err != nil {
		return err
	}
	err = encoder.Unmarshal(data, obj)
	return err
}

// Post 发起POST请求
func Post(client *http.Client, url string, contentType string, body io.Reader) ([]byte, error) {
	if url == "" || !strings.HasPrefix(url, "http") {
		return nil, ErrInvalidUrl
	}
	if client == nil {
		client = http.DefaultClient
	}
	response, err := client.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

// PostWithObj 发起POST请求并将结果发序列化到指定obj
func PostWithObj(client *http.Client, url string, contentType string, body io.Reader, encoder serialize.Codec, v interface{}) error {
	data, err := Post(client, url, contentType, body)
	if err != nil {
		return err
	}
	return encoder.Unmarshal(data, v)
}
