package main

import (
	"fmt"
	"math/rand"
	"time"
)

func example1() {
    // Channel
    var c = make(chan int)
    go process(c)
    for i := range c {
        fmt.Println(i)
    }
    // Buffer Channel
    var c2 = make(chan int, 5)
    go process(c2)
    for i := range c2 {
        fmt.Println(i)
        time.Sleep(time.Second*1)
    }

}

func process(c chan int) {
    defer close(c)
    for i := 0; i < 5; i++ {
        c <- i
    }
    fmt.Println("Exiting process")
}

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {
    var chickenChannel = make(chan string)
    var tofuChannel = make(chan string)
    var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
    for i := range websites {
        go checkChickenPrices(websites[i], chickenChannel)
        go checkTofuPrices(websites[i], tofuChannel)
    }
    sendMessage(chickenChannel, tofuChannel)

}

func checkChickenPrices(website string, chickenChannel chan string) {
    for {
        time.Sleep(time.Second*1)
        var chickenPrice = rand.Float32() * 20
        if chickenPrice <= MAX_CHICKEN_PRICE {
            chickenChannel <- website
            break
        }
    }
}

func checkTofuPrices(website string, tofuChannel chan string) {
    for {
        time.Sleep(time.Second*1)
        var tofuPrice = rand.Float32() * 20
        if tofuPrice <= MAX_TOFU_PRICE {
            tofuChannel <- website
            break
        }
    }
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
    select {
        case website := <-chickenChannel:
            fmt.Printf("\nFound a deal on chicken at %s\n", website)
        case website := <-tofuChannel:
            fmt.Printf("\nEmail Sent: Found deal on tofu at %s\n", website)
    }
}
