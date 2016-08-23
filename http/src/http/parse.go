package http

import (
  "strings"
)

type Request struct {
  NumLines int
  Length int
  RequestLine string
  Headers [](*Header)
  Body string
}

func Parse(packet string) (*Request) {
  request := Request{}
  request.Length = len(packet)

  lines := strings.Split(packet, "\n")

  request.NumLines = len(lines)
  request.RequestLine = lines[0]

  request.Headers = extractHeaders(lines)
  request.Body = extractBody(packet)

  return &request
}
