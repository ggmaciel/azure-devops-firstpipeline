package main

import "testing"

func TestHelloWorldLen(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{{
		name: "Hello World Len",
		want: 2,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			helloWorld := "Hello World"
			got := len(helloWorld)

			if got != tt.want {
				t.Errorf("got %d, wanted %d", got, tt.want)
			}
		})
	}
}
