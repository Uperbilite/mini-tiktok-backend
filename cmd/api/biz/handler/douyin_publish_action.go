// Code generated by hertz generator.

package handler

import (
	"context"
	"fmt"
	"mime/multipart"
	"mini-tiktok-backend/cmd/api/biz/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	api_publish "mini-tiktok-backend/cmd/api/biz/model/api/publish"
	"mini-tiktok-backend/kitex_gen/publish"
)

// Paras 文件类型的参数接收单独定义
type Paras struct {
	data  *multipart.FileHeader `form:"data"`
	token string                `form:"token"`
	title string                `form:"title"`
}

// DouyinPublishAction .
// @router /douyin/publish/action/ [POST]
func DouyinPublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req Paras
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(req.token, req.title, req.data)

	err = rpc.PublishVideo(ctx, &publish.PublishVideoRequest{
		UserId: 0,
		Data:   nil,
		Title:  "",
	})

	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api_publish.DouyinPublishActionResponse)

	c.JSON(consts.StatusOK, resp)
}
