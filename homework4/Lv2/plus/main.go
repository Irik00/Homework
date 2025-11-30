package main

import(
	"fmt"
	"io"
	"os"
	"time"
)

type logLevel int

const(
	levelDebug logLevel = iota
	levelInfo
	levelWarn
	levelError
)

func(l logLevel) String()string{
	switch l {
	case levelDebug:
		return "Debug"//调试
	case levelInfo:
		return "Info"//信息摘要
	case levelWarn:
		return "Warn"//提醒注意
	case levelError:
		return "Error"//故障报告
	default:
		return "mistake"
	}
}

type logger struct{
	writer io.Writer
	level logLevel
}

func getPath()string{
	path := os.Getenv("LOG_PATH")
	//func Getenv(key string) string
	if path == ""{
		path = "test.log"
	}
	return path
}

func newLogger(level logLevel) *logger{
	logPath := getPath()
	f, err := os.OpenFile(logPath,os.O_WRONLY|os.O_APPEND|os.O_CREATE,0644)
	if err != nil{
		fmt.Println("mistake",err)
	}
	multiWriter := io.MultiWriter(f,os.Stdout)//func MultiWriter(writers ...Writer) Writer
	//Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	return &logger{
		writer: multiWriter,
		level: level,
	}
}

func(l *logger) Log(level logLevel,text string){
	if level<l.level{
		return
	}
	now := time.Now()
	Line := fmt.Sprintf(
		"%s %d %s %s\n",
		now.Format("2006-01-02 15:04:05"),
		now.Unix(),
		level.String(),
		text,
	)
	l.writer.Write([]byte(Line))
}

func main(){
	log := newLogger(levelInfo)
//线上环境通常只开启Warn之后的
//Trace,Fatal,Panic
	log.Log(levelDebug,"test")
	log.Log(levelInfo,"user sign in")
	log.Log(levelWarn,"请求参数缺失")
	log.Log(levelError,"数据库连接失败")
}


