package main

import (
  "fmt"
  "strings"
)
func main() {
  var words [3]string
  words[0] = "Hello"
  words[1] = "World!"
  words[2] = "世界！"

  fmt.Println(strings.Trim(words[0], "]"))
  fmt.Println(words[0])
  
  var slice = words[0:2]
  fmt.Println(slice)
  slice[1] = "Hell-World!"
  fmt.Println(slice)
  fmt.Println(words[0], words[1], words[2])
}
