package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"mini-tiktok-backend/cmd/favorite/dal"
	"mini-tiktok-backend/kitex_gen/favorite/favoriteservice"
	"mini-tiktok-backend/pkg/consts"
	"net"
)

func Init() {
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(consts.TCP, consts.FavoriteServiceAddr)
	if err != nil {
		panic(err)
	}

	Init()

	svr := favoriteservice.NewServer(new(FavoriteServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.FavoriteServiceName}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
