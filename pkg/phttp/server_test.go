package phttp

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/panjiawan/go-lib/pkg/plog"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"testing"
)

func TestServer(t *testing.T) {
	s := &fasthttp.Server{
		MaxRequestBodySize: 1024,
	}
	server := fasthttprouter.New()
	server.GET("/test/v1", func(ctx *fasthttp.RequestCtx) {
		fmt.Println("/test/v1 call")
		plog.Info("test", zap.String("path", string(ctx.Path())))
	})
	s.Handler = server.Handler
	s.ListenAndServe(":8889")
}
