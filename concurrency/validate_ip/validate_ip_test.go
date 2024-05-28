package main

import (
	"testing"
)

func TestIsValidIP(t *testing.T) {
	tests := []struct {
		ip       string
		expected bool
	}{
		{"192.168.1.1", true},
		{"test.test", true},
		{"", false},
	}
	for _, test := range tests {
		parsedIp := isValidIP(test.ip)
		if parsedIp != test.expected {
			t.Errorf("parsed Ip(%s): %t, expected: %t", test.ip, parsedIp, test.expected)
		}
	}
}
