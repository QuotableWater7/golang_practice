package http

import (
  "regexp"
  "strings"
)

func extractHeaders(lines []string) map[string]string {
  lines_after_request := lines[1:]
  var header_stopping_point int

  for i := range lines_after_request {
    if !isHeader(lines_after_request[i]) {
      header_stopping_point = i
      break
    }
  }

  headers := lines_after_request[:header_stopping_point]
  return buildHeaderMap(headers)
}

func buildHeaderMap(headers []string) map[string]string {
  output := make(map[string]string)

  for i := range headers {
    keyValPair := strings.Split(headers[i], ": ")
    output[keyValPair[0]] = keyValPair[1]
  }

  return output
}

func isHeader(line string) bool {
  isMatch, _ := regexp.MatchString("[A-Z][a-zA-Z]+(-[A-Z][a-zA-Z]+)?:", line)
  return isMatch
}
