package phonegeocode

import (
	"testing"
)

func TestThingsThatAreFound(t *testing.T) {
	cases := []struct {
		number, country string
	}{
		{"+447999111333", "GB"},
	}

	p := New()

	for _, tc := range cases {
		cc, err := p.Country(tc.number)
		if err != nil {
			t.Errorf("Not expecting number '%s' to yield an error; got %v", tc.number, err)
		}
		if cc != tc.country {
			t.Errorf("Number '%s' did not match expected CC '%s'; got '%s'", tc.number, tc.country, cc)
		}
	}
}
