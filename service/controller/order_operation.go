package controller

import (
	"encoding/json"
	"fmt"
	"goTaxi/models/order"
	"goTaxi/models/user"
	"goTaxi/pkg/database"
	"goTaxi/pkg/redis"
	"goTaxi/pkg/types"
	"goTaxi/pkg/util/tcp"
	"time"
)

func (u *UserConn) CreateOrder(orders *order.Order) {
	orders.PassengerUserId = u.User.Id
	orders.OrderPrice = 10
	orders.Create()
	if orders.Id > 0 {
		jsons, _ := json.Marshal(orders)
		//缓存订单信息
		redis.RDB.HSet(redis.Ctx, order.PushOrderListKey, types.Uint64ToString(orders.Id), string(jsons))
		//加入未接单成员集合
		redis.RDB.SAdd(redis.Ctx, order.PushOrderListKey1, types.Uint64ToString(orders.Id))
		//返回信息给乘客
		orderInfo, _ := json.Marshal(
			tcp.OrderCreateSuccess{
				Order:  *orders,
				Msg:    "下单成功,正在为您呼叫滴滴......",
				Status: true,
			})
		err := u.Transfer.IWrite(tcp.TYPE_CREATE_ORDER, tcp.CodeSuccess, orderInfo)
		if err != nil {
			return
		}
	} else {
		orderInfo := "下单失败,请重新再试"
		err := u.Transfer.IWrite(tcp.TYPE_CREATE_ORDER, tcp.CodeFail, []byte(orderInfo))
		if err != nil {
			return
		}
	}
	return
}

func (u *UserConn) CancelOrder(orders *order.Order) {
	orderInfo := ""
	database.DB.Where("order_on = ?", orders.OrderOn).First(orders)
	if orders.Id <= 0 {
		orderInfo = "该订单不存在"

	}
	if orders.OrderStatus == order.StatusStart || orders.OrderStatus == order.StatusCancel || orders.OrderStatus == order.StatusComplete {
		orderInfo = "该订单不可取消"
	}
	if orderInfo == "" {
		database.DB.Model(order.Order{}).Where("order_on = ?", orders.OrderOn).Update("order_status", order.StatusCancel)
	}
	//移除redis
	redis.RDB.HDel(redis.Ctx, order.PushOrderListKey, types.Uint64ToString(orders.Id))
	redis.RDB.SRem(redis.Ctx, order.PushOrderListKey1, types.Uint64ToString(orders.Id))

	orderInfo = "订单已取消"
	err := u.Transfer.IWrite(tcp.TYPE_CANCEL_ORDER, tcp.CodeSuccess, []byte(orderInfo))
	//查看订单是否有司机接单,如果有,告知司机订单已取消
	if err != nil {
		return
	}
	return
}

func (u *UserConn) GrabbingOrders(orders *order.Order) {
	//移除redis集合成员,移除成功说明抢单,否则失败,
	_, err := redis.RDB.SRem(redis.Ctx, order.PushOrderListKey1, orders.Id).Result()
	if err != nil {
		//告知订单已经被抢
		u.Transfer.IWrite(tcp.TYPE_RECEIVING_ORDER, tcp.CodeFail, []byte("非常抱歉,订单已抢走"))
	}
	//获取订单信息里面的乘客id
	orderInfo, err1 := redis.RDB.HGet(redis.Ctx, order.PushOrderListKey, types.Uint64ToString(orders.Id)).Result()
	if err1 != nil {
		//告知订单已经被抢
		u.Transfer.IWrite(tcp.TYPE_RECEIVING_ORDER, tcp.CodeFail, []byte("非常抱歉,订单已抢走"))
	}
	err = json.Unmarshal([]byte(orderInfo), &orders)
	//移除集合订单信息
	redis.RDB.HDel(redis.Ctx, order.PushOrderListKey, types.Uint64ToString(orders.Id)).Result()

	//然后修改数据库订单信息为已接单
	updateData := map[string]interface{}{
		"driver_user_id": u.User.Id,
		"order_status":   order.StatusReceive,
		"created_start":  time.Now(),
	}
	database.DB.Model(order.Order{}).Where("id = ?", orders.Id).Updates(updateData)

	//绑定乘客id
	UserDriverMemory[u.User.Id].ConnectUserId = orders.PassengerUserId
	//获取用户信息
	PassengerUser, _ := user.GetUserInfo(orders.PassengerUserId)
	//告知用户司机已接订单
	str, _ := json.Marshal(tcp.ReceivingOrdersSuccess{
		User:   *PassengerUser,
		Order:  *orders,
		Msg:    "接单成功",
		Status: true,
	})
	u.Transfer.IWrite(tcp.TYPE_RECEIVING_ORDER, tcp.CodeSuccess, str)

	//获取乘客指针
	userConn := UserPassengerMemory[orders.PassengerUserId]
	//绑定司机id
	userConn.ConnectUserId = u.User.Id
	//告知用户接单成功
	str1, _ := json.Marshal(tcp.ReceivingOrdersSuccess{
		User:   *u.User,
		Order:  *orders,
		Msg:    "司机已接单",
		Status: true,
	})
	userConn.Transfer.IWrite(tcp.TYPE_RECEIVING_ORDER, tcp.CodeSuccess, str1)
	return
}

func PushOrderDriver() {
	fmt.Println("订单信息更新推送给所有司机")
	//循环所有在线司机,有新的订单信息更新,让司机自己在redis中获取所有未接单乘客信息
	for _, v := range UserDriverMemory {
		//只推送未接单的司机
		if v.ConnectUserId == 0{
			v.Transfer.IWrite(tcp.TYPE_PUSH_ORDER, tcp.CodeSuccess, []byte("接单信息有更新"))
		}
	}

}
