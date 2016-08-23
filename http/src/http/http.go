package http

import (
  "fmt"
  "strings"
)

func Parse(packet string) {
  lines := strings.Split(packet, "\n")
  request_line := lines[0]

  fmt.Printf("\n%d lines in packet.\n", len(lines))
  fmt.Printf("*******************\n\n")
  fmt.Println("REQUEST LINE")
  fmt.Println(request_line)

  fmt.Println()
  fmt.Println("******************\n")

  headers := extractHeaders(lines)
  for i := range headers {
    fmt.Printf("%s: %s\n", headers[i].key, headers[i].value)
  }
}
