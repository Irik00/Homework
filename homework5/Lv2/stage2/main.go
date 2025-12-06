package main

import(
	"fmt"
	"sync"
	"sync/atomic"
)

//任务是自增
type Task struct{
	sum *atomic.Int64
	wg *sync.WaitGroup
}

//能执行自增操作的方法
func(t *Task) Execute(){
	defer t.wg.Done()
	t.sum.Add(1)
}

func main(){
	sum := atomic.Int64{}
	wg := sync.WaitGroup{}
	sum.Store(0)
	wg.Add(100000)
	taskchan := make(chan Task,100000)
	for range 10{
		go func(){
			for task := range taskchan {
				task.Execute()
			}
		}()
	}

	for range 100000{
		taskchan <- Task{
			sum: &sum,
			wg: &wg,
		}
	}
	close(taskchan)
	wg.Wait()
	fmt.Println(sum.Load())
}

//atomic.Int64和sync.WaitGroup都禁止值拷贝
// func (wg *sync.WaitGroup) Add(delta int)
// func (wg *sync.WaitGroup) Done()
// func (wg *sync.WaitGroup) Go(f func())
// func (wg *sync.WaitGroup) Wait()