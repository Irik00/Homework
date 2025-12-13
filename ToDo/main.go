package main

//把数据库当做结构体操作

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Task结构体

type Task struct {
	gorm.Model
	Title  string `gorm:"not null"`
	Status int    `gorm:"not null"`
}

// add
func AddTask(db *gorm.DB, task *Task) {
	err := db.Create(task).Error
	if err != nil {
		fmt.Println("Add failed!", err)
	} else {
		fmt.Println(" Add success! ")
	}
}

// search
func FindTask(db *gorm.DB) []Task {
	var task []Task
	err := db.Find(&task).Error
	if err != nil {
		fmt.Println(" Find failed: %v ", err)
	} else {
		fmt.Println(" Find success! ", task)
	}
	return task
}

// update
func UpdateTask(db *gorm.DB, task *Task) error {
	return db.Save(task).Error
}

// delete
func DeleteTask(db *gorm.DB, task *Task) error {
	return db.Delete(task).Error
}
func main() {

	//配置信息
	dsn := "root:060509@tcp(localhost:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"
	//连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//处理错误
	if err != nil {
		log.Fatal("connect failed", err)
	}
	//自动创建
	db.AutoMigrate(&Task{})
	//执行任务清单增删改查操作的函数,分装调用吧
	newTask1 := Task{
		Title:  "Daily Task:Sleeping",
		Status: 0,
	}
	newTask2 := Task{
		Title:  "Study Task:Reading",
		Status: 1,
	}
	AddTask(db, &newTask1)
	AddTask(db, &newTask2)

	FindTask(db)

	newTask1.Status = 1
	err1 := UpdateTask(db, &newTask1)
	if err1 != nil {
		fmt.Println("Update Failed!", err)
	} else {
		fmt.Println(" Update success! ")
	}

	err2 := DeleteTask(db, &newTask2)
	if err2 != nil {
		fmt.Println("Delete Failed!", err)
	} else {
		fmt.Println(" Delete success! ")
	}
}


//GORM将go语言结构体自动转变：
//大写转小写，默认变复数


//AutoMigrate 会根据结构体自动修改数据库中的表的结构
//它先检查表是否存在
//存在则更新
//无则创建
