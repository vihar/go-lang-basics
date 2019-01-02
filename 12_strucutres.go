package main
import "fmt"

type person struct {
    name string
    age int
    addr string
    phone int
    access bool
}

func main() {
    fmt.Println(person{"Vihar", 21, "Hyd", 9999, false})

    vihar := person{"Vihar", 21, "Hyd", 9999, false}
    fmt.Println(vihar)

    unknown:= person{"Unknown", 00, "Unknown", 00, false}
    fmt.Println(unknown)
    fmt.Println(unknown.age)

    duplicate := &vihar
    fmt.Println("###############")
    fmt.Println(duplicate.age)
}