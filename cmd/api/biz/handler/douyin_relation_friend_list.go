// Code generated by hertz generator.

package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"mini-tiktok-backend/cmd/api/biz/mw"
	"mini-tiktok-backend/cmd/api/biz/rpc"
	"mini-tiktok-backend/kitex_gen/relation"
	pkg_consts "mini-tiktok-backend/pkg/consts"
	"mini-tiktok-backend/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	api_relation "mini-tiktok-backend/cmd/api/biz/model/api/relation"
)

// DouyinRelationFriendList .
// @router /douyin/relation/friend/list/ [GET]
func DouyinRelationFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api_relation.DouyinRelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	user, _ := c.Get(pkg_consts.IdentityKey)

	userList, err := rpc.GetFriendList(ctx, &relation.GetFriendListRequest{
		UserId: user.(*mw.User).UserId,
		TargetUserId: req.UserID,
	})
	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}

	SendResponse(c, errno.Success, utils.H{
		"user_list": userList,
	})
}
