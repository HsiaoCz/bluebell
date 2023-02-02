package redis

// redis
// redis key 注意使用命名空间的方式
const (
	KeyPrefix          = "bluebell:"
	KeyPostTimeZset    = "post:time"  //zset ;帖子及发帖时间
	KeyPostScoreZset   = "post:score" //zset 帖子及分数
	KeyPostVotedZsetPF = "post:voted" //zset 记录用户及投票的类型  参数是Post_id
	KeyCommunitySetPF  = "community:" //set 保存每个分区下帖子的id
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
