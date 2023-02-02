package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	//1、生成post id
	p.ID = snowflake.GenID()
	//2、保存到数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID, p.CommunityID)
	if err != nil {
		zap.L().Error("redis.CreatePost failed", zap.Error(err))
		return
	}
	//3、返回
	return
}

func GetPostDetailByID(pid int64) (data models.ApiPostDetail, err error) {
	//查询数据 并凭借组合我们想用的数据
	post, err := mysql.GetPostDetailByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailByID failed", zap.Error(err))
		return
	}
	//根据作者的id 查询作者的信息
	user, err := mysql.GetPostUser(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetPostUser failed", zap.Error(err))
		return
	}
	//根据社区ID 查询社区详情
	community, err := mysql.GetCommunityDetailById(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID failed", zap.Error(err))
		return
	}
	data.Post = post
	data.AuthorName = user.Username
	data.CommunityDetail = community
	return
}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList failed:", zap.Error(err))
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		//根据作者的id 查询作者的信息
		user, err := mysql.GetPostUser(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetPostUser failed", zap.Error(err))
			continue
		}
		//根据社区ID 查询社区详情
		community, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID failed", zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

// GetPostList1 获取帖子列表
func GetPostList1(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	//去redis 查询id 列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		zap.L().Error("redis.GetPostIDsInOrder failed", zap.Error(err))
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return success but return empty data")
		return
	}
	// 根据id去数据库查询帖子详细信息
	// 返回的数据还要按照给定的id 顺序返回
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		zap.L().Debug("mysql.GetPostListByIDs(ids)", zap.Any("posts", posts))
	}
	//提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		zap.L().Error("redis.GetPostVoteData(ids) failed", zap.Error(err))
		return
	}
	//将帖子的作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		//根据作者的id 查询作者的信息
		user, err := mysql.GetPostUser(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetPostUser failed", zap.Error(err))
			continue
		}
		//根据社区ID 查询社区详情
		community, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID failed", zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNumber:      voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

// 根据community 获取帖子详情
func GetCommunityPostList(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	//去redis 查询id 列表
	ids, err := redis.GetCommunityPostList(p)
	if err != nil {
		zap.L().Error("redis.GetPostIDsInOrder failed", zap.Error(err))
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return success but return empty data")
		return
	}
	// 根据id去数据库查询帖子详细信息
	// 返回的数据还要按照给定的id 顺序返回
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		zap.L().Debug("mysql.GetPostListByIDs(ids)", zap.Any("posts", posts))
	}
	//提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		zap.L().Error("redis.GetPostVoteData(ids) failed", zap.Error(err))
		return
	}
	//将帖子的作者及分区信息查询出来填充到帖子中
	for idx, post := range posts {
		//根据作者的id 查询作者的信息
		user, err := mysql.GetPostUser(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetPostUser failed", zap.Error(err))
			continue
		}
		//根据社区ID 查询社区详情
		community, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID failed", zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNumber:      voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

// GetPostListNew 两个查询逻辑合二为一的函数
func GetPostListNew(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	//根据请求参数的不同 执行不同的逻辑
	if p.CommunityID == 0 {
		//查所有
		data, err = GetPostList1(p)
	} else {
		//根据社区id查询
		data, err = GetCommunityPostList(p)
	}
	if err != nil {
		zap.L().Error("GetPostListNew failed:", zap.Error(err))
		return
	}
	return
}
