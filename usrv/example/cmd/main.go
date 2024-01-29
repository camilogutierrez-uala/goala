package main

import (
	"github.com/Bancar/goala/usrv"
	"github.com/Bancar/goala/usrv/example/service"
)

func main() {
	HTTPLocalLambda()
}

func HTTPLocalLambda() {
	usrv.LocalHTTP(service.New().Service, true)
}
