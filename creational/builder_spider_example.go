package main

import (
	"fmt"
)

type SpiderInterface interface {
	Parse()
}

type Builder interface {
	Url(url string) Builder
	UserAgent(userAgent string) Builder
	Method(method string) Builder
	Build() SpiderInterface
}

//////////Spider//////////

type Spider struct {
	url       string
	method    string
	userAgent string
}

//////////BaiduSpider//////////

type BaiduSpider struct {
	Spider
}

func NewBaiduBuilder() Builder {
	return &BaiduSpider{}
}

func (s *BaiduSpider) Url(url string) Builder {
	s.url = url
	return s
}

func (s *BaiduSpider) UserAgent(userAgent string) Builder {
	s.userAgent = userAgent
	return s
}

func (s *BaiduSpider) Method(method string) Builder {
	s.method = method
	return s
}

func (s *BaiduSpider) Build() SpiderInterface {
	return s
}

func (s *BaiduSpider) Parse() {
	fmt.Println("---Baidu Spider Parse Start---")
	fmt.Println("Url:", s.url)
	fmt.Println("Method:", s.method)
	fmt.Println("UserAgent:", s.userAgent)
	fmt.Println("---Baidu Spider Parse End---")
}

//////////GoogleSpider//////////

type GoogleSpider struct {
	Spider
}

func NewGoogleBuilder() Builder {
	return &GoogleSpider{}
}

func (s *GoogleSpider) Url(url string) Builder {
	s.url = url
	return s
}

func (s *GoogleSpider) UserAgent(userAgent string) Builder {
	s.userAgent = userAgent
	return s
}

func (s *GoogleSpider) Method(method string) Builder {
	s.method = method
	return s
}

func (s *GoogleSpider) Build() SpiderInterface {
	return s
}

func (s *GoogleSpider) Parse() {
	fmt.Println("---Google Spider Parse Start---")
	fmt.Println("Url:", s.url)
	fmt.Println("Method:", s.method)
	fmt.Println("UserAgent:", s.userAgent)
	fmt.Println("---Google Spider Parse End---")
}
func main() {
	//baidu spider
	bs := NewBaiduBuilder().Url("http://www.baidu.com").Method("GET").UserAgent("baidu-spider-1.0").Build()
	bs.Parse()

	fmt.Println()

	//google spider
	gs := NewGoogleBuilder().Url("http://www.google.com").Method("GET").UserAgent("google-spider-1.0").Build()
	gs.Parse()
}
