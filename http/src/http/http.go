package http

import (
  "fmt"
  "strings"
)

type Request struct {
  RequestLine string
  Headers [](*Header)
}

func Parse(packet string) {
  request := Request{}

  lines := strings.Split(packet, "\n")
  request.RequestLine = lines[0]

  fmt.Printf("\n%d lines in packet.\n", len(lines))
  fmt.Printf("*******************\n\n")
  fmt.Println("REQUEST LINE")
  fmt.Println(request.RequestLine)

  fmt.Println()
  fmt.Println("******************\n")

  request.Headers = extractHeaders(lines)
  for i := range request.Headers {
    fmt.Printf("%s: %s\n", request.Headers[i].key, request.Headers[i].value)
  }
}
