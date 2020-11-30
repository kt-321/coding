package main

import (
	"context"
	"fmt"
)

//⑩
func f(){ fmt.Println(("Do")) }

func main() {
	//①
	//go func() {
	//	fmt.Println("別のゴルーチン")
	//}()
	//fmt.Println("mainゴルーチン")
	//time.Sleep(50*time.Millisecond) //Sleepしないとmainゴルーチンがすぐに終了してしまう。同期を取るため記述。

	//②
	//ch := make(chan int) // 容量0
	//go func() {
	//	ch <-100 //送信時にチャネルのバッファがいっぱいだとブロックとなり処理が終わらない
	//}()
	//go func() {
	//	v := <-ch //受信時にチャネル内が空だとブロックとなり処理が終わらない
	//	fmt.Println(v)
	//}()
	//time.Sleep(2 * time.Second) //Sleepしないとmainゴルーチンがすぐに終了してしまう


	//③
	//ch1 := make(chan int)
	//ch2 := make(chan string)
	//
	//go func() { ch1 <- 100 }()
	//go func() { ch2 <- "hi"}()
	//
	//select {
	//case v1 := <-ch1:
	//	fmt.Println(v1)
	//case v2 := <-ch2:
	//	fmt.Println(v2)
	//}

	//④
	//ch1 := make(chan int)
	//var ch2 chan string //ゼロ値はnil
	//
	//fmt.Printf("%T: %v\n", ch1, ch1)
	//fmt.Printf("%T: %v\n", ch2, ch2)
	//
	//
	//go func() { ch1 <- 100 }()
	//go func() { ch2 <- "hi" }()
	//
	//select {
	//case v1 := <-ch1:
	//	fmt.Println(v1)
	//case v2 := <-ch2: //nilの場合は無視。caseのことは無かったことに
	//	fmt.Println(v2)
	//}

	//⑤チャネルのチャネル
	//ch := make(chan int)
	//go func(){
	//	//時間のかかる処理
	//	time.Sleep(6 * time.Second)
	//	ch <- 1
	//}()
	//select {
	//case <- time.After(5 * time.Second): //チャネルのチャネル time.Afterはチャネルを返す 計測のスタートはtime.Afterが呼ばれた段階
	//	fmt.Println("タイムアウト")
	//case v1 := <-ch:
	//	fmt.Println(v1)
	//}

	//⑥ロック sync.Mutex
	//var m sync.Mutex //ゼロ値で使える
	//fmt.Printf("%T: %v\n", m, m)
	//m.Lock()
	//go func() {
	//	time.Sleep(3 * time.Second)
	//	m.Unlock()
	//	fmt.Println("unlock1")
	//}()
	//m.Lock()//ここでブロック
	//m.Unlock()
	//fmt.Println("unlock2")

	//⑦sync.RWMutex
	//var m sync.RWMutex
	//m.RLock()
	//go func() {
	//	time.Sleep(3 * time.Second)
	//	m.RUnlock()
	//	fmt.Println("unlock1")
	//}()
	//m.RLock() //読み込みロックだけはブロックしない。もし上の記述でm.Lock()/m.Unlock()が使われていると、ブロックされる
	//m.RUnlock()
	//fmt.Println("unlock2")

	//⑧sync.WaitGroup 複数のゴルーチンの処理が完了するのを待つ
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	fmt.Println("done1")
	//	wg.Done()
	//}()
	//wg.Add(1)
	//go func() {
	//	fmt.Println("done2")
	//	wg.Done()
	//} ()
	//wg.Wait()
	//fmt.Println("unlock")

	//⑨TODO　golang.org/x/sync/errgroup

	//⑩sync.Once
	//var once sync.Once
	//once.Do(f)
	//once.Do(f) //2回目以降は実行されない
	//once.Do(f) //2回目以降は実行されない
	//fmt.Println("done")


	//11 ゴルーチンを跨いだ処理のキャンセル
	//gen := func(ctx context.Context) <-chan int {
	//	dst := make(chan int)
	//	n := 1
	//	go func() {
	//		//無限ループ
	//		for {
	//			select {
	//			case <-ctx.Done(): return //Doneチャネルが閉じられると反応する
	//			case dst <- n: n++
	//			}
	//		}
	//	}()
	//	return dst
	//}
	//
	//bc := context.Background() //ルートになるコンテキスト
	//ctx, cancel := context.WithCancel(bc) //ルートになるコンテキストをラップしてキャンセル機能のついたコンテキストを生成
	//defer cancel() //contextのDoneメソッドで返ってくるチャネル(Doneチャネル)がクローズされる。Doneメソッドで待っていたチャネルに終了が告げられる。

	////チャネルもfor rangeで回すことができる。
	//for n := range gen(ctx) {
	//	fmt.Println(n)
	//	if n == 20 { break }
	//}

	//参考
	//type Context interface {
	//	Deadline() (deadline time.Time, ok bool)
	//	Done() <-chan struct{}
	//	Err() error
	//	Value(key interface{}) interface{}
	//}


	//12 タイムアウト context.WithTimeout
	//bc := context.Background()
	//t := 50*time.Millisecond
	//ctx, cancel := context.WithTimeout(bc, t)
	//defer cancel()
	//select {
	//case <-time.After(1 * time.Second):
	//	fmt.Println("overslept")
	//case <-ctx.Done():
	//	fmt.Println(ctx.Err()) //「deadline exceeded」というキャンセルされたというエラーが返ってくる
	//}

}

//13 WithValueでコンテキストに値を持たせる
type withoutCacheKey struct{}
func WithoutCache(c context.Context) context.Context {
	if IsIgnoredCache(c) {
		return c
	}
	return context.WithValue(c, withoutCacheKey{}, struct{}{})
}
func IsIgnoredCache(c context.Context) bool {
	return c.Value(withoutCacheKey{}) != nil
}
