package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "Hello World", want: "Hello from em-ai-wrapper!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
