// Code generated by hertz generator.

package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	favorite "mini-tiktok-backend/cmd/api/biz/model/api/favorite"
)

// DouyinFavoriteAction .
// @router /douyin/favorite/action/ [POST]
func DouyinFavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(favorite.DouyinFavoriteActionResponse)

	c.JSON(consts.StatusOK, resp)
}
