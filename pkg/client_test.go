package pkg

import (
	"fmt"
	"net/http"
	"testing"
)

// 测试整个文件：go test -v client_test.go
// 测试单个函数：go test -v client_test.go -test.run TestClient_Send
func TestClient_Send(t *testing.T) {
	t.Run("#Client", func(t *testing.T) {

		var resp ModelResponse
		client, err := NewClient(WithMethod(http.MethodGet), WithUrl(MODEL_URL))
		if err != nil {
			fmt.Println("NewClient Err:", err)
			return
		}
		err = client.Send(&resp)
		if err != nil {
			fmt.Println("Send Err:", err)
			return
		}

		//re, _ := json.Marshal(resp)
		//fmt.Println(string(re))
		for _, item := range resp.Data {
			fmt.Println(item.ID)
		}
	})
}

func TestClient_ImageCreate(t *testing.T) {
	t.Run("#Image-Create", func(t *testing.T) {
		var resp ImageCreateResponse
		client, err := NewClient(WithMethod(http.MethodPost), WithUrl(IMG_CREATE_URL))
		if err != nil {
			fmt.Println("New Client Err:", err)
			return
		}

		err = client.Send(&resp)
		if err != nil {
			fmt.Println("Send Err:", err)
			return
		}

		fmt.Println("Total Image:", len(resp.Data))
		for _, item := range resp.Data {
			fmt.Println("Url:", item.URL)
		}
	})
}
