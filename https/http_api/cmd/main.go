package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pyihe/go-pkg/https/http_api"
	"github.com/pyihe/go-pkg/tools"
)

type game struct {
}

func (g *game) Handle(r http_api.IRouter) {
	r.GET("/get", http_api.WrapHandler(g.get))
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
	defer s.Stop()

	s.AddHandler(&game{})

	s.Run()

	tools.Wait()
}
