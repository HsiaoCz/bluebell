package redis

import (
	"errors"
	"math"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

// PostVote 投票功能
// 1、用户投票的数据
// 使用简化版的投票分数算法
// 投一票就加432分  86400/200 可以给你的帖子续一天

// 投票的几种情况:
// direction=1 时 有两种情况:
// 1.之前没有投过票 现在投赞成票  --> 更新分数和投票记录
// 2.之前投反对票，现在改投赞成票
// direction=0 时 有两种情况
// 1.之前投过赞成票 现在要取消投票
// 2.之前投过反对票，现在要取消投票
//direction=-1 时 有两种情况
// 1.之前没有投过票 现在投反对票
// 2.之前投赞成票 现在改投反对票

// 投票的规则
// 每个帖子自发表之日起一个星期之内允许用户投票 超过一个星期就不允许再投票了
// 1、到期之后讲Redis中保存的赞成票及反对票存储到MySQL表中
// 2、到期之后删除那个KeyPostVotedZsetPF
const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePrevote     = 432 //每一票值多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

func CreatePost(postID, communityID int64) (err error) {
	//这两个要添加到一个事务里面
	// 要么同时成功 要么同时失败
	pipeline := rdb.TxPipeline()
	//讲帖子创建时间添加到redis
	_, err = pipeline.ZAdd(getRedisKey(KeyPostTimeZset), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Result()
	if err != nil {
		zap.L().Error("redis.CreatePost failed", zap.Error(err))
		return
	}
	//将帖子分数添加到redis
	_, err = pipeline.ZAdd(getRedisKey(KeyPostScoreZset), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Result()
	if err != nil {
		zap.L().Error("pipeline.ZAdd err:", zap.Error(err))
		return
	}
	// 把帖子id加到社区的set
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
	pipeline.SAdd(cKey, postID)
	_, err = pipeline.Exec()
	return
}

func VoteForPost(userID, postID string, value float64) (err error) {
	//1.判断投票的限制
	// 去redis 拿到帖子发布时间
	//先判断要投票的帖子是否还在有效时间
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZset), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	//2和3需要放到一个pipeline里
	//2.更新帖子的分数
	//先查之前的投票记录
	ov := rdb.ZScore(getRedisKey(KeyPostVotedZsetPF+postID), userID).Val()
	//如果这一次投票的值和之前保存的值一致，就提示不允许重复投票
	if value == ov {
		return ErrVoteRepeated
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostTimeZset), op*diff*scorePrevote, postID)
	if ErrVoteTimeExpire != nil {
		return err
	}
	//3.记录用户为帖子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZsetPF+postID), postID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZsetPF+postID), redis.Z{
			Score:  value, //赞成票还是反对票
			Member: userID,
		})
	}
	_, err = pipeline.Exec()
	return err
}
