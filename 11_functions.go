package main
import "fmt"
func add(a int, b int) int{
    return a+b
}

func main() {

    c  := add(2, 3)
    fmt.Println(c)
}