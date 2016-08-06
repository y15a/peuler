package main
import "fmt"

func main() {
  sum := 0
  for c := 3; c < 1000; c++ {
    if c%3 == 0 || c%5 == 0 {
      sum += c
    }
  }
  fmt.Println(sum)
}
