package http

import "strings"

func extractLines(packet string) []string {
  return strings.Split(packet, "\n")
}
