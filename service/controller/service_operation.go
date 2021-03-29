//服务端获取数据分发任务
package controller

import (
	"goTaxi/models/order"
	"goTaxi/pkg/util/tcp"
	"net"
)

var LoopOrder chan *UserConnOrder

type UserConnOrder struct {
	U      *UserConn
	Orders *order.Order
	Types  string //处理类型
}

type NetConn struct {
	Conn net.Conn
}

func InformationAnalysis(conn *net.Conn) {
	//创键user实例
	n := NewSuerProcess(conn)
	mess, err := n.Transfer.IRead()
	if err != nil {
		return
	}
	//解析数据包data
	switch mess.Types {
	case tcp.TYPE_USER_REGISTER:
		n.UserRegister(&mess)
		break
	case tcp.TYPE_USER_LOGIN:
		n.UserLogin(&mess)
		break
	default:
		return
	}

}

//处理channel中的下单和接单数据,
func LoopOrders() {
	for v := range LoopOrder {
		switch v.Types {
		case tcp.TYPE_CREATE_ORDER:
			//下单,下单方法直接调用,监听里面只需要存储channel即可
			v.U.CreateOrder(v.Orders)
			//下单成功通知未接单司机可接单信息更新
			PushOrderDriver()
			break
		case tcp.TYPE_GRABBING_ORDER:
			v.U.GrabbingOrders(v.Orders)
			//抢单成功通知未接单司机可接单信息更新
			PushOrderDriver()
			break
		case tcp.TYPE_CANCEL_ORDER:
			v.U.CancelOrder(v.Orders)
			//取消订单通知未接单司机可接单信息更新
			PushOrderDriver()
			break
		default:
			break
		}
	}
}
