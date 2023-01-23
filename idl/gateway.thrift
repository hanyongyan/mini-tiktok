namespace go api

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
}

struct Video {
    1: i64 id
    2: User author
    3: string play_url
    4: string cover_url
    5: i64 favorite_count
    6: i64 comment_count
    7: bool is_favorite
    8: string title
}

struct Comment {
    1: i64 id
    2: User user
    3: string content
    4: string create_date
}

struct FeedReq {
    1: optional string latest_time
    2: optional string token
}

struct FeedResp {
    1: i64 status_code
    2: string status_message
    3: i64 next_time
    4: list<Video> video_list
}

struct UserRegisterReq {
    1: string username (api.query="username", api.vd="len($) <= 32")
    2: string password (api.query="password", api.vd="len($) <= 32>")
}

struct UserRegisterResp {
    1: i64 status_code
    2: string status_message
    3: i64 user_id
    4: string token
}



struct UserLoginReq {
    1: string username (api.query="username", api.vd="len($) <= 32")
    2: string password (api.query="password", api.vd="len($) <= 32")
}

struct UserLoginResp {
    1: i64 status_code
    2: string status_message
    3: i64 user_id
    4: string token
}

struct UserReq {
    1: string user_id
    2: string token
}

struct UserResp {
    1: i64 status_code
    2: string status_message
    3: User user
}

struct PublishActionReq {
    1: list<byte> data
    2: string token
    3: string title
}

struct PublishActionResp {
    1: i64 status_code
    2: string status_message
}

struct PublishListReq {
    1: string token
    2: string user_id
}

struct PublishListResp {
    1: i64 status_code
    2: string status_message
    3: list<Video> video_list
}

struct FavoriteActionReq {
    1: string token
    2: string video_id
    3: string action_type
}

struct FavoriteActionResp {
    1: i64 status_code
    2: string status_message

}

struct FavoriteListReq {
    1: string user_id
    2: string token
}

struct FavoriteListResp {
    1: i64 status_code
    2: string status_message
    3: list<Video> video_list
}

struct CommentActionReq {
    1: string token
    2: string video_id
    3: string action_type
    4: optional string comment_text
    5: optional string comment_id
}

struct CommentActionResp {
    1: i64 status_code
    2: string status_message
    3: Comment comment
}

struct CommentListReq {
    1: string token
    2: string video_id
}

struct CommentListResp {
    1: i64 status_code
    2: string status_message
    3: list<Comment> comment_list
}

struct RelationActionReq {
    1: string token
    2: string to_user_id
    3: string action_type
}

struct RelationActionResp {
    1: i64 status_code
    2: string status_message

}

struct RelationFollowListReq {
    1: string user_id
    2: string token
}

struct RelationFollowListResp {
    1: i64 status_code
    2: string status_message
    3: list<User> user_list

}

struct RelationFollowerListReq {
    1: string user_id
    2: string token
}

struct RelationFollowerListResp {
    1: i64 status_code
    2: string status_message
    3: list<User> user_list
}

struct RelationFriendListReq {
    1: string user_id
    2: string token
}

struct RelationFriendListResp {
    1: i64 status_code
    2: string status_message
    3: list<User> user_list
}

service ApiService {
    FeedResp Feed(1: FeedReq req) (api.get="/douyin/feed")
    UserRegisterResp UserRegister(1: UserRegisterReq req) (api.post="/douyin/user/register")
    UserLoginResp UserLogin(1: UserLoginReq req) (api.post="/douyin/user/login")
    UserResp User(1: UserReq req) (api.get="/douyin/user")
    PublishActionResp PublishAction(1: PublishActionReq req) (api.post="/douyin/publish/action")
    PublishListResp PublishList(1: PublishActionReq req) (api.get="/douyin/publish/list")
    FavoriteActionResp FavoriteAction(1: FavoriteActionReq req) (api.post="/douyin/favorite/action")
    FavoriteListResp FavoriteList(1: FavoriteListReq req) (api.get="/douyin/favorite/list")
    CommentActionResp CommentAction(1: CommentActionReq req) (api.post="/douyin/comment/action")
    CommentListResp CommentList(1: CommentListReq req) (api.get="/douyin/comment/list")
    RelationActionResp RelationAction(1: RelationActionReq req) (api.post="/douyin/relation/action")
    RelationFollowListResp RelationFollowList(1: RelationFollowListReq req) (api.get="/douyin/relation/follow/list")
    RelationFollowerListResp RelationFollowerList(1: RelationFollowerListReq req) (api.get="/douyin/relation/follower/list")
    RelationFriendListResp RelationFriendList(1: RelationFriendListReq req) (api.get="/douyin/relation/friend/list")
}