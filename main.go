package main

import (
    "fmt"
    "os"
)

const BUFSIZE = 1024

func main() {
    file, err := os.Open(`access.log`)
    if err != nil {
        // open error
    }
    defer file.Close()

    buf := make([]byte, BUFSIZE)
    for {
        n, err := file.Read(buf)
        if n == 0 {
            break
        }
        if err != nil {
            // read error
            break
        }

        fmt.Print(string(buf[:n]))
    }
}
