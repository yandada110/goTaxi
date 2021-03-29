package controller

import (
	"encoding/json"
	"fmt"
	"goTaxi/models/order"
	"goTaxi/pkg/util"
	"goTaxi/pkg/util/tcp"
	"time"
)

var (
	StartPosition, EndPosition string
	orders                     order.Order
)

func (n NetConn) CreateOrder() {
	// 用于记录用户输入的选项
	fmt.Println("----------------下单打车界面------------")
	fmt.Println("请输入初始位置:")
	fmt.Scanf("%s\n", &StartPosition)
	fmt.Println("请输入目标位置:")
	fmt.Scanf("%s\n", &EndPosition)
	//创建订单
	orders = order.Order{
		OrderStart:   StartPosition,
		OrderEnd:     EndPosition,
		OrderOn:      util.Generate(),
		OrderStatus:  order.StatusCreate,
		CreatedStart: time.Now(),
	}
	jsons, _ := json.Marshal(orders)
	//初始化实例
	n.Transfer = tcp.NewTransfer(&n.Conn)
	//发送创建请求给服务端
	err := n.Transfer.IWrite(tcp.TYPE_CREATE_ORDER, tcp.CodeSuccess, jsons)
	if err != nil {
		return
	}
}

//取消订单
func (n NetConn) CancelOrder() {
	orders = order.Order{
		OrderOn:      orders.OrderOn,
	}
	jsons, _ := json.Marshal(orders)
	//初始化实例
	n.Transfer = tcp.NewTransfer(&n.Conn)
	//发送创建请求给服务端
	err := n.Transfer.IWrite(tcp.TYPE_CANCEL_ORDER, tcp.CodeSuccess, jsons)
	if err != nil {
		return
	}
}
