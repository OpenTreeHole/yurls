# xurls

Extract urls from text using regular expressions. Requires Go 1.22 or later.

```go
import "github.com/opentreehole/yurls"

func main() {
	findUrls := xurls.Parse("https://example.com foo bar") // []string{"https://example.com"}
}
```

