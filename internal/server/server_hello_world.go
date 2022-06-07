package server

import (
	"fmt"
	"github.com/google/wire"
	"net/http"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHelloWorldServer, NewGreeter)

type HelloWorldHandler struct{}

func (h HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!"))
}

func NewHelloWorldServer(port int) *http.Server {
	portStr := fmt.Sprintf(":%v", port)
	server := &http.Server{
		Addr:    portStr,
		Handler: HelloWorldHandler{},
	}
	return server
}
