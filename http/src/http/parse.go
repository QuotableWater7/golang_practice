package http

import (
  "strings"
)

type Request struct {
  Type string
  NumLines int
  Length int
  RequestLine string
  Headers [](*Header)
  Body string
}

func Parse(packet string) (*Request) {
  request := Request{}
  request.Length = len(packet)
  request.Type = extractType(packet)

  lines := strings.Split(packet, "\n")

  request.NumLines = len(lines)
  request.RequestLine = lines[0]

  request.Headers = extractHeaders(lines)
  request.Body = extractBody(packet)

  return &request
}
