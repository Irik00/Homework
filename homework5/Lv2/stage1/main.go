package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

//原子化：要么执行完要么不执行

func main(){
	sum := atomic.Int64{}//"智能锁"
	sum.Store(0)
	num := 10
	everyone := 10000
	for range num{
		go func(){
			for range everyone{
				sum.Add(1)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(sum.Load())
}

