package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	reader := csv.NewReader(in)

	fmt.Printf("package phonegeocode\n\n")
	fmt.Printf("import(\n\tgotrie \"github.com/tchap/go-patricia/patricia\"\n)\n")

	fmt.Printf("func initPrefixes() *gotrie.Trie {\n")
	fmt.Printf("\tprefixes := gotrie.NewTrie()\n\n")

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		fmt.Printf("	prefixes.Insert(gotrie.Prefix(\"%s\"), \"%s\")\n", row[0], row[1])
	}
	fmt.Printf("\nreturn prefixes\n}\n")
}
