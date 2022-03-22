/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2021/11/29
 * +----------------------------------------------------------------------
 * |Time: 10:48 下午
 * +----------------------------------------------------------------------
 */

package register

// Service 服务抽象
type Service struct {
	Name  string  `json:"name"`
	Nodes []*Node `json:"nodes"`
}

// Node 服务节点的抽象
type Node struct {
	Id       string            `json:"id"`
	Tags     []string          `json:"tags"`
	MetaData map[string]string `json:"meta_data"`
	Scheme   string            `json:"scheme"`
	Weight   int               `json:"weight"`
	Address  string            `json:"address"`
	Port     int               `json:"port"`
}
