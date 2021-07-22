package homepage

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {

		})
	}
}
