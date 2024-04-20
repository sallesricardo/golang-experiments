package main

import "fmt"

func main() {
    var p *int32 = new(int32)
    var i int32
    fmt.Printf("The value p points to is: %v\n", *p)
    fmt.Printf("The value if i is: %v\n", i)
}
