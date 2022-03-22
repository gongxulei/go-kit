/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2021/12/23
 * +----------------------------------------------------------------------
 * |Time: 9:13 下午
 * +----------------------------------------------------------------------
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func main() {
	var urlStr = "postgres://user:pass@host.com:5432/path?name=zhangSan&k=v&k=123#f=2"
	u, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	str,_ := json.Marshal(u)
	fmt.Println(string(str))

	m,err := url.ParseQuery(u.RawQuery)
	str,_ = json.Marshal(m)
	fmt.Println(string(str))

}
