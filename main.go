package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

//既存のstructや型に対して、ServeHTTPメソッドを用意することでhttp.Handleに登録出来るようにする
type StringTest string

func (s StringTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//はじめに「F」が付いているものは、書き込み先を明示的に指定できる。
	fmt.Fprint(w, s)
}

type IntTest int

func (i IntTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%T: %v\n", i, i)
}

type Cat struct {
	ID        uint
	Name      string
	Nickname  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (c Cat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, c)
	//fmt.Fprintf(w, "%T: %v\n", c, c)

	//byteTest := []byte{'a', 'b'}

	fmt.Printf("%T: %v\n", c, c)

	byteTest, err := json.Marshal(c)
	//byteTest, err := json.Marshal("fugafuga")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T: %v", byteTest, byteTest)


	if _, err := w.Write(byteTest); err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	fmt.Println("起動")
	r := mux.NewRouter()

	//ハンドラを登録
	http.HandleFunc("/test", httpTest)
	http.HandleFunc("/hoge", hoge)

	r.HandleFunc("/formTest", formTest).Methods("POST")

	http.Handle("/string", StringTest("stringtest"))
	http.Handle("/int", IntTest(3))

	cat := Cat{
		ID:        0,
		Name:      "neko",
		Nickname:  "oneko",
	}
	http.Handle("/cat", Cat(cat))

	go postFormTest("test")

	//httpサーバー立ち上げ
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println(err)
	}

}

func httpTest(w http.ResponseWriter, r *http.Request) {
	log.Printf("%T: %v\n", r, r)
	fmt.Println("testtest")
}

func hoge(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hogehoge")
}

type Name struct {
	ID uint
	Name string
}

type Jsontest struct {
	ID uint
	Name string
	Token string
}

//POSTリクエストを受けとってjson形式でレスポンスを返す
func formTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("formTest開始")

	//curl -X POST -d 'Name=name' http://localhost:8080/formTest でリクエスト投げた
	value := r.FormValue("Name")
	log.Printf("%T: %v\n", value, value)

	d := Jsontest{
		ID:    1,
		Name:  value,
		Token: "token",
	}

	v, err := json.Marshal(d)
	if err != nil {
		fmt.Println(err)
	}

	if _, err := w.Write(v); err != nil {
		fmt.Println(err)
		return
	}

	//curl -X POST -H "Content-Type: application/json" -d '{"ID": 1, "Name":"sensuikan1973"}' http://localhost:8080/formTest
	//log.Printf("r.Body...%T: %v", r.Body, r.Body)
	//log.Printf("r.Header...%T: %v", r.Header, r.Header)
	//
	//var d Jsontest
	//
	//if err := json.NewDecoder(r.Body).Decode(&d); err != nil{
	//	fmt.Println(err)
	//}
	//
	//fmt.Printf("d...%T: %v\n", d, d)
	//fmt.Printf("d.ID...%T: %v\n", d.ID, d.ID)
	//fmt.Printf("d.Name...%T: %v\n", d.Name, d.Name)
	//fmt.Printf("d.Token...%T: %v", d.Token, d.Token)
}


//PostForm(url, data)関数の演習
func postFormTest(value string) {
	fmt.Println("postFormTest開始")

	values := url.Values{}
	values.Add("Name", value)

	url := "http://localhost:8080/formTest"

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			fmt.Println("ゴルーチン実行中")

			resp, err := http.PostForm(url, values)
			if err != nil {
				fmt.Println(err)
				return
			}
			//fmt.Printf("postFormTest...%T: %v\n", resp.Body, resp.Body)

			var j Jsontest

			if err := json.NewDecoder(resp.Body).Decode(&j); err !=nil {
				fmt.Println(err)
			}

			fmt.Printf("%T: %v\n", j, j)
		}()
	}
	wg.Wait()

	fmt.Println("10回実行ずみ")

	var strings =  []string{
		"hoge",
		"dddd",
		"rrrr",
		"fff",
		"ddgggdd",
		"dd12dd",
		"uuddddgdgd",
		"dddfagdd",
		"agddddd",
		"hdddd",
	}


	for _, str := range strings {
		wg.Add(1)
		go clientDoTest(str, wg)
	}
	wg.Wait()
	fmt.Println("clientDoTest全て実行完了")
}

//client.Do(Request)メソッドでPOST実装を示す
func clientDoTest(str string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("clientDoTest開始")

	fmt.Printf("clientDoTestにて%v\n", str)

	values := url.Values{}
	values.Add("Name", str)

	url := "http://localhost:8080/formTest"

	req, err := http.NewRequest("POST", url, strings.NewReader(values.Encode()))
	if err != nil {
		fmt.Println(err)
		return
	}

	//Content-Typeリクエストヘッダをapplication/x-www-form-urlencodedに設定する
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	//
	defer resp.Body.Close()

	//fmt.Printf("%T: %v", resp.Body, resp.Body)

	var j Jsontest

	if err := json.NewDecoder(resp.Body).Decode(&j); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%T: %v\n", j, j)

	fmt.Println("resp取得完了")
}
