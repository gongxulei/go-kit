/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/6
 * +----------------------------------------------------------------------
 * |Time: 3:19 下午
 * +----------------------------------------------------------------------
 */

package request

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ContentType string

const (
	HTML  ContentType = "text/html; charset=UTF-8"
	XML   ContentType = "text/xml; charset=UTF-8"
	PLAIN ContentType = "text/plain; charset=UTF-8"
	JSON  ContentType = "application/json; charset=UTF-8"
	FORM  ContentType = "application/x-www-form-urlencoded; charset=UTF-8"
	// STREAM 二进制流文件下载
	STREAM ContentType = "application/octet-stream; charset=UTF-8"
	// FORM_DATA 表单文件上传
	FORM_DATA ContentType = "multipart/form-data; charset=UTF-8"
)

// Repaly 发送请求
func Repaly(requestUrl string, method string, param map[string][]string, headerMap map[string]string, contentType ContentType) (response *http.Response, err error) {
	var (
		req         = &http.Request{}
		paramStr    string
		paramByte   []byte
		urlParse    *url.URL
		queryValues url.Values
		// ["name=zhangSan", "age=10"]
		querySlice = make([]string, 0)
	)

	// get 请求参数合并到url
	if method == http.MethodGet && param != nil {
		querySlice = append(querySlice, url.Values(param).Encode())

		urlParse, err = url.Parse(requestUrl)
		if err != nil {
			return
		}
		if urlParse.RawQuery != "" {
			queryValues, err = url.ParseQuery(urlParse.RawQuery)
			if err != nil {
				return
			}
			querySlice = append(querySlice, queryValues.Encode())
		}
		urlParse.RawQuery = strings.Join(querySlice, "&")
		requestUrl = urlParse.String()
	}

	if method == http.MethodPost {
		switch contentType {
		case JSON:
			paramByte, err = json.Marshal(param)
			if err != nil {
				return
			}
			paramStr = string(paramByte)

		case FORM:
			for k, v := range param {
				for _, value := range v {
					queryValues.Set(k, value)
				}
			}
			paramStr = queryValues.Encode()
		}
	}
	req, err = http.NewRequest(method, requestUrl, strings.NewReader(paramStr))
	if err != nil {
		return
	}
	for key, value := range headerMap {
		req.Header.Set(key, value)
	}
	// 发送请求
	response, err = (&http.Client{}).Do(req)
	if err != nil {
		return
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("请求异常,httpCode:" + strconv.Itoa(response.StatusCode))
	}
	return
}
