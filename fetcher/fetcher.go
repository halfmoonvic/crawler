package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func determineEncoding(r io.Reader) encoding.Encoding {
    bytes, err := bufio.NewReader(r).Peek(1024)
    if err != nil {
    	log.Printf("Fetcher error: %v", err)
    	return unicode.UTF8
    }
    e, _, _ := charset.DetermineEncoding(bytes, "")
    return e
}

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		log.Printf("Fetcher error: %v", err)
	    return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
	    return nil, fmt.Errorf("Wrong status code: %d", resp.StatusCode)
	}

	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}
