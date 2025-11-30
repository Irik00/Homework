package main

import (
	//"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type timestampWriter struct{
	logFile io.Writer
}

func newTimestampWriter(w io.Writer) *timestampWriter{
	return &timestampWriter{logFile:w}
}

func (tw *timestampWriter) Write(p []byte)(n int,err error){
	now := time.Now()
	line := fmt.Sprintf("%s %d %s",now.Format("20006-01-02 15:04:05"),now.Unix(),p)

	return tw.logFile.Write([]byte(line))
}

func main(){
	f, err := os.OpenFile("test.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	if err != nil{
		fmt.Println("mistake ",err)
	}
	defer f.Close()

	logwriter := newTimestampWriter(f)
	fmt.Fprintln(logwriter,"sign in")
	time.Sleep(time.Second*2)
	fmt.Fprintln(logwriter,"hello")
	time.Sleep(time.Second)
	fmt.Fprintln(logwriter,"hi")
}