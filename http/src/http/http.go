package http

import (
  "fmt"
  "strings"
)

func Parse(packet string) {
  fmt.Println(strings.Split(packet, "k"))
  _ = "breakpoint"
  fmt.Println(packet)
}
