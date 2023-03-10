// Code generated by hertz generator.

package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"mime/multipart"
	"mini-tiktok-backend/cmd/api/biz/mw"
	"mini-tiktok-backend/cmd/api/biz/rpc"
	"mini-tiktok-backend/kitex_gen/publish"
	pkg_consts "mini-tiktok-backend/pkg/consts"
	"mini-tiktok-backend/pkg/errno"
)

// Paras 文件类型的参数接收单独定义
type Paras struct {
	Data  *multipart.FileHeader `form:"data"`
	Token string                `form:"token"`
	Title string                `form:"title"`
}

// DouyinPublishAction .
// @router /douyin/publish/action/ [POST]
func DouyinPublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req Paras
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}

	file, _ := req.Data.Open()
	defer file.Close()
	fileContent, _ := ReadFileContent(file)

	user, _ := c.Get(pkg_consts.IdentityKey)

	err = rpc.PublishVideo(ctx, &publish.PublishVideoRequest{
		UserId: user.(*mw.User).UserId,
		Data:   fileContent,
		Title:  req.Title,
	})

	if err != nil {
		SendResponse(c, err, utils.H{})
		return
	}

	SendResponse(c, errno.Success, utils.H{})
}

func ReadFileContent(file multipart.File) ([]byte, error) {
	const fileChunkSize = 1 * (1 << 20) // read 1MB each time.
	fileContent := make([]byte, 0, 0)

	for {
		fileChunk := make([]byte, fileChunkSize)
		n, _ := file.Read(fileChunk)
		if n == 0 {
			break
		}
		fileChunk = fileChunk[:n]
		fileContent = append(fileContent, fileChunk...)
	}

	return fileContent, nil
}
