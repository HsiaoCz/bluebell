package controllors

import (
	"bluebell/logic"
	"bluebell/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	//1.获取参数及参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) failed", zap.Any("err:", err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//从c 取到当前发请求的用户的ID值
	userID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	//2、创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3、返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子详情的处理函数
func GetPostDetailHandler(c *gin.Context) {
	//1、获取URL中帖子的id
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//2、根据id 取出帖子数据
	data, err := logic.GetPostDetailByID(pid)
	if err != nil {
		zap.L().Error("logic.GetPostDetailByID failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler 获取帖子列表处理函数
// 查询结果分页
func GetPostListHandler(c *gin.Context) {
	//获取分页参数
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 0
	}
	size, err := strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	//获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	//返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler1 升级帖子列表接口
// 根据前端传来的参数动态的获取帖子列表
// 按创建时间排序 或者 按照分数排序
// 1、获取参数
// 2、去redis查询id 列表
// 3、根据id数据库查询帖子详细信息
func GetPostListHandler1(c *gin.Context) {
	//初始化结构体时指定一些初始参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime, //magic string
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler1 failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// c.ShouldBindJson()  获取json格式的数据
	//获取数据
	data, err := logic.GetPostListNew(p) //更新合二为一
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	//返回响应
	ResponseSuccess(c, data)

}

// 根据社区去查询帖子列表
// func GetCommunityPostListHandler(c *gin.Context) {
// 	//初始化结构体时指定一些初始参数
// 	p := &models.ParamPostList{
// 		Page:  1,
// 		Size:  10,
// 		Order: models.OrderTime,
// 	}
// 	if err := c.ShouldBindQuery(p); err != nil {
// 		zap.L().Error("GetCommunityPostListHandler failed", zap.Error(err))
// 		ResponseError(c, CodeInvalidParam)
// 		return
// 	}
// 	// c.ShouldBindJson()  获取json格式的数据
// 	//获取数据
// 	data, err := logic.GetCommunityPostList(p)
// 	if err != nil {
// 		zap.L().Error("logic.GetPostList failed", zap.Error(err))
// 		ResponseError(c, CodeServerBusy)
// 	}
// 	//返回响应
// 	ResponseSuccess(c, data)
// }
