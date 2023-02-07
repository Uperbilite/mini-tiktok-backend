package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"mini-tiktok-backend/cmd/video/dal"
	"mini-tiktok-backend/cmd/video/rpc"
	"mini-tiktok-backend/kitex_gen/video/videoservice"
	"mini-tiktok-backend/pkg/consts"
	"net"
)

func Init() {
	dal.Init()
	rpc.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(consts.TCP, consts.VideoServiceAddr)
	if err != nil {
		panic(err)
	}

	Init()

	svr := videoservice.NewServer(new(VideoServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.VideoServiceName}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
