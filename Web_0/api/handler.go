package api

import (
	"Web_0/dao"
	"Web_0/model"
	"Web_0/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "注册失败"})
		return
	}
	if dao.FindUser(user.Username, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user already exist"})
		return
	}
	dao.AddUser(user.Username, user.Password)
	c.JSON(http.StatusOK, gin.H{"message": "Add success"})
}

func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "登录失败"})
		return
	}
	if !dao.FindUser(user.Username, user.Password) {
		c.JSON(http.StatusOK, gin.H{"message": "Not Found"})
		return
	}
	token, err := utils.MakeToken(user.Username, time.Now().Add(5*time.Minute))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "fail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "login success", "token": token})
}

func Ping1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func ReviseKey(c *gin.Context) {
	type ReviseRequest struct {
		Username string `json:"username"`
		Old      string `json:"old"`
		New      string `json:"new"`
	}
	var request ReviseRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数绑定失败"})
		return
	}
	success := dao.RevisePassword(request.Username, request.Old, request.New)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "revise success"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "修改失败"})
}
