// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videoservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	video "mini-tiktok-backend/kitex_gen/video"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*video.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetVideos": kitex.NewMethodInfo(getVideosHandler, newVideoServiceGetVideosArgs, newVideoServiceGetVideosResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "video",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func getVideosHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceGetVideosArgs)
	realResult := result.(*video.VideoServiceGetVideosResult)
	success, err := handler.(video.VideoService).GetVideos(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceGetVideosArgs() interface{} {
	return video.NewVideoServiceGetVideosArgs()
}

func newVideoServiceGetVideosResult() interface{} {
	return video.NewVideoServiceGetVideosResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetVideos(ctx context.Context, req *video.GetVideosRequest) (r *video.GetVideosResponse, err error) {
	var _args video.VideoServiceGetVideosArgs
	_args.Req = req
	var _result video.VideoServiceGetVideosResult
	if err = p.c.Call(ctx, "GetVideos", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
