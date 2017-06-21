package main

import (
	"fmt"
)

type Spider interface {
	Parse()
}

type SpiderType int

const (
	Baidu SpiderType = 1 << iota
	Google
)

type BaiduSpider struct {
}

func (s *BaiduSpider) Parse() {
	fmt.Println("---BaiduSpider Parse---")
}

type GoogleSpider struct {
}

func (s *GoogleSpider) Parse() {
	fmt.Println("---GoogleSpider Parse---")
}

func NewSpider(t SpiderType) Spider {
	switch t {
	case Baidu:
		return &BaiduSpider{}
	case Google:
		return &GoogleSpider{}
	}
	return nil
}

func main() {
	s := NewSpider(Baidu)
	s.Parse()

	g := NewSpider(Google)
	g.Parse()
}
