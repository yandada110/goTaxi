package controller

import (
	"goTaxi/models/order"
	"goTaxi/pkg/util/tcp"
)

var (
	orders order.Order
)

func (n NetConn) GrabbingOrders(OrderID string) {
	//初始化实例
	n.Transfer = tcp.NewTransfer(&n.Conn)
	//发送要抢的订单号
	err := n.Transfer.IWrite(tcp.TYPE_GRABBING_ORDER, tcp.CodeSuccess, []byte(OrderID))
	if err != nil {
		return
	}
}
