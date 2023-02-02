package snowflake

import (
	"bluebell/setting"
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
)

var node *snowflake.Node

func Init() (err error) {
	app := setting.Conf.AppConfig
	var st time.Time
	fmt.Println(app.StartTime)
	st, err = time.Parse("2006-01-02", viper.GetString("app.start_time"))
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(app.MachineID)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}

// func main() {
// 	if err := Init("2020-07-01", 1); err != nil {
// 		fmt.Printf("init failed,err:%v\n", err)
// 		return
// 	}
// 	id := GenID()
// 	fmt.Println(id)
// }
