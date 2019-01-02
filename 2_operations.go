package main
import "fmt"

func main(){
    fmt.Println("Let's ask Go the values")
    fmt.Println("Addition", 2+2)
    fmt.Println("Subtraction", 2-2)
    fmt.Println("Multiplication", 2*2)
    fmt.Println("Division", 2/2)

    fmt.Println("Logical AND true, true", true && true)
    fmt.Println("Logical AND true, false", true && false)
    fmt.Println("Logical AND false, true", false && true)
    fmt.Println("Logical AND false, false", false && false)

    fmt.Println("Logical OR true, true", true || true)
    fmt.Println("Logical OR true, false", true || false)
    fmt.Println("Logical OR false, true", false || true)
    fmt.Println("Logical OR false, false", false || false)

    fmt.Println(!true)
}