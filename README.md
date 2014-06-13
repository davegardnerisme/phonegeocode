# Phone Prefix

Internationalised phone number geocoding by country for Go.

I built this because I needed to turn phone numbers into countries, and that's
really _all_ I needed. If [libphonenumber](https://code.google.com/p/libphonenumber/)
existed for Go, I would probably just use that. AFAIK it doesn't.

This is based on work in [github.com/relops/prefixes](https://github.com/relops/prefixes),
however it has a different implementation, using a [Trie](http://en.wikipedia.org/wiki/Trie)
data structure - specifically [github.com/tchap/go-patricia](https://github.com/tchap/go-patricia).

The way it works is that we have a list of prefixes that identify a country, and
we simply match the _most specific prefix_ to find the country code. This deals
with Canada where the country code is `+1` and shared with US.

All the data lives in a CSV and can be used via codegen to create our Trie.

```
cat data/prefixes.csv | go run data/codegen.go > ./data.go && go fmt
```

## Usage

```
// cc = GB, err = nil
cc, err := phonegeocode.New().Country("+447999111222")

// cc = "", err = phonegeocode.ErrCountryNotFound
cc, err := phonegeocode.New().Country("+999999999998")
```