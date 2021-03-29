package controller

import (
	"fmt"
	"os"
)

var Key string
var Conn = &NetConn{}

func Index() {
index:
	for {
		// 用于记录用户输入的选项
		fmt.Println("----------------司机接单系统------------")
		fmt.Println("\t 1 登录")
		fmt.Println("\t 2 注册")
		fmt.Println("\t 3 退出系统")
		fmt.Println("\t 请选择(1-3):")
		fmt.Scanf("%s", &Key)
		switch Key {
		case "1": //登录
			n, status, err := Conn.Login()
			if err != nil || !status {
				break
			} else {
				n.Operation()
				break index
			}
		case "2": //注册
			n, status, err := Conn.Register()
			if err != nil || !status {
				break
			} else {
				n.Operation()
				break index
			}
		case "3":
			fmt.Println("你退出了系统")
			os.Exit(0)
		default:
			fmt.Println("该选项不存在")
			break
		}

	}
}

func (n *NetConn) Operation() {
operation:
	for {
		// 用于记录用户输入的选项
		fmt.Println("----------------司机接单界面------------")
		fmt.Println("\t  选择接单id")
		fmt.Println("\t r 退出登录")
		fmt.Println("\t q 退出系统")
		fmt.Println("\t 请输入:")
		fmt.Scanf("%s", &Key)
		switch Key {
		case "r":
			n.Conn.Close()
			fmt.Println("这里是退出登录")
			//删除redis缓存数据
			Index()
			break operation
		case "q": //
			fmt.Println("你退出了系统")
			os.Exit(0)
		default:
			//抢单
			n.GrabbingOrders(Key)
			status := <-ScanStatus
			if status == "chat" {
				n.Chat()
				break operation
			}
			break
		}
	}
}

func (n *NetConn) Chat() {
	for {
		fmt.Println("---------------聊天界面------------")
		fmt.Println("\t 1  聊天")
		fmt.Println("\t 2 订单结束")
		fmt.Println("\t 请选择(1-2):")
		fmt.Scanf("%s", &Key)
		switch Key {
		case "1":
			n.ChatPassenger()
		case "2": //
			//订单结束
		default:
			fmt.Println("没有该选项")
			break
		}
	}
}
