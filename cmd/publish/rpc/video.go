package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"mini-tiktok-backend/kitex_gen/video"
	"mini-tiktok-backend/kitex_gen/video/videoservice"
	"mini-tiktok-backend/pkg/consts"
	"mini-tiktok-backend/pkg/errno"
)

var videoClient videoservice.Client

func initVideo() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	c, err := videoservice.NewClient(
		consts.VideoServiceName,
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func GetVideos(ctx context.Context, req *video.GetVideosRequest) ([]*video.Video, error) {
	resp, err := videoClient.GetVideos(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.Videos, nil
}