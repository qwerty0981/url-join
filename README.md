# `url-join` A simple library to help construct URLs

How often do you catch yourself wanting to simply append one string to another to create a URL (even though you know
that it is bad practice)? Ever run into problems where you forget if you left a trailing slash as you are building a
URL? If any of these problems sound familiar to you you might find this little library helpful.

This library was inspired by the JavaScript library of the same name [url-join](https://www.npmjs.com/package/url-join)
and it serves basically the same purpose.

## Usage

```golang
import (
    "fmt"

    urlJoin "github.com/qwerty0981/url-join"
)

func main() {
    // Create a url by concating strings together
    fullUrl := urlJoin.Join("https://example.com", "a", "b", "?foo=bar")

    fmt.Println(fullUrl) // https://example.com/a/b?foo=bar
}
```

Check the test cases in the [urlJoin_test](urlJoin_test.go) file to see the expected behavior for different inputs.
