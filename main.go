package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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
	//
	//wg.Add(1)
	//go func() {
	//	fmt.Println("done2")
	//	wg.Done()
	//} ()

	//var wg sync.WaitGroup
	//for i:=0; i<10; i++ {
	//	wg.Add(1)
	//	go func(i int) { //引数を追加
	//		defer wg.Done()
	//		fmt.Println(i)
	//	}(i) //関数実行時に現在の値を渡している
	//
	//	//go func() {
	//	//	defer wg.Done()
	//	//	fmt.Println(i) //これだとループの最後の値になってしまうことが多い
	//	//} ()
	//}
	//wg.Wait()
	//fmt.Println("unlock1")
	//
	//for i:=0; i<20; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		fmt.Printf("%v回目\n", i)
	//	}(i)
	//}
	//wg.Wait()
	//fmt.Println("unlock2")



	//⑨TODO　golang.org/x/sync/errgroup

	//⑩sync.Once
	//var once sync.Once
	//once.Do(f)
	//once.Do(f) //2回目以降は実行されない
	//once.Do(f) //2回目以降は実行されない
	//fmt.Println("done")


	//11 ゴルーチンを跨いだ処理のキャンセル
	//cf. 関数の書き方
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
	//
	////チャネルもfor rangeで回すことができる。チャネルを受け取るたびnに値(int)が入る。
	//for n := range gen(ctx) {
	//	//fmt.Println(n)
	//	fmt.Printf("%T: %v\n", n, n)
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


	//13 sync.errgroup(errgroup.Group)
	test1 := func(ctx context.Context) error {
		fmt.Printf("test1にて%T: %v\n", ctx, ctx)

		errCh := make(chan error)
		//fmt.Printf("%T: %v\n", errCh, errCh)

		//defer close(errCh)

		go func() {
			//チャネルクローズするのを忘れないように
			defer close(errCh)
			errCh <-errors.New("test1のエラー")
		}()


		select {
		case <-ctx.Done():
			fmt.Println("test1にてctx.Doneチャネル受信")
			return ctx.Err()
		case err:= <-errCh:
			fmt.Printf("test1にてerrCh受信:%T:%v\n", err, err)
			return err
		}
	}
	test2 := func(ctx context.Context) error {
		//time.Sleep(3*time.Second)
		//return errors.New("test2")

		errCh := make(chan error)
		go func(){
			errCh <- errors.New("test2のエラー")
		}()

		select {
		case <-ctx.Done():
			fmt.Println("test2にてctx.Doneチャネル受信")
			return ctx.Err()
		case err := <-errCh:
			fmt.Printf("test2にて%T:%v\n", err, err)
			return err
		}
	}

	ctx := context.Background()

	var eg *errgroup.Group
	eg, ctx = errgroup.WithContext(ctx)

	eg.Go(func() error {
		return test1(ctx)
	})
	eg.Go(func() error {
		return test2(ctx)
	})
	eg.Go(func() error {
		//fmt.Printf("%T: %v\n", ctx.Done(), ctx.Done())
		// この書き方覚える。受信専用チャネル。
		<-ctx.Done()
		//fmt.Printf("%T: %v\n", ctx.Err(), ctx.Err())
		return ctx.Err()
	})

	if err := eg.Wait(); err != nil {
		log.Println(err)
	}

	//13の練習
	//defer fmt.Println("test")
	//os.Exit(run(context.Background()))

	//練習用
	exercise()
}

//13 WithValueでコンテキストに値を持たせる
//type withoutCacheKey struct{}
//func WithoutCache(c context.Context) context.Context {
//	if IsIgnoredCache(c) {
//		return c
//	}
//	return context.WithValue(c, withoutCacheKey{}, struct{}{})
//}
//func IsIgnoredCache(c context.Context) bool {
//	return c.Value(withoutCacheKey{}) != nil
//}


