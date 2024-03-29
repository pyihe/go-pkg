package http_api

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/go-pkg/syncs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

const (
	authorization = "Authorization"
	keyJwtJTI     = "jti"
	KeyClientID   = "client_id"
)

type IRouter = gin.IRouter

// response 回复格式
type response struct {
	Code    int32       `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// APIHandler 处理各个HTTP请求的handler
type APIHandler interface {
	Handle(IRouter)
}

// Config 服务配置项
type Config struct {
	SwaggerURL  string // swagger文档地址，如果填写则开启swagger
	Name        string // 服务名称
	Addr        string // 服务地址
	RoutePrefix string // 路由前缀
	CertFile    string // 证书文件
	KeyFile     string // 密钥文件
}

type HttpServer struct {
	config Config          // 服务配置
	engine *gin.Engine     // 路由
	server *http.Server    // HTTP服务
	wg     syncs.WgWrapper // waiter
}

func NewHTTPServer(config Config) *HttpServer {
	engine := gin.Default()
	engine.Use(MidCORS())

	s := &HttpServer{}
	s.config = config
	s.engine = engine
	s.server = &http.Server{
		Addr:    config.Addr,
		Handler: engine,
	}

	if config.SwaggerURL != "" {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL(config.SwaggerURL)))
	}

	return s
}

func (s *HttpServer) Run() {
	s.wg.Wrap(func() {
		var config = s.config
		if config.CertFile != "" && config.KeyFile != "" {
			s.server.ListenAndServeTLS(config.CertFile, config.KeyFile)
		} else {
			s.server.ListenAndServe()
		}
	})
}

func (s *HttpServer) Use(middleware ...gin.HandlerFunc) {
	s.engine.Use(middleware...)
}

func (s *HttpServer) AddHandler(handler APIHandler) {
	engine := s.engine
	prefix := s.config.RoutePrefix
	handler.Handle(engine.Group(prefix))
}

func (s *HttpServer) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}
	s.wg.Wait()
	return nil
}

func JWT(method jwt.SigningMethod, publicKey interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var token *jwt.Token
		var header = c.Request.Header
		var tokenStr = header.Get(authorization)

		if tokenArray := strings.Split(tokenStr, " "); len(tokenArray) == 2 {
			tokenStr = tokenArray[1]
		}
		token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		}, jwt.WithValidMethods([]string{method.Alg()}))

		if err != nil || token == nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			c.Set(KeyClientID, claims[keyJwtJTI])
		}
		c.Next()
	}
}

func MidCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		var header = c.Writer.Header()
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Credentials", "true")
		header.Set("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT,OPTIONS")
		header.Set("Access-Control-Allow-Headers", "Sec-Websocket-Key, Connection, Sec-Websocket-Version, Sec-Websocket-Extensions, Upgrade, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func IndentedJSON(c *gin.Context, err error, data interface{}) {
	status := http.StatusOK
	if err != nil {
		status = http.StatusBadRequest
	}
	rsp := &response{}
	if err != nil {
		switch err.(type) {
		case *errors.Error:
			e := err.(*errors.Error)
			rsp.Code = e.Code()
			rsp.Message = e.Message()
		default:
			rsp.Message = err.Error()
		}
	} else {
		rsp.Message = "SUCCESS"
		rsp.Data = data
	}
	c.IndentedJSON(status, rsp)
}

func WrapHandler(handler func(*gin.Context) (interface{}, error)) func(*gin.Context) {
	return func(c *gin.Context) {
		if handler != nil {
			result, err := handler(c)
			IndentedJSON(c, err, result)
		}
	}
}

// Apply 申请生成Token
type Apply struct {
	ClientID   string            // 客户端ID
	PrivateKey interface{}       // 对应签名类型的私钥
	Expire     time.Duration     // Token有效期
	Method     jwt.SigningMethod // JWT 签名类型
}

// NewToken 生成JSON WEB TOKEN
// 生成普通的Token
func NewToken(apply Apply) (string, error) {
	var now = jwt.TimeFunc()
	var claims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(apply.Expire)),
		NotBefore: jwt.NewNumericDate(now),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        apply.ClientID,
	}
	var token = jwt.NewWithClaims(apply.Method, claims)
	return token.SignedString(apply.PrivateKey)
}
