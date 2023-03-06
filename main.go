package main

import (
	"github.com/halfmoonvic/crawler/engine"
	"github.com/halfmoonvic/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://localhost:8888/mock/www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

    // all, err := fetcher.Get("http://localhost:8888/mock/www.zhenai.com/zhenghun");
    // if err != nil {
    // 	return
    // }

    // // printCityList(all)
    // fmt.Println(parser.ParseCityList(all))
}
