package main

import "fmt"

func main() {
    var value1 int = 1
    fmt.Println("Hello World!", value1)
    var input string
    fmt.Println("Type something:")
    fmt.Scan(&input)
    fmt.Printf("You typed: \"%v\"", input)
}
