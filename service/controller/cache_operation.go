//缓存客户端数据记录
package controller

import (
	"encoding/json"
	"goTaxi/models/user"
	"goTaxi/pkg/redis"
	"goTaxi/pkg/util/tcp"
	"net"
	"time"
)

var (
	//乘客内存
	UserPassengerMemory map[uint64]*UserInfo
	//司机内存
	UserDriverMemory map[uint64]*UserInfo
)

//开辟一块内存储存登录用户信息和指针
func InitCache() {
	UserPassengerMemory = make(map[uint64]*UserInfo, 1024)
	UserDriverMemory = make(map[uint64]*UserInfo, 1024)
}


//用户信息信息
type UserInfo struct {
	UserId        uint64        `json:"userId"`        //存储用户id,用户信息获取可以在redis或者数据库获取
	ConnectUserId uint64        `json:"connectUserId"` //如果这是司机的信息connectUserId为接单成功的用户id,否则相反
	LastTime      time.Time     `json:"lastTime"`      //最后操作时间
	PushTime      time.Time     `json:"pushTime"`      //最后推送时间,司机一分钟最多推送一次,
	Overtime      int           `json:"overtime"`      //超时时间单位分钟
	Conn          *net.Conn     `json:"conn"`          //缓存地址
	Transfer      *tcp.Transfer `json:"transfer"`      //缓存地址
}

func (p *UserInfo) CacheUserInfo(u user.User) {
	switch u.Types {
	case user.TypePassenger:
		UserPassengerMemory[u.Id] = p
		//缓存用户信息到redis
		jsons, _ := json.Marshal(u)
		redis.RDB.HSet(redis.Ctx, user.UserPassengerKey, u.Id, jsons)
		break
	case user.TypeDriver:
		UserDriverMemory[u.Id] = p
		//缓存用户信息到redis
		jsons, _ := json.Marshal(u)
		redis.RDB.HSet(redis.Ctx, user.UserDriverKey, u.Id, jsons)
		break
	default:
		return
	}

}
