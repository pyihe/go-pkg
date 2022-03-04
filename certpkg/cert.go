package certpkg

import (
	"crypto/tls"
	"encoding/pem"
	"errors"
	"io/ioutil"

	"golang.org/x/crypto/pkcs12"
)

func P12ToPem(p12Path string, password string) (cert *tls.Certificate, err error) {
	if p12Path == "" {
		err = errors.New("empty path")
		return
	}
	data, err := ioutil.ReadFile(p12Path)
	if err != nil {
		return nil, err
	}
	blocks, err := pkcs12.ToPEM(data, password)
	if err != nil {
		return nil, err
	}

	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}
	pemCert, err := tls.X509KeyPair(pemData, pemData)
	return &pemCert, err
}
