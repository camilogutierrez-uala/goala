package main

import (
	"github.com/camilogutierrez-uala/goala/usrv"
	"github.com/camilogutierrez-uala/goala/usrv/example/service"
	"github.com/camilogutierrez-uala/goala/usrv/otel"
)

func main() {

}

func BaseLambda() {
	srv := service.New()
	usrv.LambdaServe(srv, service.Metrics()...)
}

func OtelLambda() {
	srv := service.New()
	if err := otel.UseTrace(); err != nil {
		panic(err)
		return
	}
	defer otel.Shutdown()

	usrv.LambdaOTELServe(srv, otel.LambdaOptions(), service.Metrics()...)
}
