package http

type Request struct {
  Type string
  NumLines int
  Length int
  RequestLine string
  Headers map[string]string
  Body string
}

func Parse(packet string) (*Request) {
  lines := extractLines(packet)

  return &Request{
    Length: len(packet),
    Type: extractType(packet),
    NumLines: len(lines),
    RequestLine: extractRequestLine(lines),
    Headers: extractHeaders(lines),
    Body: extractBody(packet),
  }
}
