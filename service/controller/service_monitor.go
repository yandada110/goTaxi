package controller

import (
	"encoding/json"
	"goTaxi/models/order"
	"goTaxi/pkg/types"
	"goTaxi/pkg/util/tcp"
)

//监听客户端事件
func (u UserConn) MonitorPassenger() {
	var (
		orders order.Order
	)
	for {
		mess, err := u.Transfer.IRead()
		if err != nil {
			return
		}
		switch mess.Types {
		case tcp.TYPE_CREATE_ORDER, tcp.TYPE_CANCEL_ORDER:
			//添加到channel中
			err := json.Unmarshal([]byte(mess.Data), &orders)
			if err != nil {
				return
			}
			//加入channel中
			LoopOrder <- &UserConnOrder{
				U:      &u,
				Orders: &orders,
				Types:  mess.Types,
			}
			break
		case tcp.TYPE_GRABBING_ORDER:
			orders.Id = types.StringToUint64(mess.Data)
			//加入channel中
			LoopOrder <- &UserConnOrder{
				U:      &u,
				Orders: &orders,
				Types:  mess.Types,
			}
			break
		case tcp.TYPE_CHAT_DRIVER:
			//获取司机id
			DriverId := UserPassengerMemory[u.User.Id].ConnectUserId
			//获取司机conn
			err := UserDriverMemory[DriverId].Transfer.IWrite(tcp.TYPE_CHAT_DRIVER, tcp.CodeSuccess, []byte(mess.Data))
			if err != nil {
				return
			}
			break
		case tcp.TYPE_CHAT_PASSENGER:
			//获取用户id
			PassengerId := UserDriverMemory[u.User.Id].ConnectUserId
			//获取用户conn
			err := UserPassengerMemory[PassengerId].Transfer.IWrite(tcp.TYPE_CHAT_PASSENGER, tcp.CodeSuccess, []byte(mess.Data))
			if err != nil {
				return
			}
			break
		default:
			break
		}
		//监听用户输入的操作行为进行下一步操作
		//switch mess.Types {
		//
		//}
	}
}
