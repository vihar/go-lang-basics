package main

import "fmt"

const name string = "Keats"
func main(){
    fmt.Println(name)
    const n = 500000000
    const a = 3e20 / n
    fmt.Println(a)
    fmt.Println(int64(a))
}