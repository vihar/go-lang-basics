package main
import "fmt"


func main() {
    numbers := make([]int, 3)

    fmt.Println(numbers)

    
    numbers = append(numbers, 12)
    numbers = append(numbers, 13, 14)

    fmt.Println(numbers)
}