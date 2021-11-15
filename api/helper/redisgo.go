package helper

import (
	"app/api/config"
	"app/api/model"
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func GetPosts(list interface{}) bool {
	isset, err := redis.Int64(config.Client.Do("EXISTS", "posts"))
	Err(err)
	if isset == 0 {
		return false
	}
	return true
}

func SetPosts(posts *[]model.Post) bool {
	json, err := json.Marshal(posts)
	Err(err)
	fmt.Println("")
	fmt.Println("json : ", string(json))
	val, err := config.Client.Do("set", "posts", string(json))
	Err(err)
	if val == 0 {
		return false
	}
	return true
}

func Err(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
