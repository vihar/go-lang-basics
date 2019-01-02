package main
import "fmt"

func main(){
    var numbers[7] int;

    fmt.Println(numbers)

    numbers[4] = 19

    fmt.Println(numbers)

    var primes = [5]int{2, 3, 5, 7, 11}
    fmt.Println(primes)

    var i = 0
    for i=0; i<5; i++{
        fmt.Println(primes[i])
    }
}