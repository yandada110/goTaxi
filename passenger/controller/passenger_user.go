package controller

import (
	"encoding/json"
	"fmt"
	"goTaxi/models/user"
	"goTaxi/pkg/util/tcp"
	"goTaxi/pkg/verification"
	"net"
)

type NetConn struct {
	Conn     net.Conn
	Transfer *tcp.Transfer
}

var (
	Username, Phone, Password string
)

//注册
func (n *NetConn) Register() (ns *NetConn, status bool, err error) {
	status = false
	fmt.Println("----------------用户注册------------")
	fmt.Println("\t注册:名字 手机号 密码 ")
	fmt.Println("\t返回:r")
	fmt.Println("\t退出系统:q ")
	fmt.Println("\t请选择:")
	fmt.Scanln(&Username, &Phone, &Password)
	//注册验证参数
	if !verification.RegisterVerification(Username, Phone, Password) {
		return
	}
	//连接服务端注册帐号
	n.Conn, err = net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Printf("服务端未开启:%v\n", err)
		return
	}
	defer n.Conn.Close()
	//打包注册数据
	u := user.User{
		UserName: Username,
		Phone:    Phone,
		Password: Password,
		Types:    user.TypePassenger,
	}
	jsons, _ := json.Marshal(u)
	//初始化实例
	n.Transfer = tcp.NewTransfer(&n.Conn)
	//发送创建请求给服务端
	err = n.Transfer.IWrite(tcp.TYPE_USER_REGISTER, tcp.CodeSuccess, jsons)
	//读取服务端创建请求成功之后返回数据状态
	mess, err1 := n.Transfer.IRead()
	if err1 != nil {
		fmt.Println("服务端数据错误")
	}
	if mess.Code != tcp.CodeSuccess {
		fmt.Println(mess.Data)
		return
	}
	var data tcp.UserSuccess
	//解析返回的数据包
	err = json.Unmarshal([]byte(mess.Data), &data)
	if err != nil {
		return
	}
	fmt.Println(data.Msg)
	if data.Status {
		status = true
		//注册成功开启监听服务端数据信息
		go n.ServiceMonitor()
	}else{
		defer n.Conn.Close()
	}
	ns = n
	return
}

//登录
func (n *NetConn) Login() (ns *NetConn, status bool, err error) {
	status = false
	fmt.Println("----------------用户登录------------")
	fmt.Println("\t登录:手机号 密码 ")
	fmt.Println("\t返回:r")
	fmt.Println("\t退出系统:q ")
	fmt.Println("\t请选择:")
	fmt.Scanln(&Phone, &Password)
	//注册验证参数
	if !verification.LoginVerification(Phone, Password) {
		return
	}
	//连接服务端注册帐号
	n.Conn, err = net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Printf("服务端未开启")
		return
	}
	//打包注册数据
	u := user.User{
		Phone:    Phone,
		Password: Password,
	}
	jsons, _ := json.Marshal(u)
	//初始化实例
	n.Transfer = tcp.NewTransfer(&n.Conn)
	//发送创建请求给服务端
	err = n.Transfer.IWrite(tcp.TYPE_USER_LOGIN, tcp.CodeSuccess, jsons)
	//读取服务端创建请求成功之后返回数据状态
	mess, err1 := n.Transfer.IRead()
	if err1 != nil {
		fmt.Println(err1)
	}
	if mess.Code != tcp.CodeSuccess {
		fmt.Println(mess.Data)
		return
	}
	var data tcp.UserSuccess
	//解析返回的数据包
	err = json.Unmarshal([]byte(mess.Data), &data)
	if err != nil {
		return
	}
	fmt.Println(data.Msg)
	if data.Status {
		status = true
		//注册成功开启监听服务端数据信息
		go n.ServiceMonitor()
	}else{
		defer n.Conn.Close()
	}
	ns = n
	return
}
