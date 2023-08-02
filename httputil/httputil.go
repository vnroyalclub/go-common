package httputil

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type huoysClient struct {
	client *http.Client
}

var (
	clients          = make(map[string]*huoysClient)
	errInvaildMethod = errors.New("invalid request method")
	errInvaildUrl    = errors.New("request url is empty")
)

func init() {
	//初始化一个client,重复使用
	//TODO 后续具体配置待定
	normalClient := &huoysClient{
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					//内网使用，不强制校验服务端的证书
					InsecureSkipVerify: true,
				},
				DisableKeepAlives:   false,
				MaxIdleConns:        20,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     60 * time.Second,
			},
			//CheckRedirect:
			//Jar:
			Timeout: time.Second * 30,
		},
	}
	clients["normal"] = normalClient
}

func (c *huoysClient) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

func Request(method, url, body string, header map[string]string) (int, []byte, error) {
	return request(clients["normal"], method, url, body, header)
}

func request(client *huoysClient, method, url, body string, header map[string]string) (code int, res []byte, err error) {
	//判断请求的方法是否正确
	switch method {
	case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete:
	default:
		fmt.Println("[ERROR] invalid request method:", method)
		err = errInvaildMethod
		return
	}

	//判断url是否有效
	if "" == url {
		fmt.Println("[ERROR] request url is empty")
		err = errInvaildUrl
		return
	}

	strings.NewReader(body)
	//创建一个request
	r, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		fmt.Println("[ERROR] internal server error:", err)
	}

	if len(header) > 0 {
		for k, v := range header {
			if "" != k {
				r.Header.Set(k, v)
			}
		}
	}

	//设置代理信息
	r.Header.Set("User-Agent", "huoys")

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println("[ERROR] failed to do respond ,err:", err)
		return
	}

	//限制响应消息体,防止获得恶意攻击信息
	bodyRead := http.MaxBytesReader(nil, resp.Body, 1024*10240)

	code = resp.StatusCode
	defer bodyRead.Close()
	res, err = ioutil.ReadAll(bodyRead)
	if err != nil {
		fmt.Println("[ERROR] failed to get respond body,err:", err)
		return
	}

	return
}
