//用户记录操作
package controller

import (
	"encoding/json"
	"goTaxi/models/user"
	"goTaxi/pkg/util/tcp"
	"net"
	"time"
)

type UserConn struct {
	Conn     *net.Conn
	Transfer *tcp.Transfer
	User     *user.User
}

func NewSuerProcess(conn *net.Conn) *UserConn {
	return &UserConn{
		Conn:     conn,
		Transfer: tcp.NewTransfer(conn),
	}
}

func (u *UserConn) UserRegister(mess *tcp.Mess) {
	var users user.User
	err := json.Unmarshal([]byte(mess.Data), &users)
	status, msg := users.Create()
	//初始化数据
	r := tcp.UserSuccess{
		User:   users,
		Msg:    msg,
		Status: status,
	}
	jsons, _ := json.Marshal(r)
	//注册成功redis记录struct conn链接指针地址和用户数据到redis缓存,
	if status {
		err = u.Transfer.IWrite(tcp.TYPE_USER_REGISTER, tcp.CodeSuccess, jsons)
		if err != nil {
			return
		}
		//创建缓存结构体
		UserInfo := UserInfo{
			UserId:   users.Id,
			LastTime: time.Now(),
			Conn:     u.Conn,
			Transfer: u.Transfer,
		}
		UserInfo.CacheUserInfo(users)
		u.User = &users
		go u.MonitorPassenger()
	} else {
		err = u.Transfer.IWrite(tcp.TYPE_USER_REGISTER, tcp.CodeFail, []byte(msg))
		if err != nil {
			return
		}
	}
}

func (u *UserConn) UserLogin(mess *tcp.Mess) {
	var users user.User
	err := json.Unmarshal([]byte(mess.Data), &users)
	//用户登录
	status, msg := users.Login()
	//初始化数据
	r := tcp.UserSuccess{
		User:   users,
		Msg:    msg,
		Status: status,
	}
	jsons, _ := json.Marshal(r)
	//注册成功redis记录struct conn链接指针地址和用户数据到redis缓存,
	if status {
		err = u.Transfer.IWrite(tcp.TYPE_USER_LOGIN, tcp.CodeSuccess, jsons)
		if err != nil {
			return
		}
		//创建缓存结构体
		UserInfo := UserInfo{
			UserId:   users.Id,
			LastTime: time.Now(),
			Conn:     u.Conn,
			Transfer: u.Transfer,
		}
		UserInfo.CacheUserInfo(users)
		u.User = &users
		go u.MonitorPassenger()
	} else {
		err = u.Transfer.IWrite(tcp.TYPE_USER_LOGIN, tcp.CodeFail, []byte(msg))
		if err != nil {
			return
		}
	}
}
