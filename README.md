# slices

Never write a function to tell whether a string slice contains a string again.

## Usage

```go
package main

import (
    "fmt"

    "github.com/rumpl/slices"
)

func main() {
    fmt.Println(slices.Contains([]string{"a", "b", "c"}, "d")) // Outputs: false
}
```
