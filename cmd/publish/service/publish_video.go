package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gofrs/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"image/jpeg"
	"mini-tiktok-backend/cmd/publish/dal/db"
	"mini-tiktok-backend/kitex_gen/publish"
	"mini-tiktok-backend/pkg/consts"
	"os"
	"strings"
)

type PublishVideoService struct {
	ctx context.Context
}

func NewPublishVideoService(ctx context.Context) *PublishVideoService {
	return &PublishVideoService{ctx: ctx}
}

func (s *PublishVideoService) PublishVideo(req *publish.PublishVideoRequest) error {
	// push video to OSS and get url
	client, err := oss.New(consts.OSSEndPoint, consts.AccessKeyId, consts.AccessKeySecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(consts.OSSBucketName)
	if err != nil {
		return err
	}

	videoUid, _ := uuid.NewV4()
	videoUidString := videoUid.String()

	var videoObjectKey strings.Builder
	videoObjectKey.Grow(consts.ObjKeyLen)
	videoObjectKey.WriteString("video/")
	videoObjectKey.WriteString(videoUidString)
	videoObjectKey.WriteString(".mp4")

	err = bucket.PutObject(videoObjectKey.String(), bytes.NewReader(req.Data))
	if err != nil {
		return err
	}

	videoSignedUrl, err := bucket.SignURL(videoObjectKey.String(), oss.HTTPGet, 600)
	if err != nil {
		return err
	}
	cover, _ := GetCoverFromVideo(videoSignedUrl)

	var coverObjectKey strings.Builder
	coverObjectKey.Grow(consts.ObjKeyLen)
	coverObjectKey.WriteString("cover/")
	coverObjectKey.WriteString(videoUidString)
	coverObjectKey.WriteString(".jpg")

	err = bucket.PutObject(coverObjectKey.String(), bytes.NewReader(cover))
	if err != nil {
		return err
	}

	return db.CreateVideo(s.ctx, &db.Video{
		AuthorId: req.UserId,
		PlayURL:  videoObjectKey.String(),
		CoverURL: coverObjectKey.String(),
		Title:    req.Title,
	})

}

func GetCoverFromVideo(videoUrl string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)

	_ = ffmpeg.Input(videoUrl).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()

	img, _, _ := image.Decode(reader)

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), nil
}
