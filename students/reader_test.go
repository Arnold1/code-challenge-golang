package students

import "testing"

func TestNewReader(t *testing.T) {
	r, _ := NewReader("data/comma.txt", ',')
	if r != nil {
		t.Error("Expected not nil, got nil")
	}
}
