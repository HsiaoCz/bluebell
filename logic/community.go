package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"

	"go.uber.org/zap"
)

// GetCommunityList 查询数据
func GetCommunityList() (data []models.Community, err error) {
	//查找到所有的community并返回
	data, err = mysql.GetCommunityList()
	zap.L().Error("mysql.GetCommunityList() failed", zap.Error(err))
	return data, err
}

// GetCommunityDetail 根据ID查询社区详情
func GetCommunityDetail(id int64) (data []models.CommunityDetail, err error) {
	//查询所有分类的详情
	//要查询东西 我们先要有一个结构体进行映射
	data, err = mysql.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("获取数据失败", zap.Error(err))
	}
	return data, err
}
