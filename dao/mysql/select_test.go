package mysql

import (
	"bluebell/models"
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestSelect(T *testing.T) {
	dsn := "root:admin123@tcp(127.0.0.1:3306)/bluebell?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("connect err")
		return
	}
	var communityList []models.Community
	sqlStr := "select community_id,community_name from community where id>?"
	db.Select(&communityList, sqlStr, 0)
	fmt.Printf("%v\n", communityList)
}
