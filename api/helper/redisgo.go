package helper

import (
	"app/api/config"
	"fmt"
)

func SetOrGet(list interface{}) []interface{} {
	isset, err := config.Client.Do("exists", "posts")
	Err(err)
	if isset == 0 {
		val, err := config.Client.Do("hset", "posts", "name", "deneme")
		Err(err)

		fmt.Println(val)
		fmt.Println("a")
	}

	return nil
}

func Err(err error) {
	if err != nil {
		panic(err)
	}
}
