package main

import(
	"github.com/Irik00/Lanshan-Go-2025-Homework/Homework/homework5/Lv2/stage3/goroutinepool"
	"fmt"
)

type UseTask struct{
	Num int
}

func (t *UseTask) Execute(){
	fmt.Println("start",t.Num)
}

func main(){
	pool := goroutinepool.Newpool(5,10)
	for i:=1;i<=5;i++{
		pool.TaskChan <- &UseTask{Num:i}
	}
	pool.Wait()
}