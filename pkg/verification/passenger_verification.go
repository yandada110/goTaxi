//乘客输入验证
package verification

import (
	"fmt"
	"strings"
)

//注册验证
func RegisterVerification(Username, Phone, Password string) (status bool) {
	status = false
	//验证参数
	if (strings.Count(Phone, "") - 1) < 3 {
		fmt.Println("手机号最少不低于3位数")
		return
	}
	if (strings.Count(Password, "") - 1) < 3 {
		fmt.Println("密码最少不低于3位数")
		return
	}
	if (strings.Count(Username, "") - 1) < 3 {
		fmt.Println("昵称最少不低于3位数")
		return
	}
	status = true
	return
}

//登录验证
func LoginVerification(Phone, Password string) (status bool) {
	status = false
	//验证参数
	if (strings.Count(Phone, "") - 1) < 3 {
		fmt.Println("手机号最少不低于3位数")
		return
	}
	if (strings.Count(Password, "") - 1) < 3 {
		fmt.Println("密码最少不低于3位数")
		return
	}
	status = true
	return
}
