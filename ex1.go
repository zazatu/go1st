package main

import (
    "flag"
    "fmt"
    "math"
)

func main() {
    var (
        name = flag.String("name", "John", "Enter your name.")
        ip   = flag.Int("ip", 12345, "What is your ip?")
    )
    flag.Parse()
    fmt.Println("name:", *name)
    fmt.Println("ip:", *ip)
}