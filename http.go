package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var Headers map[string]string

type Jar struct {
	cookies []*http.Cookie
}

func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.cookies = cookies
}
func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
	return jar.cookies
}

var Global_jar *Jar

func CookieInit() {
	Global_jar = new(Jar)
	Headers = map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0", "Accept": "*/*", "Accept-Language": "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3"}
	//Headers = map[string]string{"User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0", "Accept": "*/*", "Accept-Language": "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3", "Accept-Encoding": "gzip, deflate", "Content-Type": "application/x-www-form-urlencoded", "X-Requested-With": "XMLHttpRequest", "Connection": "keep-alive"}
}

func HttpGet(urlStr string) ([]byte, error) {
	client := &http.Client{nil, nil, Global_jar, 99999999999992}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		fmt.Println("Get error", err)
		return nil, err
	}
	//set header
	for k2, v2 := range Headers {
		req.Header.Set(k2, v2)
	}
	fmt.Println(req.Header)
	
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Connect error:", urlStr, err)
		return nil, err
	}
	fmt.Println("返回值：", resp.StatusCode, urlStr)
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Println(resp.ContentLength)
	defer resp.Body.Close()
	return content, err
}
