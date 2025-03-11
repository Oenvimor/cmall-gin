//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 10:58:11
 * @LastEditors: congz
 * @LastEditTime: 2020-08-12 22:32:04
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname  string `form:"nickname" json:"nickname" binding:"required,min=2,max=10"`
	UserName  string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
	Password  string `form:"password" json:"password" binding:"required,min=8,max=16"`
	Challenge string `form:"challenge" json:"challenge"`
	Validate  string `form:"validate" json:"validate"`
	Seccode   string `form:"seccode" json:"seccode"`
}

// Valid 验证表单
func (service *UserRegisterService) Valid() *serializer.Response {
	var code int
	count := 0
	err := model.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count).Error
	if err != nil {
		code = e.ERROR_DATABASE
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if count > 0 {
		code = e.ERROR_EXIST_NICK
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	count = 0
	err = model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count).Error
	if err != nil {
		code = e.ERROR_DATABASE
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if count > 0 {
		code = e.ERROR_EXIST_USER
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() *serializer.Response {
	user := model.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
	}
	code := e.SUCCESS
	// 表单验证
	if res := service.Valid(); res != nil {
		return res
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = e.ERROR_FAIL_ENCRYPTION
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
