package httppkg

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pyihe/go-pkg/serialize"
)

type TLSClient struct {
	client      *http.Client
	ca          *x509.CertPool
	certificate tls.Certificate
	encoder     serialize.Serializer
}

// NewTLSClient 创建TLS客户端
func NewTLSClient(caCrt, clientCrt, clientKey string, encoder serialize.Serializer) *TLSClient {
	var err error
	c := &TLSClient{
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
	return c
}

// ListenAndServeTLS start
func (h *TLSClient) ListenAndServeTLS(serverCrt, serverKey string) error {
	server := &http.Server{
		Addr: ":443",
		TLSConfig: &tls.Config{
			ClientCAs:  h.ca,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}
	return server.ListenAndServeTLS(serverCrt, serverKey)
}

// Get GET请求
func (h *TLSClient) Get(url string) ([]byte, error) {
	response, err := h.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

// GetWithObj GET并反序列化
func (h *TLSClient) GetWithObj(url string, obj interface{}) error {
	data, err := h.Get(url)
	if err != nil {
		return err
	}
	if h.encoder == nil {
		return ErrInvalidEncoder
	}

	return h.encoder.Decode(data, obj)
}

// Post POST请求
func (h *TLSClient) Post(url, contentType string, body io.Reader) ([]byte, error) {
	response, err := h.client.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

// PostWithObj POST请求并反序列化
func (h *TLSClient) PostWithObj(url, contentType string, body io.Reader, obj interface{}) error {
	data, err := h.Post(url, contentType, body)
	if err != nil {
		return err
	}
	if h.encoder == nil {
		return ErrInvalidEncoder
	}
	return h.encoder.Decode(data, obj)
}

func loadCA(caFile string) (*x509.CertPool, error) {
	p := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, err
	}
	p.AppendCertsFromPEM(ca)
	return p, nil
}
