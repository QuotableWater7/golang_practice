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

  output := extractHeaders(lines)
  fmt.Println(strings.Join(output, "\n"))
}
