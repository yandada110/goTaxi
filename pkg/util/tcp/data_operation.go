//tcp数据包打包和解析
package tcp

import (
	"encoding/binary"
	"encoding/json"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [4096]byte
}

func NewTransfer(conn *net.Conn) *Transfer {
	return &Transfer{
		Conn: *conn,
	}
}

//打包数据
func Package(types string, code int, data []byte) (jsons []byte, err error) {
	m := Mess{
		Data:  string(data),
		Types: types,
		Code:  code,
	}
	jsons, err = json.Marshal(m)
	return
}

//发送数据
func (t *Transfer) IWrite(types string, code int, data []byte) (err error) {
	var str []byte
	str, err = Package(types, code, data)
	//发送数据长度
	_, err = t.Conn.Write(t.setPkgLen(str))
	if err != nil {
		return
	}
	//发送数据包
	_, err = t.Conn.Write([]byte(str))
	if err != nil {
		return
	}
	return nil
}

//获取数据包数据
func (t *Transfer) IRead() (res Mess, err error) {
	var n int
	//获取数据包长度
	_, err = t.Conn.Read(t.Buf[:4])
	if err != nil {
		return
	}
	pkgLen := t.getPkgLen()
	//获取数据包
	n, err = t.Conn.Read(t.Buf[:pkgLen])
	if err != nil {
		return
	}
	err = json.Unmarshal(t.Buf[:n], &res)
	return
}

func (t *Transfer) setPkgLen(data []byte) []byte {
	// 先定义一个uint32变量
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(t.Buf[:4], pkgLen)
	return t.Buf[:4]

}

func (t *Transfer) getPkgLen() int {
	resPkgLen := binary.BigEndian.Uint32(t.Buf[:4])
	return int(resPkgLen)
}