func exercise() {
	//6の練習
	//var m sync.Mutex
	//m.Lock()
	//fmt.Println("locked")
	//go func() {
	//	time.Sleep(3*time.Second)
	//	fmt.Println("lock解除")
	//	m.Unlock()
	//}()
	//m.Lock()
	//m.Unlock()
	//fmt.Println("unlocked")

	//7の練習
	//var m sync.RWMutex
	//m.RLock()
	//fmt.Println("locked")
	//go func() {
	//	time.Sleep(3*time.Second)
	//	fmt.Println("lock解除する")
	//	m.RUnlock()
	//}()
	//m.RLock()
	//fmt.Println("ここは出力されない")
	//m.RUnlock()
	//fmt.Println("unlocked")

	//10の練習
	//var once sync.Once
	////once.Do(onceTest)
	//
	//for i:=1; i<10; i++ {
	//	go func(i int) {
	//		once.Do(onceTest)
	//		fmt.Println(i)
	//	}(i )
	//}
	//
	//go func() {
	//	fmt.Println("ループ外")
	//	once.Do(onceTest)
	//}()
	//
	//time.Sleep(10*time.Second) //スリープしないとgo分より先にmainが終了されてしまう


	//11の練習
	//bc := context.Background()
	//ctx, cancel := context.WithCancel(bc)
	//defer cancel()
	//
	//gen := func(ctx context.Context) <-chan int {
	//	dst := make(chan int)
	//	n := 1
	//	go func() {
	//		for {
	//			select {
	//			case <-ctx.Done(): return
	//			//送信の書き方
	//			case dst <- n: n++
	//			}
	//		}
	//	}()
	//	return dst
	//}
	//
	////for rangeの書き方 https://qiita.com/najeira/items/71a0bcd079c9066347b4
	//for n:= range gen(ctx) {
	//	fmt.Println(n)
	//	if n == 20 { break }
	//}

	//12の練習
	//bc := context.Background()
	//t := 50*time.Millisecond
	//ctx, cancel := context.WithTimeout(bc, t)
	//defer cancel()
	//
	//select {
	//case <- time.After(1 * time.Second):
	//	fmt.Println("overslept")
	//case <- ctx.Done():
	//	fmt.Println(ctx.Err())
	//}

}

//10の練習用
func onceTest() {
	fmt.Println("一回だけ実行される")
}

//13
func run(ctx context.Context) int {
	eg, ctx := errgroup.WithContext(ctx)
	//fmt.Printf("%T: %v\n", eg, eg)
	//fmt.Printf("%T: %v\n", ctx, ctx)

	eg.Go(func() error {
		//fmt.Println("error1")
		//return errors.New("test1")
		return runServer(ctx)
	})
	eg.Go(func() error {
		//fmt.Println("error2")
		//return errors.New("test2")
		return acceptSignal(ctx)
	})
	eg.Go(func() error{
		<-ctx.Done()
		fmt.Println("ctx.Done()")
		return ctx.Err()
	})

	if err := eg.Wait(); err != nil {
		log.Println(err)
		return 1
	}
	log.Println("0!")
	return 0
}

func runServer(ctx context.Context) error {
	s := &http.Server{
		Addr: ":8888",
		Handler: nil,
	}

	errCh := make(chan error)
	go func() {
		defer close(errCh)
		if err := s.ListenAndServe(); err != nil {
			//チャネルへの値の入れ方
			errCh <- err
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			fmt.Println("runServer処理中...")
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("runSeverのctx.Doneチャネル")
		return s.Shutdown(ctx)
	case err := <-errCh:
		fmt.Println("runSeverのerrCh!!")
		return err
	}


	//return errors.New("runServerError")
}

func acceptSignal(ctx context.Context) error {
	sigCh := make(chan os.Signal, 1)

	//受け取るシグナルを設定
	//"os.Interrupt"は、Ctrl+C で中断されるときに送られるシグナル(SIGINT)のこと
	signal.Notify(sigCh, os.Interrupt)

	// シグナル待機中にやりたい処理 ※goroutine(並行処理)で書く
	go func() {
		for {
			time.Sleep(2 * time.Second)
			fmt.Println("acceptSignal処理中...")
		}
	}()


	select {
	case <-ctx.Done():
		fmt.Println("acceptSignalのctx.Doneチャネル")
		signal.Reset()
		return ctx.Err()
	case sig := <-sigCh:
		fmt.Println("acceptSignalのerrCh!!")
		return fmt.Errorf("signal received: %v", sig.String())
	}

	//return errors.New("signalError")
}
