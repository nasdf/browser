# Browser

Open web pages in the default browser.

```go
import (
    "context"
    
    "github.com/nasdf/browser"
)

func main() {
    browser.Open(context.Background(), "http://localhost:8080")
}
```
