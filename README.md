# Casper

A little package that brings an animation to your **terminal**

```
go get -u github.com/JAbduvohidov/casper
```

## Example

```go
package main

import (
	"github.com/JAbduvohidov/casper"
	"time"
)

func main() {
	frames := casper.NewFrames([]string{
		"", "h", "he", "hel", "hell", "hello",
		"hello ", "hello w", "hello wo", "hello wor",
		"hello worl", "hello world", "hello world!",
		"hello world!", "hello world!"})

	animation := casper.NewAnimation(frames)

	animation.Play(time.Millisecond)
}

```
Build the code above and run the executable file to see the result
