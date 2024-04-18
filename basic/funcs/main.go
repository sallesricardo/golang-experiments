package main

import (
    "errors"
    "fmt"
)


func main() {
    var value1, value2 int = 11, 0
    var result, remainder, err = intDivision(value1, value2)
    if (err != nil) {
        fmt.Println(err.Error())
    } else {
        fmt.Printf("Result %v Remainder %v\n", result, remainder)
    }

}

func intDivision(numerator int, denominator int) (int, int, error) {
    var err error
    if (denominator == 0) {
        err = errors.New("Cannot divide by zero")
        return 0, 0, err
    }
    var result int = numerator / denominator
    var remainder int = numerator % denominator

    return result, remainder, err
}
