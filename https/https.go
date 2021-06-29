package https

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pyihe/go-pkg/encoding"
)

type TLSClient interface {
	Get(url string) ([]byte, error)
	GetWithObj(url string, obj interface{}) error
	Post(url, contentType string, body io.Reader) ([]byte, error)
	PostWithObj(url, contentType string, body io.Reader, obj interface{}) error
	ListenAndServeTLS(serverCrt, serverKey string) error
}

// httpClient
type httpClient struct {
	client      *http.Client
	ca          *x509.CertPool
	certificate tls.Certificate
	encoder     encoding.Encoding
}

// NewTLSClient
func NewTLSClient(caCrt, clientCrt, clientKey string, encoder encoding.Encoding) {
	var err error
	c := &httpClient{
		encoder: encoder,
	}

	c.ca, err = loadCA(caCrt)
	if err != nil {
		panic(err)
	}

	c.certificate, err = tls.LoadX509KeyPair(clientCrt, clientKey)
	if err != nil {
		panic(err)
	}

	c.client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      c.ca,
				Certificates: []tls.Certificate{c.certificate},
			},
		},
	}

}

// ListenAndServeTLS
func (h *httpClient) ListenAndServeTLS(serverCrt, serverKey string) error {
	server := &http.Server{
		Addr: ":443",
		TLSConfig: &tls.Config{
			ClientCAs:  h.ca,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}
	return server.ListenAndServeTLS(serverCrt, serverKey)
}

// Get
func (h *httpClient) Get(url string) ([]byte, error) {
	response, err := h.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

// GetWithObj
func (h *httpClient) GetWithObj(url string, obj interface{}) error {
	data, err := h.Get(url)
	if err != nil {
		return err
	}
	if h.encoder == nil {
		return ErrInvalidEncoder
	}

	return h.encoder.Unmarshal(data, obj)
}

// Post
func (h *httpClient) Post(url, contentType string, body io.Reader) ([]byte, error) {
	response, err := h.client.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

// PostWithObj
func (h *httpClient) PostWithObj(url, contentType string, body io.Reader, obj interface{}) error {
	data, err := h.Post(url, contentType, body)
	if err != nil {
		return err
	}
	if h.encoder == nil {
		return ErrInvalidEncoder
	}
	return h.encoder.Unmarshal(data, obj)
}

// loadCA 加载身份证书
func loadCA(caFile string) (*x509.CertPool, error) {
	p := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, err
	}
	p.AppendCertsFromPEM(ca)
	return p, nil
}
