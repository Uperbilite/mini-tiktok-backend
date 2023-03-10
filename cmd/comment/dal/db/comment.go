package db

import (
	"context"
	"gorm.io/gorm"
	"mini-tiktok-backend/pkg/consts"
	"strconv"
	"strings"
)

type Comment struct {
	gorm.Model
	UserId     int64  `json:"user_id"`
	VideoId    int64  `json:"video_id"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type Video struct {
	gorm.Model
	AuthorId      int64  `json:"author_id"`
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	Title         string `json:"title"`
	FavoriteCount uint   `json:"favorite_count"`
	CommentCount  uint   `json:"comment_count"`
}

func (c *Comment) TableName() string {
	return consts.CommentTableName
}

// GetVideoKey Key format is "video:{video_id}"
func GetVideoKey(videoId int64) string {
	var res strings.Builder
	res.WriteString("video:")
	res.WriteString(strconv.FormatInt(videoId, 10))
	return res.String()
}

func CreateComment(ctx context.Context, comment *Comment) (*Comment, error) {
	var err error
	db := DB.Begin()

	if err = db.WithContext(ctx).Create(comment).Error; err != nil {
		db.Rollback()
	}

	if err = db.WithContext(ctx).
		Model(&Video{}).
		Where("id = ?", comment.VideoId).
		Update("comment_count", gorm.Expr("comment_count + ?", 1)).
		Error; err != nil {
		db.Rollback()
	}

	db.Commit()

	// delete redis key for consistent
	RDB.HDel(ctx, GetVideoKey(comment.VideoId), consts.CommentCount)

	if err != nil {
		return nil, err
	}
	return comment, nil
}

func DeleteComment(ctx context.Context, id int64) error {
	var videoId int64
	if err := DB.WithContext(ctx).
		Model(&Comment{}).
		Select("video_id").
		Where("id = ?", id).
		Find(&videoId).Error; err != nil {
		return err
	}

	var err error
	db := DB.Begin()

	if err = db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&Comment{}).Error; err != nil {
		db.Rollback()
	}

	if err = db.WithContext(ctx).
		Model(&Video{}).
		Where("id = ?", videoId).
		Update("comment_count", gorm.Expr("comment_count - ?", 1)).
		Error; err != nil {
		db.Rollback()
	}

	db.Commit()

	// delete redis key for consistent
	RDB.HDel(ctx, GetVideoKey(videoId), consts.CommentCount)

	return err
}

func GetCommentsByVideoId(ctx context.Context, videoId int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := DB.WithContext(ctx).
		Model(&Comment{}).
		Where("video_id = ?", videoId).
		Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}
