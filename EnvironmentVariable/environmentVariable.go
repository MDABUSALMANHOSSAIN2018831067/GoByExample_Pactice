package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {

    os.Setenv("FOO", "1")
	os.Setenv("BAR", "2")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))

    fmt.Println()
    for v, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)

        fmt.Println(v,pair[0])
    }
}