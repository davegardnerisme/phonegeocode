package phonegeocode

import (
	"errors"
	"strings"

	gotrie "github.com/tchap/go-patricia/patricia"
)

var ErrCountryNotFound = errors.New("Could not identify country from phone number.")

type Geocoder interface {
	Country(number string) (string, error)
}

// New initialises a new thread-safe geocoder
func New() Geocoder {
	return &trieGeocoder{
		data: initPrefixes(),
	}
}

type trieGeocoder struct {
	data *gotrie.Trie
}

// Country tries to identify the country for a phone number
func (g *trieGeocoder) Country(number string) (cc string, err error) {
	number = strings.TrimPrefix(number, "+")

	maxLen := -1
	g.data.VisitPrefixes(gotrie.Prefix(number), func(prefix gotrie.Prefix, item gotrie.Item) error {
		if len(prefix) > maxLen {
			cc = item.(string)
		}
		return nil
	})

	if len(cc) == 0 {
		err = ErrCountryNotFound
	}

	return
}
