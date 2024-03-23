# Chanspy

> Go Channel Spy

This module contains helpers for Go channel types like:

* Checking if channel is closed without read.
* Obtaining channel length and capacity.
* etc.

## Disclaimer

This channel (ab)uses unsafe techniques, use with caution.

> I'm not responsible for broken deployments, segfaults, thermonuclear war, or you getting fired because undefined behavior in the code led to a company going bankrupt.
> Please do some research if you have any concerns about features included in this package before using it!
> YOU are choosing to make these modifications, and if you point the finger at me for messing up your program, I will laugh at you.

## Usage

```
package main

import (
    "fmt"

    "github.com/x1unix/chanspy"
)

func main() {
    ch := make(chan int)
    fmt.Println(chanspy.IsClosed(ch)) // Prints: false

    close(ch)
    fmt.Println(chanspy.IsClosed(ch)) // Prints: true 
}
```

See [examples](./examples).

