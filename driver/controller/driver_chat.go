//聊天
package controller

import (
	"goTaxi/pkg/util/tcp"
)

func (n *NetConn) ChatPassenger(Contents string) {
	//发送创建请求给服务端
	err := n.Transfer.IWrite(tcp.TYPE_CHAT_PASSENGER, tcp.CodeSuccess, []byte(Contents))
	if err != nil {
		return
	}
}
