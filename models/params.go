package models

// 定义注册请求的参数的结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 投票
type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"`              //帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` //赞成票(1) 反对票(-1)
}

// ParamPostList  获取帖子列表query string 参数
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`
	Page        int64  `json:"page" form:"page"`
	Size        int64  `json:"size" form:"size"`
	Order       string `json:"order" form:"order"`
}

const (
	OrderTime  = "time"
	OrderScore = "score"
)
