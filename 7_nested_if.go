package main
import "fmt"

func main(){
    var water = 0;

    if(water>20) {
        fmt.Println("Fill Water to 40")
    } else if (water>40){
        fmt.Println("Water is full!")
    } else if (water==0){
        fmt.Println("Please buy some water!")
    } else{
        fmt.Println("Do Nothing!")
    }
}