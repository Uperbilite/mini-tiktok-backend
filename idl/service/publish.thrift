namespace go publish

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

struct Video {
    1: required i64 id
    2: required User author
    3: required string play_url
    4: required string cover_url
    5: required i64 favorite_count
    6: required i64 comment_count
    7: required bool is_favorite
    8: required string title
}

struct PublishVideoRequest {
    1: required i64 user_id (vt.gt = "0")
    2: required binary data
    3: required string title (vt.min_size = "1")
}

struct PublishVideoResponse {
    1: required BaseResp base_resp
}

service PublishService {
    PublishVideoResponse PublishVideo(1: PublishVideoRequest req)
}