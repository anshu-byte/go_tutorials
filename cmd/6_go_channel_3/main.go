package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3


func main() {
    chickenChannel := make(chan string)
    tofuChannel := make(chan string)
    websites := []string{"walmart.com", "costco.com", "wholefoods.com"}

    for i:= range websites{
        go checkChickenPrices(websites[i], chickenChannel)
        go checkTofuPrices(websites[i], tofuChannel)
    }
    sendMessage(chickenChannel, tofuChannel)
} 

func checkTofuPrices(website string, c chan string){
    for{
        time.Sleep(time.Second*1)
        var tofu_price = rand.Float32() * 20
        if tofu_price < MAX_TOFU_PRICE{
            c <- website
            break
        }
    }
}

func checkChickenPrices(website string, c chan string){
    for{
        time.Sleep(time.Second*1)
        var chicken_price = rand.Float32() * 20
        if chicken_price < MAX_CHICKEN_PRICE{
            c <- website
            break
        }
    }
} 

func sendMessage(chickenChannel chan string, tofuChannel chan string){
    select{
        case website := <-chickenChannel:
            fmt.Printf("Chicken price is low at %v\n", website)
        case website := <-tofuChannel:
            fmt.Printf("Tofu price is low at %v\n", website)    
    }
}