package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main(){
//无缓冲

	f1, err1 := os.Create("test_01.txt")//在当前工作目录创建
	if err1 != nil {
		fmt.Println("mistake",err1)
	}

	defer f1.Close()//记得关

	start1 := time.Now()

	for i := 0;i<10000;i++{
		f1.WriteString("test  1\n")
	}
	t1 := time.Since(start1)//自动匹配单位
	fmt.Println("不带缓冲：",t1)
//带缓冲

	f2,err2 := os.Create("test_02.txt")

	if err2 != nil{
		fmt.Println("mistake",err2)
	}

	defer f2.Close()

	w := bufio.NewWriter(f2)

	start2 := time.Now()
	for k:=0;k<10000;k++{
		w.WriteString("test  2\n")
	}
	w.Flush()//刷新
	t2 := time.Since(start2)
	fmt.Println("带缓冲：",t2)
}

//当k=i=5000   
// 不带缓冲： 26.191ms
// 带缓冲： 0s  <1ms显示0

//k=i=10000
// 不带缓冲： 52.4786ms
// 带缓冲： 525.2µs

/*
os.Create  新建空白文件，如有则覆盖重写
func Create(name string) (file *File, err error)
*/

/*
func (f *File) Close() error
*/

/*
type Duration int64    max:290
func Since(t Time) Duration
*/

/*
func NewWriter(w io.Writer) *Writer
*/

