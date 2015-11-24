package utils

import (
	"testing"
)



const (
	// I'm being funny
	url = "https://godoc.org/golang.org/x/net/html"
)



func TestParseLink(t *testing.T) {
	println("Testing parse link for " + url)
	ParseLink(url)
}