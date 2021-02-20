# Env
### Environment Variables Helper
##### A small set of routines to make easier and safer to deal with environment variables, in a more structured way

#### Example
```lang=golang
package main

import (
	"fmt"
	"os"

	"github.com/golangsugar/env"
)

func main() {
	fmt.Println(env.AsString("PATH", ""))

	_ = os.Setenv("ENV_N", "115.0")

	fmt.Println(env.AsInt("ENV_N", 0))

}
```
Run at **Go Playground**
https://play.golang.org/p/7pSPZIn5LTT
