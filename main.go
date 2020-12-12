package main

import (
	"context"
	"fmt"
	//"github.com/devlights/try-golang/lib/output"
	"runtime"
	"sync"
	"time"
)

func main() {
	if err := RuntimeMemoryStats(); err != nil {
		fmt.Println(err)
	}
}

// RuntimeMemoryStats は、runtime.MemoryStats() のサンプルです.
//
// REFERENCES::
//   - https://golangcode.com/print-the-current-memory-usage/
func RuntimeMemoryStats() error {
	var (
		rootCtx         = context.Background()
		mainCtx, cancel = context.WithCancel(rootCtx)
		wg              = &sync.WaitGroup{}
	)

	// 初期の状態を表示
	runtime.GC()
	printMemoryStats("init")

	// ---------------------------------------
	// データを500ミリ秒毎に増やしていく処理
	// ---------------------------------------
	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		var (
			tick  = time.Tick(500 * time.Millisecond)
			count = 0
			items = make([][]byte, 0, 5)
		)

		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case _ = <-tick:
				count++
				data := make([]byte, 0, 1024*1024)
				items = append(items, data)

				// output.Stderrf("[append]", "count=%d\ttick=%v\n", count, t)
			}
		}
	}(mainCtx, wg)

	// ---------------------------------------
	// 現在のメモリ量を2000ミリ秒毎に出力する処理
	// ---------------------------------------
	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		var (
			tick  = time.Tick(2000 * time.Millisecond)
			count = 0
		)

		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case <-tick:
				count++
				printMemoryStats(fmt.Sprintf("(%d)", count))
			}
		}
	}(mainCtx, wg)

	// 10秒したら終わり
	select {
	case <-time.After(10 * time.Second):
		cancel()
		wg.Wait()
	}

	// 最後の状態を表示
	printMemoryStats("latest")

	// GC後の状態を表示
	runtime.GC()
	printMemoryStats("after runtime.GC()")

	return nil
}

func printMemoryStats(prefix string) {
	// --------------------------------------------------------
	// runtime.MemoryStats() から、現在の割当メモリ量などが取得できる.
	//
	// まず、データの受け皿となる runtime.MemStats を初期化し
	// runtime.ReadMemStats(*runtime.MemStats) を呼び出して
	// 取得する.
	// --------------------------------------------------------
	var (
		ms runtime.MemStats
	)

	fmt.Println(prefix, "--------------------")

	//output.Stdoutl(prefix, "----------------------------")
	runtime.ReadMemStats(&ms)
	//
	//// Alloc は、現在ヒープに割り当てられているメモリ
	//// HeapAlloc と同じ.
	//output.Stdoutl("Alloc", toKb(ms.Alloc))
	//output.Stdoutl("HeapAlloc", toKb(ms.HeapAlloc))
	//
	//// TotalAlloc は、ヒープに割り当てられたメモリ量の累積
	//// Allocと違い、こちらは増えていくが減ることはない
	//output.Stdoutl("TotalAlloc", toKb(ms.TotalAlloc))
	//
	//// HeapObjects は、ヒープに割り当てられているオブジェクトの数
	//output.Stdoutl("HeapObjects", toKb(ms.HeapObjects))
	//
	//// Sys は、OSから割り当てられたメモリの合計量
	//output.Stdoutl("Sys", toKb(ms.Sys))
	//
	//// NumGC は、実施されたGCの回数
	//output.Stdoutl("NumGC", ms.NumGC)

	// Alloc は、現在ヒープに割り当てられているメモリ
	fmt.Println("Alloc", toKb(ms.Alloc))

	// HeapAlloc と同じ
	fmt.Println("HeapAlloc", toKb(ms.HeapAlloc))

	// TotalAlloc は、ヒープに割り当てられたメモリ量の累積
	// Allocと違い、こちらは増えていくが減ることはない
	fmt.Println("TotalAlloc", toKb(ms.TotalAlloc))

	// HeapObjects は、ヒープに割り当てられているオブジェクトの数
	//fmt.Println("HeapObjects", toKb(ms.HeapObjects)) //正しく動いてるかわからない

	// Sys は、OSから割り当てられたメモリの合計量
	fmt.Println("Sys", toKb(ms.Sys))

	// NumGC は、実施されたGCの回数
	//fmt.Println("NumGC", toKb(ms.NumGC)) //正しく動いてるかわからない
}

func toKb(bytes uint64) uint64 {
	return bytes / 1024
}

//noinspection GoUnusedFunction
func toMb(bytes uint64) uint64 {
	return toKb(bytes) / 1024
}