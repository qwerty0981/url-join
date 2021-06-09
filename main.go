package main

import (
	"fmt"

	urlJoin "github.com/qwerty0981/url-join/pkg/url-join"
)

func main() {
	fmt.Println(urlJoin.UrlJoin("https://", "test.com", "test"))
}
