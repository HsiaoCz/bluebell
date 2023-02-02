package redis

import (
	"bluebell/models"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func getIDsFormKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	return rdb.ZRevRange(key, start, end).Result()
}

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	//从redis 获取ID
	//1.根据用户请求中携带的order 参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZset)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZset)
	}
	//2.确定查询的索引起始点
	//3. ZREVRANGE 查询  按照分数从大到小的顺序查询 查询指定数量的元素
	return getIDsFormKey(key, p.Page, p.Size)
}

// 根据ids查询每篇帖子的赞成票数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	// data = make([]int64, 0, len(ids))
	// for _, id := range ids {
	// 	key := getRedisKey(KeyPostVotedZsetPF + id)
	// 	//查找key中分数是1的数量
	// 	v := rdb.ZCount(key, "1", "1").Val()
	// 	data = append(data, v)
	// }
	//使用pipeline 一次发送多条命令 减少rtt
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZsetPF + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

// GetCommunityPostList 按照社区查询ids
func GetCommunityPostList(p *models.ParamPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZset)
	if p.Order == models.OrderScore {
		orderKey = getRedisKey(KeyPostScoreZset)
	}
	// 使用zinterstore 把分区的帖子set与帖子分数的zset 生成一个新的zset
	// 针对新的zset 按之前的逻辑取数据
	ckey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))
	//利用缓存key 减少zinterstore 执行的次数
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if rdb.Exists(orderKey).Val() < 1 {
		//不存在，需要计算
		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, ckey, orderKey)
		pipeline.Expire(key, 60*time.Second) //设置超时时间
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	return getIDsFormKey(key, p.Page, p.Size)
}
