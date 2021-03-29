package controller

import (
	"encoding/json"
	"fmt"
	"goTaxi/models/order"
	"goTaxi/pkg/redis"
	"goTaxi/pkg/util/tcp"
)

var ScanStatus chan string

func (n *NetConn) ServiceMonitor() {
	ScanStatus = make(chan string, 50)
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
		case tcp.TYPE_PUSH_ORDER:
			//获取redis中的订单信息显示司机界面中
			orders, err1 := redis.RDB.HGetAll(redis.Ctx, order.PushOrderListKey).Result()
			fmt.Println("-----订单信息已更新-------")
			if err1 != nil {
				fmt.Println("暂无可接订单信息")
			}
			orderInfo := order.Order{}
			for _, v := range orders {
				err := json.Unmarshal([]byte(v), &orderInfo)
				if err == nil {
					fmt.Println("订单id", orderInfo.Id, "\t起点位置:", orderInfo.OrderStart, "\t终点位置:", orderInfo.OrderEnd)
				}
			}
			break
		case tcp.TYPE_RECEIVING_ORDER:
			//接单成功打开成功接单界面
			ScanStatus <- "chat"
			break
		case tcp.TYPE_CHAT_DRIVER:
			//用户发送聊天信息
			fmt.Println("乘客跟你说:",mess.Data)
			fmt.Println("---------------聊天界面------------")
			fmt.Println("请输入聊天信息:")
			break
		default:
			break
		}
	}
}
