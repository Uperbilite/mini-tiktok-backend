package consts

const (
	UserTableName       = "users"
	VideoTableName      = "videos"
	FavoriteTableName   = "favorites"
	CommentTableName    = "comments"
	FollowTableName     = "follows"
	MessageTableName    = "messages"
	SecretKey           = "secret key"
	IdentityKey         = "id"
	UserIdKey           = "user_id"
	MySQLDefaultDSN     = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	RedisAddr           = "127.0.0.1:6379"
	RedisPassword       = ""
	RedisDB             = 0
	OSSEndPoint         = "oss-cn-beijing.aliyuncs.com"
	AccessKeyId         = "LTAI5tBQzy3pKUKA1juX68Mq"
	AccessKeySecret     = "iNWQt90rnVTZJYENX4EDWZSUfNDtBu"
	OSSBucketName       = "mini-tiktok-backend"
	ObjKeyLen           = 60
	TCP                 = "tcp"
	ExportEndpoint      = ":4317"
	ETCDAddress         = "127.0.0.1:2379"
	ApiServiceName      = "api"
	UserServiceName     = "user"
	PublishServiceName  = "publish"
	FavoriteServiceName = "favorite"
	VideoServiceName    = "video"
	CommentServiceName  = "comment"
	RelationServiceName = "relation"
	FavoriteCount       = "favorite_count"
	CommentCount        = "comment_count"
	ApiServiceAddr      = "0.0.0.0:8080"
	UserServiceAddr     = "127.0.0.1:8090"
	PublishServiceAddr  = "127.0.0.1:8100"
	FavoriteServiceAddr = "127.0.0.1:8110"
	VideoServiceAddr    = "127.0.0.1:8120"
	CommentServiceAddr  = "127.0.0.1:8130"
	RelationServiceAddr = "127.0.0.1:8140"
)
