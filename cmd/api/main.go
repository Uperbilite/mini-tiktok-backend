// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"mini-tiktok-backend/cmd/api/biz/mw"
	"mini-tiktok-backend/cmd/api/biz/rpc"
)

func Init() {
	rpc.Init()
	mw.InitJWT()
}

func main() {
	Init()

	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	register(h)
	h.Spin()
}