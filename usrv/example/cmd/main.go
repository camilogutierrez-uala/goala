package main

import (
	"github.com/camilogutierrez-uala/goala/usrv"
	"github.com/camilogutierrez-uala/goala/usrv/example/service"
)

func main() {
	HTTPLocalLambda()
}

func HTTPLocalLambda() {
	usrv.LocalHTTP(service.New().Service, true)
}
