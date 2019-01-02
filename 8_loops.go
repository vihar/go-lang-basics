package main
import "fmt"

func main(){
    var i = 1

    for i<5 {
        fmt.Println(i)
        fmt.Println("Looping for the", i,"th time")
        i = i+1
    }

    fmt.Println("Looping is now over!")


    for i = 1; i < 5; i++ {
        fmt.Println(i)
    }

    // for {
    //     fmt.Println("I am Infinite!")
    // }
}