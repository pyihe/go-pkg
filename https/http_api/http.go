package http_api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pyihe/go-pkg/errors"
	"github.com/pyihe/go-pkg/syncs"
)

const (
	successCode      = http.StatusOK
	defaultErrorCode = 100000
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

func (s *HttpServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}
	s.wg.Wait()
	return nil
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

func IndentedJSON(c *gin.Context, status int, err error, data interface{}) {
	rsp := &response{}
	if err != nil {
		switch err.(type) {
		case *errors.Error:
			e := err.(*errors.Error)
			rsp.Code = e.Code()
			rsp.Message = e.Message()
		default:
			rsp.Code = defaultErrorCode
			rsp.Message = err.Error()
		}
	} else {
		rsp.Code = successCode
		rsp.Message = "SUCCESS"
		rsp.Data = data
	}
	c.IndentedJSON(status, rsp)
}

func Wrap(handler func(*gin.Context) (interface{}, error)) func(*gin.Context) {
	return func(c *gin.Context) {
		if handler != nil {
			result, err := handler(c)
			var status = http.StatusOK
			if err != nil {
				status = http.StatusBadRequest
			}
			IndentedJSON(c, status, err, result)
		}
	}
}
