# Chanspy

> Go Channel Spy

[![Goreportcard](https://goreportcard.com/badge/github.com/x1unix/chanspy)](https://goreportcard.com/report/github.com/x1unix/chanspy)
[![Go Reference](https://pkg.go.dev/badge/github.com/x1unix/chanspy.svg)](https://pkg.go.dev/github.com/x1unix/chanspy)

This module contains helpers for Go channel types like:

* Checking if channel is closed without read.
* Obtaining channel length and capacity.
* etc.

## Disclaimer

This package (ab)uses unsafe techniques, use with caution.

> I'm not responsible for broken deployments, segfaults, thermonuclear war, or you getting fired because undefined behavior in the code led to a company going bankrupt.
> Please do some research if you have any concerns about features included in this package before using it!
> YOU are choosing to make these modifications, and if you point the finger at me for messing up your program, I will laugh at you.

## Usage

```go
package main

import (
    "fmt"

    "github.com/x1unix/chanspy"
)

func main() {
    // This is a fast way to check if channel is closed without block.
    // Use chanspy.ValueOf(ch, chanspy.WithLock) for thread-safe way.
    ch := make(chan int)
    fmt.Println(chanspy.IsClosed(ch)) // Prints: false

    close(ch)
    fmt.Println(chanspy.IsClosed(ch)) // Prints: true
}
```

See [examples](./example_test.go) for more details.

