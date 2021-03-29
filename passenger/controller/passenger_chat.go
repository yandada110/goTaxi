//聊天
package controller

import (
	"fmt"
	"goTaxi/pkg/util/tcp"
)

var Contents string

func (n *NetConn) ChatDriver() {
	fmt.Println("----------------聊天界面------------")
	fmt.Println("\t请输入聊天内容:")
	fmt.Scanf("%s", &Contents)
	//发送创建请求给服务端
	err := n.Transfer.IWrite(tcp.TYPE_CHAT_DRIVER, tcp.CodeSuccess, []byte(Contents))
	if err != nil {
		return
	}
}
