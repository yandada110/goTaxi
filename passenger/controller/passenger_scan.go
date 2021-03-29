package controller

import (
	"fmt"
	"os"
)

var Key int
var Conn = &NetConn{}

func Index() {
index:
	for {
		// 用于记录用户输入的选项
		fmt.Println("----------------乘客打车系统------------")
		fmt.Println("\t 1 登录")
		fmt.Println("\t 2 注册")
		fmt.Println("\t 3 退出系统")
		fmt.Println("\t 请选择(1-3):")
		fmt.Scanf("%d", &Key)
		switch Key {
		case 1: //登录
			n, status, err := Conn.Login()
			if err != nil || !status {
				break
			} else {
				n.Operation()
				break index
			}
		case 2: //注册
			n, status, err := Conn.Register()
			if err != nil || !status {
				break
			} else {
				n.Operation()
				break index
			}
		case 3:
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
		fmt.Println("----------------乘客打车界面------------")
		fmt.Println("\t 1 下单打车")
		fmt.Println("\t 2 退出登录")
		fmt.Println("\t 3 退出系统")
		fmt.Println("\t 请选择(1-3):")
		fmt.Scanf("%d", &Key)
		switch Key {
		case 1: //打车
			n.CreateOrder()
			status := <-ScanStatus
			if status == "operation" {
				n.WaitingOrder()
				break operation
			}
		case 2: //
			n.Conn.Close()
			fmt.Println("这里是退出登录")
			//删除redis缓存数据
			Index()
			break operation
		case 3:
			fmt.Println("你退出了系统")
			os.Exit(0)
		default:
			fmt.Println("该选项不存在")
			break
		}
	}
}

func (n *NetConn) WaitingOrder() {
waitingOrder:
	for {
		fmt.Println("----------------等待司机接单------------")
		fmt.Println("\t 1 取消订单")
		fmt.Println("\t 2 继续等待")
		fmt.Println("\t 请选择(1-2):")
		fmt.Scanf("%d", &Key)
		switch Key {
		case 1: //取消订单
			chat := <-ScanChat
			if chat == "chat" {
				n.Chat(Key)
			} else {
				n.CancelOrder()
				status := <-ScanStatus
				if status == "operation" {
					n.Operation()
					break waitingOrder
				}
			}
		case 2:
			fmt.Println("正在为您呼叫滴滴,请耐心等待......")
			break
		default:
			fmt.Println("该选项不存在")
			break
		}
	}
}

func (n *NetConn) Chat(key int) {
	status := false
Chat:
	for {
		if status {
			fmt.Println("---------------司机正常赶往出发地------------")
			fmt.Println("\t 1 聊天")
			fmt.Println("\t 2 取消订单")
			fmt.Println("\t 请选择(1-2):")
			fmt.Scanf("%d", &key)
		} else {
			status = true
		}
		switch key {
		case 1: //聊天
			n.ChatDriver()
		case 2: //取消订单
			n.CancelOrder()
			status := <-ScanStatus
			if status == "operation" {
				n.Operation()
				break Chat
			}
		default:
			fmt.Println("该选项不存在")
			break
		}
	}
}
