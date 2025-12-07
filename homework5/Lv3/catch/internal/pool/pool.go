package pool

import "sync"

//任务的模版
type Task struct {
	Filepath string
	Keyword  string
}

//结果的模版
type Result struct {
	Filepath string
	LineNum  int
	Content  string
}

//工作池
type Pool struct {
	TaskChan   chan Task //发送任务的通道
	ResultChan chan Result //接收结果的通道
	Wg         sync.WaitGroup
}

func NewPool(workerNum int) *Pool {
	//初始化
	p := &Pool{
		TaskChan:   make(chan Task, 100),
		ResultChan: make(chan Result, 100),
	}
	//接着启动
	for i := 0; i < workerNum; i++ {
		p.Wg.Add(1)
		go p.worker()
	}
	return p
}

//工作函数
func (p *Pool) worker() {
	defer p.Wg.Done()

}

//关闭池
func (p *Pool) Close() {
	//close(p.TaskChan)
	p.Wg.Wait()
	//close(p.ResultChan)
	//在root中已经关闭通道，不要重复关
}
