// Code generated by hertz generator.

package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"

	"github.com/cloudwego/hertz/pkg/app"
	api_user "mini-tiktok-backend/cmd/api/biz/model/api/user"
	"mini-tiktok-backend/cmd/api/biz/mw"
	"mini-tiktok-backend/cmd/api/biz/rpc"
	"mini-tiktok-backend/kitex_gen/user"
)

// DouyinUserRegister .
// @router /douyin/user/register/ [POST]
func DouyinUserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api_user.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}

	err = rpc.CreateUser(context.Background(), &user.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}

	// login after register and return login response message
	mw.JwtMiddleware.LoginHandler(ctx, c)
}
