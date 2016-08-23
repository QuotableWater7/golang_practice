package http

import "strings"

func extractBody(packet string) string {
  emptyLineIndex := strings.LastIndex(packet, "\n\n")

  return packet[emptyLineIndex + 2:]
}
