//对应tcp数据包解析
package tcp

import (
	"goTaxi/models/order"
	"goTaxi/models/user"
)

//数据包Types,Source类型
const (
	TYPE_USER_LOGIN      = "login"           //登录
	TYPE_USER_REGISTER   = "register"        //注册
	TYPE_CREATE_ORDER    = "create_order"    //创建订单信息
	TYPE_CANCEL_ORDER    = "cancel_order"    //取消订单
	TYPE_PUSH_ORDER      = "push_order"      //告知司机有新的订单消息
	TYPE_GRABBING_ORDER  = "grabbing_order"  //司机抢单
	TYPE_RECEIVING_ORDER = "receiving_order" //告知司机抢单情况-告知乘客接单情况
	TYPE_CHAT_DRIVER     = "chat_driver"     //用户发送消息给司机
	TYPE_CHAT_PASSENGER  = "chat_passenger"  //司机给用户发送消息

)
const (
	CodeSuccess = 200 //成功
	CodeFail    = 204 //登录失败msg不为空
)

//所有数据包传输结构数据
type Mess struct {
	Code  int    `json:"code"` //状态码
	Data  string `json:"data"` //数据包
	Types string `json:"type"` //请求数据类型
}

//用户操作所需数据
type UserSuccess struct {
	User   user.User `json:"user"`
	Msg    string    `json:"msg"`
	Status bool      `json:"status"` //成功1,失败0
}

//订单创建成功操作所需数据
type OrderCreateSuccess struct {
	Order  order.Order `json:"order"`
	Msg    string      `json:"msg"`
	Status bool        `json:"status"` //成功1,失败0
}

//接单成功推送消息给司机和乘客所需数据
type ReceivingOrdersSuccess struct {
	User   user.User   `json:"user"`  //司机信息
	Order  order.Order `json:"order"` //订单信息
	Msg    string      `json:"msg"`
	Status bool        `json:"status"` //成功1,失败0
}

//type MessagePackage interface {
//	ParsingPackage()
//}

////解析数据包
//func (r *Register) ParsingPackage() {
//	//n := r.u.Phone
//}
