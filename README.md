<h1 align="center">👾 LlamaLog - Simple & Fancy 👾</h1>

<h3>Installation:</h3>
    
    go get github.com/LlamaNite/llamalog

<h3>Usage:</h3>

```go
package main

import "github.com/LlamaNite/llamalog"

var log = llamalog.NewLogger("My Project", "Example")

func main() {
	log.Info("Hi! This is %s.", "LlamaNite")
	log.Warn("This is Warning")
	log.Error("oh! error occurred!")
}
```
<h3>Output:</h3>

![image](https://user-images.githubusercontent.com/60406325/128930067-f27121cf-2b10-4f2b-98a8-779ee663d04f.png)

