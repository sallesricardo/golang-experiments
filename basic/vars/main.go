package main

import (
	"fmt"
	"math"
	//"math/rand"
    "unicode/utf8"
)

func main() {
    var value1 int = 1
    var vInt8 int8 = 127
    var vInt16 int16 = 32767
    var vInt32 int32 = math.MaxInt32
    var vInt64 int64 = math.MaxInt64
    var uValue1 uint = 1
    var vUint8 uint8 = 255
    var vUint16 uint16 = 65535
    var vUint32 uint32 = math.MaxUint32
    var vUint64 uint64 = math.MaxUint64

    fmt.Println("Int types:")
    fmt.Println(value1, vInt8, vInt16, vInt32, vInt64, uValue1, vUint8, vUint16, vUint32, vUint64)

    var vFloat32 float32 = math.MaxFloat32
    var vFloat64 float64 = math.MaxFloat64
    fmt.Println("Float types")
    fmt.Println(vFloat32, vFloat64)

    fmt.Println("Strings type:")
    var stringUtf8 string = "My String with accent ã é ç"
    fmt.Println(stringUtf8, len(stringUtf8), utf8.RuneCountInString(stringUtf8))

    var char rune = 'a'
    fmt.Println(char)

    var vBool bool = true
    fmt.Println("Boolean type:")
    fmt.Println(vBool)

    const pi float32 = 3.1415
    fmt.Println(pi)

    fmt.Println("Hello World!", value1)
    var input string
    fmt.Println("Type something:")
    fmt.Scan(&input)
    fmt.Printf("You typed: \"%v\"", input)
    // var value = rand.Rand(100)
    // fmt.Println(value)
}
