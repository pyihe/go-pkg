package https

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pyihe/go-pkg/encoding"
	"github.com/pyihe/go-pkg/errors"
)

var (
	ErrInvalidUrl     = errors.New("url must start with 'http'")
	ErrInvalidEncoder = errors.New("invalid encoder")
)

// Get
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

// GetWithObj
func GetWithObj(client *http.Client, url string, encoder encoding.Encoding, obj interface{}) error {
	data, err := Get(client, url)
	if err != nil {
		return err
	}
	err = encoder.Unmarshal(data, obj)
	return err
}

// Post
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

// PostWithObj
func PostWithObj(client *http.Client, url string, contentType string, body io.Reader, encoder encoding.Encoding, v interface{}) error {
	data, err := Post(client, url, contentType, body)
	if err != nil {
		return err
	}

	return encoder.Unmarshal(data, v)
}
