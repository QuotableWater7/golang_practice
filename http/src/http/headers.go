package http

import (
  "regexp"
  "strings"
)

type Header struct {
  Key string
  Value string
}

func extractHeaders(lines []string) [](*Header) {
  lines_after_request := lines[1:]
  var header_stopping_point int

  for i := range lines_after_request {
    if !isHeader(lines_after_request[i]) {
      header_stopping_point = i
      break
    }
  }

  headers := lines_after_request[:header_stopping_point]
  return buildHeaderStructs(headers)
}

func buildHeaderStructs(headers []string) [](*Header) {
  headerStructs := make([](*Header), len(headers))

  for i := range headers {
    keyValPair := strings.Split(headers[i], ": ")
    headerStructs[i] = &Header{Key: keyValPair[0], Value: keyValPair[1]}
  }

  return headerStructs
}

func isHeader(line string) bool {
  isMatch, _ := regexp.MatchString("[A-Z][a-zA-Z]+(-[A-Z][a-zA-Z]+)?:", line)
  return isMatch
}
