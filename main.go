package main

import "fmt"

func main() {
    ch := make(chan rune)

    go func() {
        for i := 1; i <= 3; i++ {
            ch <- 'w'
        }
        close(ch)
    }()

    for val := range ch {
        fmt.Println(string(val))
    }
}