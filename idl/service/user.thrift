namespace go user

enum ErrCode {
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
}

struct BaseResp {
    1: required i64 status_code
    2: required string status_msg
    3: required i64 service_time
}

struct User {
    1: required i64 id
    2: required string name
    3: required i64 follow_count
    4: required i64 follower_count
    5: required bool is_follow
}

struct CheckUserRequest {
    1: required string username (vt.min_size = "1")
    2: required string password (vt.min_size = "1")
}

struct CheckUserResponse {
    1: required BaseResp base_resp
    2: required i64 user_id
}

struct CreateUserRequest {
    1: required string username (vt.min_size = "1")
    2: required string password (vt.min_size = "1")
}

struct CreateUserResponse {
    1: required BaseResp base_resp
}

struct QueryUserRequest {
    1: required i64 user_id (vt.gt = "0")
    2: required i64 target_user_id (vt.gt = "0")
}

struct QueryUserResponse {
    1: required BaseResp base_resp
    2: required User user
}

service UserService {
    CheckUserResponse CheckUser(1: CheckUserRequest req)
    CreateUserResponse CreateUser(1: CreateUserRequest req)
    QueryUserResponse QueryUser(1: QueryUserRequest req)
}