// Code generated by Validator v0.1.4. DO NOT EDIT.

package relation

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
func (p *Message) IsValid() error {
	return nil
}
func (p *RelationActionRequest) IsValid() error {
	if p.UserId <= int64(0) {
		return fmt.Errorf("field UserId gt rule failed, current value: %v", p.UserId)
	}
	if p.ToUserId <= int64(0) {
		return fmt.Errorf("field ToUserId gt rule failed, current value: %v", p.ToUserId)
	}
	_src := []int32{int32(1), int32(2)}
	var _exist bool
	for _, src := range _src {
		if p.ActionType == int32(src) {
			_exist = true
			break
		}
	}
	if !_exist {
		return fmt.Errorf("field ActionType in rule failed, current value: %v", p.ActionType)
	}
	return nil
}
func (p *RelationActionResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *GetFollowListRequest) IsValid() error {
	if p.UserId <= int64(0) {
		return fmt.Errorf("field UserId gt rule failed, current value: %v", p.UserId)
	}
	if p.TargetUserId <= int64(0) {
		return fmt.Errorf("field TargetUserId gt rule failed, current value: %v", p.TargetUserId)
	}
	return nil
}
func (p *GetFollowListResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *GetFollowerListRequest) IsValid() error {
	if p.UserId <= int64(0) {
		return fmt.Errorf("field UserId gt rule failed, current value: %v", p.UserId)
	}
	if p.TargetUserId <= int64(0) {
		return fmt.Errorf("field TargetUserId gt rule failed, current value: %v", p.TargetUserId)
	}
	return nil
}
func (p *GetFollowerListResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *GetFriendListRequest) IsValid() error {
	if p.UserId <= int64(0) {
		return fmt.Errorf("field UserId gt rule failed, current value: %v", p.UserId)
	}
	if p.TargetUserId <= int64(0) {
		return fmt.Errorf("field TargetUserId gt rule failed, current value: %v", p.TargetUserId)
	}
	return nil
}
func (p *GetFriendListResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *MessageActionRequest) IsValid() error {
	if p.UserId <= int64(0) {
		return fmt.Errorf("field UserId gt rule failed, current value: %v", p.UserId)
	}
	if p.ToUserId <= int64(0) {
		return fmt.Errorf("field ToUserId gt rule failed, current value: %v", p.ToUserId)
	}
	if len(p.Content) < int(1) {
		return fmt.Errorf("field Content min_len rule failed, current value: %d", len(p.Content))
	}
	return nil
}
func (p *MessageActionResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *MessageChatRequest) IsValid() error {
	if p.UserId <= int64(0) {
		return fmt.Errorf("field UserId gt rule failed, current value: %v", p.UserId)
	}
	if p.ToUserId <= int64(0) {
		return fmt.Errorf("field ToUserId gt rule failed, current value: %v", p.ToUserId)
	}
	return nil
}
func (p *MessageChatResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
