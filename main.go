package main

import (
	"encoding/json"
	"fmt"
	"os"
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

	byt := []byte(`{"num":6.13, "strs":["a", "b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T: %v: %T: %v: %T: %v\n", dat, dat, dat["num"], dat["num"], dat["strs"], dat["strs"])

	num := dat["num"].(float64)
	fmt.Printf("%T: %v:\n", num, num)

	strs := dat["strs"].([]interface{})
	fmt.Printf("%T: %v:\n", strs, strs)
	str1 := strs[0].(string)
	fmt.Printf("%T: %v:\n", str1, str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	fmt.Printf("%T: %v:\n", str, str)
	res := &response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Printf("%T: %v: %T: %v\n", res, res, res.Fruits[0], res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce" :7}
	enc.Encode(d)
}