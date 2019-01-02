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

**Nested if-else**
If you have multiple conditions to check(more than two) you can nest using `else if` statements. This is aslo called as nesting conditionals.

```go
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
```

### Loops in GO

Loops help you repeat stuff, kind of automating things. In Go for loops can be written in different styles. Below are the code snippets how you can write them,

Go lang doesn't have while loops, instead the syntax can be used similarly to that with the `for` keyword. We initialize the variable, give it the condition and then iterate, as soon as the condition turns false, the loop terminates.

```go
var i = 1
for i<5 {
    fmt.Println(i)
    i = i+1
}

```

This is the generic way of writing `for` loops in GO, we first initialize the variable, then give the condition and then iterate it. Until the condition is true the loop keeps on continuing, once the condition is false, it proceeds to the next statements.

```go
for i = 1; i < 5; i++ {
    fmt.Println(i)
}
```

Infinite loops are fun, In Go, we can write an endless loop by merely using `for` and not giving any conditions, by default it takes the condition true leaving the loop to be infinite. You need to break it by hitting `CTRL + C`
```go
for {
    fmt.Println("I am Infinite")
    break
}
```

### Arrays in GO

Arrays store multiple variables of the same datatype. In Go, we declare arrays using "Square Braces" `[]`. We give the type of array, and the number of elements array needs to carry. By default, an array is zero-valued, which for integer types - 0s.

```go
package main
import "fmt"

func main(){
    var numbers[7] int;

    fmt.Println(numbers)
}
```
The above code snippet creates an array of seven zeros.

We can assign the values of elements by using the indices of the arrays, In Go, the indices start with 0, say if we want to change the value in the 4th element we need to use the third index `a[3]=19`.

If we want to manually initialise the elements in the array we need to use `{}`. Below is the code snippet, 

```
package main
import "fmt"

func main(){
    var primes = [5]int{2, 3, 5, 7, 11}
    fmt.Println(primes)
}
```

We can iterate arrays over loops, to find the values in them or want to manipulate anything. Below is the example, 


```go
package main
import "fmt"

func main(){
    var i = 0
    for i=0; i<5; i++{
        fmt.Println(primes[i])
    }
}
```

### Slices in GO

Unlike arrays, slices are typed only by the elements they contain (not the number of elements). We can add elements to arrays by using `append` method. Slices support copying and removing elements too. To create an empty slice with non-zero length, use the builtin `make`. We can use the `len` method to find the total number of elements present in the slices.

In addition to these basic operations, slices support several more that make them richer than arrays. One is the builtin append, which returns a slice containing one or more new values.

```go
package main
import "fmt"


func main() {
    numbers := make([]int, 3)

    fmt.Println(numbers)

    
    numbers = append(numbers, 12)
    numbers = append(numbers, 13, 14)

    fmt.Println(numbers)
}
```

### Functions in GO

Functions are reusable pieces of code, once the functions are declared we can call them a number of times. We have already been using our `main()` function, the main function is invoked as soon as the program is run, we need not explicitly call it. Functions in GO are declared using the keyword `func` followed by the function name. They can take in arguments, and return datatypes. Let's now write our first example function,

```go
package main
import "fmt"
func add(a int, b int) int{
    return a+b
}

func main() {

    c  := add(2, 3)
    fmt.Println(c)
}
```
In the above code snippet, as we mentioned functions could return variables, the return type should be specified after the arguments are given, in the above case the return type of the function is `int`.







