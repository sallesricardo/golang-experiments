package main

import "fmt"

func main() {
    var myMap map[string]int32 = make(map[string]int32)
    fmt.Println(myMap)

    var myMap2 = map[string]int32{"Adam": 23, "Sarah": 45}
    fmt.Println(myMap2)

    var age, ok = myMap2["Jason"]

    delete(myMap2, "Adam")
    fmt.Println(myMap2)

    if ok {
        fmt.Println("The age is:", age)
    } else {
        fmt.Println("Invalid Name")
    }

    for name := range myMap2 {
        fmt.Println("Name:", name)
    }

    for name, age := range myMap2 {
        fmt.Println("Name:", name, "age:", age)
    }
}
