package urlJoin

import (
	"fmt"
	"net/url"
)

func UrlJoin(paths ...string) (*url.URL, error) {
	fmt.Println("Test")
	return url.Parse("https://test.com")
}
