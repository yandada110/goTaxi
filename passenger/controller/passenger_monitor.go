package controller

import (
	"encoding/json"
	"fmt"
	"goTaxi/pkg/util/tcp"
)

var (
	ScanStatus chan string
	ScanChat chan string
)

func (n *NetConn) ServiceMonitor() {
	ScanStatus = make(chan string, 50)
	ScanChat = make(chan string, 20)
	var (
		orderCreateSuccess tcp.OrderCreateSuccess
	)
	//监听乘客事件
	for {
		mess, err := n.Transfer.IRead()
		if err != nil {
			return
		}
		if mess.Code != tcp.CodeSuccess {
			fmt.Println(mess.Data)
			return
		}
		switch mess.Types {
		case tcp.TYPE_CREATE_ORDER:
			err := json.Unmarshal([]byte(mess.Data), &orderCreateSuccess)
			if err != nil {
				return
			}
			fmt.Println("您下单:", orderCreateSuccess.Order.OrderStart, "-", orderCreateSuccess.Order.OrderEnd, "-", orderCreateSuccess.Msg)
			ScanStatus <- "operation"
			break
		case tcp.TYPE_CANCEL_ORDER:
			fmt.Println("您的订单:", mess.Data)
			ScanStatus <- "operation"
			break
		case tcp.TYPE_RECEIVING_ORDER:
			fmt.Println("-----订单已接------")
			fmt.Println("\t 1 聊天")
			fmt.Println("\t 2 取消订单")
			ScanChat <- "chat"
			break
		case tcp.TYPE_CHAT_PASSENGER:
			fmt.Println("司机跟你说:",mess.Data)
			fmt.Println("----------------等待司机接单------------")
			fmt.Println("\t 1 取消订单")
			fmt.Println("\t 2 继续等待")
			fmt.Println("\t 请选择(1-2):")
		default:
			break
		}
		//监听用户输入的操作行为进行下一步操作
		//switch mess.Types {
		//
		//}
	}
}
