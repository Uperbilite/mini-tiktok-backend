// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	api_favorite "mini-tiktok-backend/cmd/api/biz/router/api/favorite"
	api_feed "mini-tiktok-backend/cmd/api/biz/router/api/feed"
	api_publish "mini-tiktok-backend/cmd/api/biz/router/api/publish"
	api_relation "mini-tiktok-backend/cmd/api/biz/router/api/relation"
	api_user "mini-tiktok-backend/cmd/api/biz/router/api/user"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	api_favorite.Register(r)

	api_feed.Register(r)

	api_publish.Register(r)

	api_user.Register(r)

	api_relation.Register(r)

}
