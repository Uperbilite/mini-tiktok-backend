// Code generated by hertz generator.

package Publish

import (
	"github.com/cloudwego/hertz/pkg/app"
	"mini-tiktok-backend/cmd/api/biz/mw"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinpublish_ctionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _listMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinpublishlistMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		// use jwt mw
		mw.JwtMiddleware.MiddlewareFunc(),
	}
}
