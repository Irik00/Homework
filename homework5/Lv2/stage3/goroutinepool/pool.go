package goroutinepool

import "sync"

type Task interface{
	Execute()
}

type Pool struct{
	TaskChan chan Task
	wg sync.WaitGroup
}

func Newpool(Num,Cap int)*Pool{
	p := &Pool{
		TaskChan: make(chan Task,Cap),
	}
	for range Num{
		p.wg.Add(1)
		go func(){
			defer p.wg.Done()
			for task := range p.TaskChan{
				task.Execute()
			}
		}()
	}
	return p
}

func(p *Pool)Wait(){
	close(p.TaskChan)
	p.wg.Wait()
}