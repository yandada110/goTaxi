package user

import (
	"encoding/json"
	"fmt"
	"goTaxi/pkg/database"
	"goTaxi/pkg/redis"
	"goTaxi/pkg/types"
	"golang.org/x/crypto/bcrypt"
)

//添加用户信息
func (u *User) Create() (status bool, msg string) {
	status = false

	//查询用户是否存在
	if _, err := GetByPhone(u.Phone); err == nil {
		fmt.Println("用户已存在")
		msg = "用户已存在"
		return
	}
	//密码加密
	pwd, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(pwd)
	database.DB.Create(&u)
	return true, "注册成功"
}

//用户登录
func (u *User) Login() (status bool, msg string) {
	status = false
	Password := u.Password
	d := database.DB.Where("phone = ?", u.Phone).First(&u)
	if d.Error != nil {
		fmt.Println("用户名或密码错误")
		msg = "用户名或密码错误"
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(Password)) != nil {
		fmt.Println("用户名或密码错误")
		msg = "用户名或密码错误"
		return
	}
	msg = "登录成功"
	status = true
	return
}

// GetByEmail -
func GetByPhone(phone string) (*User, error) {
	user := &User{}
	d := database.DB.Where("phone = ?", phone).First(&user)
	return user, d.Error
}

func GetUserInfo(id uint64) (*User, error) {
	users, err := redis.RDB.HGet(redis.Ctx, UserPassengerKey, types.Uint64ToString(id)).Result()
	user := &User{}
	if err != nil {
		err = database.DB.Where("id = ?", id).First(&user).Error
	} else {
		err = json.Unmarshal([]byte(users), &user)
	}
	return user, err
}
