package http

import (
  "fmt"
  "strings"
)

func Parse(packet string) {
  lines := strings.Split(packet, "\n")
  request_line := lines[0]

  fmt.Printf("\n%d lines in packet.\n", len(lines))
  fmt.Printf("*******************\n")
  fmt.Println(request_line)
  fmt.Println(strings.Join(lines[1:], "\n"))
}
