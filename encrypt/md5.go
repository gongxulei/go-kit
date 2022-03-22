/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2021/12/21
 * +----------------------------------------------------------------------
 * |Time: 9:34 上午
 * +----------------------------------------------------------------------
 */

package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"math/rand"
	"time"
)

func Md5V1(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func Md5V2(text string) string {
	ctx := md5.Sum([]byte(text))
	//return fmt.Sprintf("%x", ctx)
	return hex.EncodeToString(ctx[:])
}

func Md5V3(text string) string {
	ctx := md5.New()
	_, _ = io.WriteString(ctx, text)
	return hex.EncodeToString(ctx.Sum(nil))
}

// GetRandomString 生成随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := make([]byte, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
