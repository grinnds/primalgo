package primalgo

import "testing"

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}
	return obj
}

func TestMorrisonTest(t *testing.T) {
	tests := map[string]struct {
		N        *MorrisonNumber
		expected bool
	}{
		"2147483647-Prime": {N: Must(NewMorrisonNumber(1, 31)), expected: true},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := MorrisonTest(tt.N)
			if got != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}
