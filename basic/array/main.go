package main

import "fmt"

func main() {
    var intArr[3] int = [...]int{1,2,3}
    fmt.Println(intArr)
    fmt.Println(intArr[0])
    fmt.Println(intArr[1:3])
    fmt.Println(&intArr[1])


    var intSlice []int = []int{3,4,5}
    fmt.Println(intSlice)
    fmt.Printf("The length is %v with capacity of %v\n", len(intSlice), cap(intSlice))
    intSlice = append(intSlice, 7)
    fmt.Printf("The length is %v with capacity of %v\n", len(intSlice), cap(intSlice))

    var intSlice2 []int = []int{8,9}
    intSlice = append(intSlice, intSlice2...)
    fmt.Println(intSlice)

    var intSlice3 []int32 = make([]int32, 3, 8)
    fmt.Printf("The length is %v with capacity of %v\n", len(intSlice3), cap(intSlice3))
}
