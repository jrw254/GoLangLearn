package main

import "fmt"

func main() {
  for i := 1; i <= 100; i++ {
    if i % 3 == 0 {
      fmt.Println(i)
    }
  }

  for ii := 1; ii <= 100; ii++ {
    if ii % 15 == 0 {
      fmt.Println("FizzBuzz")
    } else if ii % 5 == 0 {
      fmt.Println("Buzz")
    } else if ii % 3 == 0 {
      fmt.Println("Fizz")
    } else {
      fmt.Println(ii)      
    } 
  }
    
}  