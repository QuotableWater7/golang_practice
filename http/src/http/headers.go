package http

import "regexp"

func extractHeaders(lines []string) []string {
  lines_after_request := lines[1:]
  var header_stopping_point int

  for i := range lines_after_request {
    if !isHeader(lines_after_request[i]) {
      header_stopping_point = i
      break
    }
  }

  return lines_after_request[:header_stopping_point]
}

func isHeader(line string) bool {
  isMatch, _ := regexp.MatchString("[A-Z][a-zA-Z]+(-[A-Z][a-zA-Z]+)?:", line)
  return isMatch
}
