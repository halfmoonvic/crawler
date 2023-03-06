package parser

import (
	"regexp"

	"github.com/halfmoonvic/crawler/engine"
)

const cityListRe = `<a href="(http://localhost:8888/mock/www.zhenai.com/zhenghun/[a-z0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		// fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFunc: engine.NilParser})
	}

	return result
}
