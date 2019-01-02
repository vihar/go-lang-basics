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

### Value Operations in GO

In Go, similar to other programming languages we can apply all the arithmetic and logical operations on variables or values. Below is the example how you can use the operations, for arithmetics we can directly use `+`, `-`, `*`, `/` . To perform a logical `and` operation we use `&&`, and to perform a logical `or` we us `||`.

```go
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
```

### Comments in GO

For a programmer, comments help to document and debug things more efficiently, In "Go Lang" we do it the C-Way. Comments in GO are similar to comments in C or C++. We use `//` for single line comments and `/* */` for multiple line comments. Everything that goes inside the astrids doesn't work.

```go
package main

import "fmt"

func main(){

    // fmt.Println("This doesn't print :/")

    fmt.Println("This Works")

    /*
        Code doesn't work her!
    */
}
```

### Constants in GO

Constants in GO can perform arthematics with arbitary precisions. A numeric constant has no type until it’s given one, such as by an explicit cast.

```go
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
```

### Conditionals In GO

The order in which the program carries instructions is called Control Flow. Conditionals are only run if certain conditions are met. If the given condition is true, the specific set of instructions are run, else the other set of instructions are executed. In "Go" we do it condition by using `if` and `else` keywords. Without an `if` there can't be an else statement.

```go
package main
import "fmt"

func main() {
    var emergency = true

    if emergency==true {
        fmt.Println("Apply First-Aid")
    }
    else{
        fmt.Println("Take Some Rest!")
    }
}
```

**You can also have multiple if-statements**, below is the code snippet where we use multiple if-statements to find if the given number is even or odd, if the conditions inside the `if` block isn't satisfied the statements will just be passed on.

```go
package main
import "fmt"

func main(){
    var number = 5

    if number % 2 == 0 {
        fmt.Println("Even Number")
    }

    if number % 2 !=0 {
        fmt.Println("Odd Number")
    }
}
```