package http

import "regexp"

func extractType(packet string) string {
  isResponse, _ := regexp.MatchString(packet, `\AHTTP`)

  if isResponse {
    return "response"
  } else {
    return "request"
  }
}
