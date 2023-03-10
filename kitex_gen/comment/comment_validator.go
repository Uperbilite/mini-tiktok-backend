// Code generated by Validator v0.1.4. DO NOT EDIT.

package comment

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *BaseResp) IsValid() error {
	return nil
}
func (p *User) IsValid() error {
	return nil
}
func (p *Comment) IsValid() error {
	if p.User != nil {
		if err := p.User.IsValid(); err != nil {
			return fmt.Errorf("filed User not valid, %w", err)
		}
	}
	return nil
}
func (p *CreateCommentRequest) IsValid() error {
	if p.UserId <= int64(0) {
		return fmt.Errorf("field UserId gt rule failed, current value: %v", p.UserId)
	}
	if p.VideoId <= int64(0) {
		return fmt.Errorf("field VideoId gt rule failed, current value: %v", p.VideoId)
	}
	if len(p.Content) < int(1) {
		return fmt.Errorf("field Content min_len rule failed, current value: %d", len(p.Content))
	}
	return nil
}
func (p *CreateCommentResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	if p.Comment != nil {
		if err := p.Comment.IsValid(); err != nil {
			return fmt.Errorf("filed Comment not valid, %w", err)
		}
	}
	return nil
}
func (p *DeleteCommentRequest) IsValid() error {
	if p.CommentId <= int64(0) {
		return fmt.Errorf("field CommentId gt rule failed, current value: %v", p.CommentId)
	}
	return nil
}
func (p *DeleteCommentResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *GetCommentListRequest) IsValid() error {
	if p.UserId <= int64(0) {
		return fmt.Errorf("field UserId gt rule failed, current value: %v", p.UserId)
	}
	if p.VideoId <= int64(0) {
		return fmt.Errorf("field VideoId gt rule failed, current value: %v", p.VideoId)
	}
	return nil
}
func (p *GetCommentListResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
