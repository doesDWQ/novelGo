package util

import (
	"io/ioutil"
	"net/http"
)

// HttpGet 直接使用get获取到网页内容
func HttpGet(url string) (data string, err error) {
	resp, err := http.Get(url)

	if err != nil {
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	// 返回数据为转换为字符串
	data = string(bytes)

	return
}
