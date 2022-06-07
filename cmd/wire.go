// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/Yuukiiii/geekbangWeek04/internal/biz"
	"github.com/google/wire"
	"net/http"
)

// initApp init application.
func initServer(portHelloWorld int, m string) *http.Server {
	panic(wire.Build(biz.ProviderSet, server.ProviderSet))
}
