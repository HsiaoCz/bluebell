package mysql

import (
	"bluebell/models"
	"strings"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(post_id,title,content,author_id,community_id) values (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostDetailByID 根据ID获取帖子详情
func GetPostDetailByID(pid int64) (data models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time from post where post_id=?`
	err = db.Get(&data, sqlStr, pid)
	return
}

// GetPostList 获取帖子列表函数
func GetPostList(page, size int64) (data []models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time from post order by create_time limit ?,?`
	err = db.Select(&data, sqlStr, page, size)
	if err != nil {
		zap.L().Error("GetPostList failed:", zap.Error(err))
		return nil, err
	}
	return
}

// 根据给定的id 列表查询帖子数据
func GetPostListByIDs(ids []string) (postList []models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time from post where post_id in (?) order by FIND_IN_SET(post_id,?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
