package main

import (
	"crypto/rsa"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pyihe/go-pkg/https/http_api"
	"github.com/pyihe/go-pkg/tools"
)

const (
	publicData = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuXws30YUeNhHeEqaxvCF
sQy8OrnFpj7nQ36vBVp6W3jid0QhBIOGDJicKPvpNsYiil/a/N2bn2oFpEGWW2UL
9NK5GSlxsJR1lUX+pHCvMqAuUMtkbAUFN+5x81yD5s4IlKeB4o4+5gTPbykTd1Xr
bWElHKci9o/qeLWqefVtahULs8V8lPDXu9aHQRZxhFnK49hFj1NnjjGkRRxqIwVB
B8En6lcbkfbeMvoucLtVjjPJmQeqTBbhUPSb3Bq7XFwS8CalpTGMNq+/8etDfNvr
Mx3fsR7Xl0eLeRZwHzmDSD/e8isU8o09MYJyuf12Gn+Vk57+CGp8wCSbcEBgxIVw
YwIDAQAB
-----END PUBLIC KEY-----
`

	privateData = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC5fCzfRhR42Ed4
SprG8IWxDLw6ucWmPudDfq8FWnpbeOJ3RCEEg4YMmJwo++k2xiKKX9r83ZufagWk
QZZbZQv00rkZKXGwlHWVRf6kcK8yoC5Qy2RsBQU37nHzXIPmzgiUp4Hijj7mBM9v
KRN3VettYSUcpyL2j+p4tap59W1qFQuzxXyU8Ne71odBFnGEWcrj2EWPU2eOMaRF
HGojBUEHwSfqVxuR9t4y+i5wu1WOM8mZB6pMFuFQ9JvcGrtcXBLwJqWlMYw2r7/x
60N82+szHd+xHteXR4t5FnAfOYNIP97yKxTyjT0xgnK5/XYaf5WTnv4IanzAJJtw
QGDEhXBjAgMBAAECggEAXUAkfwuId0ikhcFcFvANBnnUl+GVkILHWZNzAkn+gcZA
dJ13zAEjf2mp+KDNGpB6BP8z5TS0bGys8WtX4BLo8+iMzi2bzp6Ujrtlzd4o9la5
cj0O+496lCf4edTlK0Ah6DpefjvtO07nyoboFnLHrXHNap5MYZDw3EbYsf/FRjL7
Gq23+e0igcnTYJiB+9R+OffwR/KvTp1z3tAVsUf88RlJNtWGUB8AEFP8/Bwy62Gx
Dc8qzoaeoCha+LOwlDuinxcwJlFIsjyGw74QgW85LrfyGh4dyv3IbiLeATzafXRh
vqGvheJ2mXc2DT0cI2T+sfRutc/t1UnbUXt1KZyzgQKBgQDkWL1jq46IG62RhQJB
8hYhJvFzNc/2SdzHesCxVLxA+pLOkAazDEeYyUQ75tGlamW0sGgDcXj9oJp1O3/+
kYHftrxfHOt4+V7//Mo/rwP0M2Y1l6qKCeAcY+4psRu2YQ+HOjqE9+6i8nHBKiou
Thd8Rs96aVmxlDFnvZtsVgBisQKBgQDP8qFENTV11JaRRISMZwt8n2vypFRNDKmg
nnGRE2xpzkpVX6tmRY7SrGgNc95SEO9lMLnQ1DQONv7Yw/up7Yj/tGmw37g8rAKj
zDDs7wHe1dBVw+wlM+4a28JlalVnQYOURMOqzcG35YMBxiCYUim80PdqLloE6FLr
4PZEy17BUwKBgQCMjsc9k/uvcoIbwikKmM7gZ01W4rf5XawGKlx0i7k5skQt3GAT
VKq5tKJI0SMZVG34lGHiRLX6QSLyqMZ31+9+2sgHMBEOLUo5/swr+TpQ1lbDBHHY
eI24TBbtGPT7BbH+RmyBLvB44w38nkzKpg001Y2fRzwL4DGtLvx96k5gcQKBgFW0
l5joIT+OPfxjdAn2Enrrre8UoZYcCPGlPANiMQWuu15Sju8Y7hOQcVZSEihayIA5
Q+x4+Xd+XSz0IY5Y02Uoc4Mtwd5nurLN3sBYhbnVAAfJN1PiAlnZh1aLK+Xhz5xV
dxu3sAbeNk+N3DNLcd5bdg2ySvHI2xxS3M1f0I73AoGAHyIZ5TGbrnOPUtqFjUrr
cnVv9HE5xN8d9XmmTGoR7jcOFaQFPnDX3XnNCMdhdPxIiNmQWXqlLVBWJa6z/wwR
rKPjKpXU/OD6XQ6+ywHyYWT0u7xdR9Mj9ZhAveXbjBuuK6jC2u2gYWslOUJOc5JD
BzB9RGx1kJOYPqhd3nyeNLY=
-----END PRIVATE KEY-----
`
)

var (
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
)

func init() {
	var err error
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(publicData))
	if err != nil {
		panic(err)
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(privateData))
	if err != nil {
		panic(err)
	}
}

type game struct {
}

func (g *game) Handle(r http_api.IRouter) {
	r.GET("/login", http_api.WrapHandler(g.login))

	group := r.Group("", http_api.JWT(jwt.SigningMethodRS256, publicKey))
	{
		group.GET("/get", http_api.WrapHandler(g.get))
	}
}

func (g *game) login(c *gin.Context) (interface{}, error) {
	token, err := http_api.Token(jwt.SigningMethodRS256, privateKey, 20*time.Second)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (g *game) get(c *gin.Context) (interface{}, error) {
	//time.Sleep(10 * time.Second)
	return 100, nil
}

func main() {
	config := http_api.Config{
		Name:        "game",
		Addr:        ":8888",
		RoutePrefix: "",
		CertFile:    "",
		KeyFile:     "",
	}
	s := http_api.NewHTTPServer(config)
	defer s.Close()

	s.AddHandler(&game{})

	s.Run()

	tools.Wait()
}
