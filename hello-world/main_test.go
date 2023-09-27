package main

import (
	"testing"
)

func TestHandler(t *testing.T) {
	t.Run("returns a response with a body containing 'Hello World! I am written in Go.'", func(t *testing.T) {
		response, err := handler()
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}

		if response.StatusCode < 200 {
			t.Errorf("Expected status code >= 200, but got %v", response.StatusCode)
		}

	})

}
