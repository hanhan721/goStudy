package main

import (
	"encoding/json"
	"fmt"
)

// 使用tag标签配置json库,处理json数据
type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 1998, []string{"周星驰", "张柏芝"}}
	// 结构体转为json
	jsonString, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("parse json error", err)
	} else {
		fmt.Printf("json->%s", jsonString)
		fmt.Println()
	}
	// json转为结构体
	movie2 := Movie{}
	err2 := json.Unmarshal(jsonString, &movie2)
	if err2 != nil {
		fmt.Println("parse struct error", err2)
	} else {
		fmt.Println("movie2->", movie2)
	}
}
