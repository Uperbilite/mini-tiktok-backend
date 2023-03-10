// Code generated by hertz generator.

package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"mini-tiktok-backend/cmd/api/biz/mw"
	"mini-tiktok-backend/cmd/api/biz/rpc"
	"mini-tiktok-backend/kitex_gen/favorite"
	pkg_consts "mini-tiktok-backend/pkg/consts"
	"mini-tiktok-backend/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	api_favorite "mini-tiktok-backend/cmd/api/biz/model/api/favorite"
)

// DouyinFavoriteList .
// @router /douyin/favorite/list/ [GET]
func DouyinFavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api_favorite.DouyinFavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}

	user, _ := c.Get(pkg_consts.IdentityKey)

	videos, err := rpc.GetFavoriteList(ctx, &favorite.GetFavoriteListRequest{
		UserId:       user.(*mw.User).UserId,
		TargetUserId: req.UserID,
	})
	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}

	SendResponse(c, errno.Success, utils.H{
		"video_list": videos,
	})
}
