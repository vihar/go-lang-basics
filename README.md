# Mastering GO Lang!

### The traditional program, "Hello World!"

```go
package main
import "fmt"
func main() {
    fmt.Println("hello world")
}
```
The keyword “import” is used for importing a package into other packages.
In our first program we have imported the package “fmt” for using the function Println. 
Println is the function used to print strings in "go" similar to `print()` in Python.
The package “fmt” comes from the Go standard library. 


### Variables in Go

In Go we declare variables using the keyword `var` followed by the variable name,
the typecasting is done automaticaly else the datatype can be declared after the
variable name.

Go Lang has four main datatypes namely,

* Boolean
* Numeric
* String
* Derived

In each of the datatypes we have different allocations such as (true, false, uint8, uint16...)

In dervied we have special datatypes consisting all pointers, arrays, unions, functions,
slices, interfaces, maps and few other. Below is the code snippet how you can declare variables in "GO"

```go
package main

import "fmt"

func main(){
    var name = "Start"
    fmt.Println(name)

    var number int=1
    fmt.Println(number)

    // Initiliases the number to Zero.
    var a int
    fmt.Println(un_init)

    var bool_value = true
    fmt.Println(bool_value)
}
```