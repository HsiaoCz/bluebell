package controllors

import (
	"bluebell/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 社区相关的
//controllors 做参数处理

func CommunityHandler(c *gin.Context) {
	//查询到所有的社区(community_id,community_name) 以列表的形式返回
	communityList, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		//不轻易把服务端报错暴露给外部
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, communityList)
}

// CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	//1、获取URL上的id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//查询到所有的社区(community_id, community_name)
	//查询到所有的社区(community_id,community_name) 以列表的形式返回
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed", zap.Error(err))
		//不轻易把服务端报错暴露给外部
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
