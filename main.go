package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page int
	Fruits []string
}

type response2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	//fmt.Print("test")

	//test := "hoge"
	//test := 1
	//test := 2.34
	//test := []string {"hoge", "fuge", "aaa"}
	//test := map[string]int {"hoge": 1, "fuge": 2, "aaa": 3}
	//test := &response1{
	//	Page: 1,
	//	Fruits: []string{"hoge", "fuga", "aaa"},
	//}
	test := &response2{
		Page: 1,
		Fruits: []string{"hoge", "fuga", "aaa"},
	}
	testMarshaled, err := json.Marshal(test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(testMarshaled))
}