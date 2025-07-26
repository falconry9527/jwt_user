package controllers

import (
	"github.com/gin-gonic/gin"
	"jwt_user/models"
	"jwt_user/msg"
	"jwt_user/services"
	"strconv"
)

func CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		msg.ErrorCode(ctx, msg.ParamError)
		return
	}
	if err := services.CreateUser(&user); err != nil {
		msg.Error(ctx, err)
		return
	}
	msg.Success(ctx, user)
}

func GetUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		msg.ErrorCode(ctx, msg.ParamError)
		return
	}
	user, err := services.GetUserByID(uint(id))
	if err != nil {
		msg.ErrorCode(ctx, msg.NotFind)
		return
	}
	msg.Success(ctx, user)
}

func UpdateUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		msg.ErrorCode(ctx, msg.ParamError)
		return
	}
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		msg.ErrorCode(ctx, msg.ParamError)
		return
	}
	user.ID = uint(id)
	if err := services.UpdateUser(&user); err != nil {
		msg.ErrorCode(ctx, msg.ParamError)
		return
	}
	msg.Success(ctx, user)
}

func DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		msg.ErrorCode(ctx, msg.ParamError)
		return
	}
	if err := services.DeleteUser(uint(id)); err != nil {
		msg.Error(ctx, err)
		return
	}
	msg.Success(ctx, nil)
}

func GetAllUsers(ctx *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		msg.Error(ctx, err)
		return
	}
	msg.Success(ctx, users)
}
