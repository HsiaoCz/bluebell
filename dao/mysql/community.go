package mysql

import (
	"bluebell/models"
	"database/sql"
	"fmt"

	"go.uber.org/zap"
)

// GetCommunityList 从数据库查询出community 并返回
func GetCommunityList() (communityList []models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	if err = db.Select(&communityList, sqlStr); err != nil {
		fmt.Printf("select err:%v\n", err)
		if err == sql.ErrNoRows {
			zap.L().Error("there is no community in db")
			err = nil
		}
	}
	return
}

// GetCommunityDetail 获取社区分类详情
func GetCommunityDetail(id int64) (communityDetail []models.CommunityDetail, err error) {
	sqlStr := "select community_id,community_name,introduct_name,create_time from community where community_id=?"
	//这里虽然只查一行 但是依然要使用select
	err = db.Select(&communityDetail, sqlStr, id)
	if err == sql.ErrNoRows {
		err = ErrorInvalidID
	}
	return
}

func GetCommunityDetailById(id int64) (community models.CommunityDetail, err error) {
	sqlStr := "select community_id,community_name,introduct_name,create_time from community where community_id=?"
	err = db.Get(&community, sqlStr, id)
	if err != nil {
		zap.L().Error("get communityDetail failed", zap.Error(err))
		return
	}
	return
}
