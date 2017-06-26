package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/axgle/mahonia"
	"github.com/golang/glog"
)

func Utf8toGBK(utfStr string) string {

	var enc mahonia.Encoder
	enc = mahonia.NewEncoder("gbk")
	gbkStr := enc.ConvertString(utfStr)
	return gbkStr
}

func GBKtoUtf8(gbkStr string) string {
	var dec mahonia.Decoder

	dec = mahonia.NewDecoder("gbk")
	utfStr := dec.ConvertString(gbkStr)
	return utfStr
}

type CQSSC struct {
	BuyHost  string
	TimeHost string
	Last50   [50]string
	Last100  [100]string
	Last1    string
	Numbers  map[string]string
}

func (c *CQSSC) Init() {
	CookieInit()
	urlstr := "http://buy.cqcp.net/game/cqssc/"

	_, err := HttpGet(urlstr)
	if err != nil {
		glog.Errorln("Login Error:", err)
	}
}

func (c *CQSSC) GetLastNum() (string, error) {

	urlstr := c.BuyHost + "?itype=11&name=0.02410333935306941"
	fmt.Println(urlstr)
	data, err := HttpGet(urlstr)
	if err != nil {
		glog.Errorln("ERR:", err, urlstr)
	}

	return GBKtoUtf8(string(data)), err

}

func (c *CQSSC) GetLast10() {
	//+?itype=3
}

//GET http://buy.cqcp.net/trend/ssc/scchart_11.aspx HTTP/1.1
//Host: buy.cqcp.net
//Connection: keep-alive
//Upgrade-Insecure-Requests: 1
//User-Agent: Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36
//Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
//Accept-Encoding: gzip, deflate, sdch
//Accept-Language: zh-CN,zh;q=0.8

func (c *CQSSC) GetLastNumByCount(count int) (string, error) {
	urlstr := "http://buy.cqcp.net/trend/ssc/scchart_11.aspx"
	data, err := HttpGet(urlstr)
	if err != nil {
		glog.Errorln(err)
		return "", err
	}
	return string(data), err

}

func (c *CQSSC) GetTime() (string, error) {
	//fmt.Println("Timehost:", c.TimeHost)
	data, err := HttpGet(c.TimeHost)
	if err != nil {
		glog.Errorln(err)
		return "", err
	}
	fmt.Println(data)
	return string(data), err
}

func (c *CQSSC) Run() error {
	c.Init()
	last50 := make(map[string]string)
	data, err := c.GetLastNumByCount(0)
	if err != nil {
		panic(err)
	}
	htmlstr := string(data)

	ret := strings.Index(htmlstr, "Con_BonusCode")
	numstr := htmlstr[ret:]
	retend := strings.Index(numstr[0:], "\";")

	//fmt.Println(numstr[0+17 : retend])
	numbers := strings.Split(numstr[0+17:retend], ";")
	//fmt.Println(numbers)

	for _, r := range numbers {
		ns := strings.Split(r, "=")
		//fmt.Println(ns[0], ns[1])
		last50[ns[0]] = ns[1]
	}
	fmt.Println(last50)

	for {
		lastnum, err := c.GetLastNum()
		if err != nil {
			glog.Errorln("GetLastNum:", err, string(lastnum))
		}
		fmt.Println("LastNum:", lastnum)
		ln := strings.Split(lastnum, "|")

		if v, ok := last50[ln[1]]; ok {
			fmt.Println(v, ok)
		} else {
			fmt.Println("New:", ln[0], ln[1])
			last50[ln[0]] = ln[1]
		}

		/*
			timestr, err := c.GetTime()
			if err != nil {
				glog.Errorln("GetTime:", err)
			}
			fmt.Println(lastnum, timestr)
		*/

		fmt.Println(last50)
		time.Sleep(time.Second * 15)

	}
	return nil
}
