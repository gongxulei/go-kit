/**
 * +----------------------------------------------------------------------
 * |Created by GoLand.
 * +----------------------------------------------------------------------
 * |User: gongxulei <email:790707988@qq.com>
 * +----------------------------------------------------------------------
 * |Date: 2022/1/6
 * +----------------------------------------------------------------------
 * |Time: 4:41 下午
 * +----------------------------------------------------------------------
 */

package request

import (
	"net/http"
	"reflect"
	"testing"
)

func TestRepaly(t *testing.T) {
	type args struct {
		requestUrl  string
		method      string
		param       map[string][]string
		headerMap   map[string]string
		contentType ContentType
	}
	tests := []struct {
		name         string
		args         args
		wantResponse *http.Response
		wantErr      bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				requestUrl:  "http://127.0.0.1?name=zhangsan#aaa=1",
				method:      "POST",
				param:       map[string][]string{"a": {"cc", "dd"}},
				headerMap:   nil,
				contentType: JSON,
			},
			wantErr:      true,
			wantResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := Repaly(tt.args.requestUrl, tt.args.method, tt.args.param, tt.args.headerMap, tt.args.contentType)
			t.Logf("response: %#v, err: %#v", gotResponse, err)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repaly() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Repaly() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
