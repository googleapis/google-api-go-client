// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package iterator_test

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
)

var (
	client *Client
	ctx    = context.Background()
)

var pageTemplate = template.Must(template.New("").Parse(`
<table>
  {{range .Entries}}
    <tr><td>{{.}}</td></tr>
  {{end}}
</table>
{{with .Next}}
  <a href="/entries?pageToken={{.}}">Next Page</a>
{{end}}
`))

// This example demonstrates how to use Pager to support
// pagination on a web site.
func Example_webHandler(w http.ResponseWriter, r *http.Request) {
	const pageSize = 25
	it := client.Items(ctx)
	var items []int
	pageToken, err := iterator.NewPager(it, pageSize, r.URL.Query().Get("pageToken")).NextPage(&items)
	if err != nil {
		http.Error(w, fmt.Sprintf("getting next page: %v", err), http.StatusInternalServerError)
	}
	data := struct {
		Items []int
		Next  string
	}{
		items,
		pageToken,
	}
	var buf bytes.Buffer
	if err := pageTemplate.Execute(&buf, data); err != nil {
		http.Error(w, fmt.Sprintf("executing page template: %v", err), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := buf.WriteTo(w); err != nil {
		log.Printf("writing response: %v", err)
	}
}

// This example demonstrates how to use a Pager to page through an iterator in a loop.
func Example_pageLoop() {
	// Ask for 10 pages of 5 item each from a Prime number iterator.
	const pageSize = 5
	p := iterator.NewPager(Primes(), pageSize, "" /* start from the beginning */)
	for page := 0; page < 10; page++ {
		var items []int
		pageToken, err := p.NextPage(&items)
		if err != nil {
			log.Fatalf("Iterator paging failed: %v", err)
		}
		fmt.Printf("Page %d: %v\n", page, items)
		if pageToken == "" {
			break
		}
	}
	// Output:
	// Page 0: [2 3 5 7 11]
	// Page 1: [13 17 19 23 29]
	// Page 2: [31 37 41 43 47]
	// Page 3: [53 59 61 67 71]
	// Page 4: [73 79 83 89 97]
	// Page 5: [101 103 107 109 113]
	// Page 6: [127 131 137 139 149]
	// Page 7: [151 157 163 167 173]
	// Page 8: [179 181 191 193 197]
	// Page 9: [199 211 223 227 229]
}

// The example demonstrates how to use a Pager to request a page from a given token.
func Example_pageToken() {
	const pageSize = 10
	const pageToken = "1337"
	p := iterator.NewPager(Primes(), pageSize, pageToken)

	var items []int
	nextPage, err := p.NextPage(&items)
	if err != nil {
		log.Fatalf("Iterator paging failed: %v", err)
	}
	fmt.Printf("Primes: %v\nToken:  %q\n", items, nextPage)
	// Output:
	// Primes: [1361 1367 1373 1381 1399 1409 1423 1427 1429 1433]
	// Token:  "1434"
}

// This example demonstrates how to get exactly the items in the buffer, without
// triggering an extra RPC.
func Example_serverPages() {
	// The iterator returned by Primes has a default page size of 20, which means
	// a single call will returns all the primes in the range [2, 21).
	it := Primes()
	var items []int
	for {
		item, err := it.Next()
		if err != nil && err != iterator.Done {
			log.Fatal(err)
		}
		if err == iterator.Done {
			break
		}
		items = append(items, item)
		if it.PageInfo().Remaining() == 0 {
			break
		}
	}
	fmt.Println(items)
	// Output:
	// [2 3 5 7 11 13 17 19]
}

// Primes returns a iterator which returns all known primes (given
// enough time).
func Primes() *SieveIterator {
	it := &SieveIterator{max: 2}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(
		it.fetch,
		func() int { return len(it.items) },
		func() interface{} { b := it.items; it.items = nil; return b })
	return it
}

// SieveIterator is an iterator that returns primes using the sieve of
// Eratosthenes. It is a demonstration of how an iterator might work.
// Internally, it uses "page size" as the number of ints to consider,
// and "page token" as the first number to consider (defaults to 2).
type SieveIterator struct {
	pageInfo *iterator.PageInfo
	nextFunc func() error
	p        []int // Primes encountered so far.
	max      int   // Largest number we've calculated so far.
	items    []int
}

// PageInfo returns a PageInfo, which supports pagination.
func (it *SieveIterator) PageInfo() *iterator.PageInfo { return it.pageInfo }

func (it *SieveIterator) fetch(pageSize int, pageToken string) (string, error) {
	start := 2
	if pageToken != "" {
		s, err := strconv.Atoi(pageToken)
		if err != nil || s < 2 {
			return "", fmt.Errorf("invalid token %q", pageToken)
		}
		start = s
	}
	if pageSize == 0 {
		pageSize = 20 // Default page size.
	}
	x0 := start // x0 is the first number to calculate.
	if x0 > it.max {
		x0 = it.max
	}
	if start+pageSize > it.max {
		it.max = start + pageSize
	}

	// Inefficient but fun way to find primes.
outer:
	for x := x0; x < start+pageSize; x++ {
		for _, y := range it.p {
			switch {
			case x == y:
				// Previously-found prime.
				if x >= start {
					it.items = append(it.items, x)
				}
				continue outer
			case x%y == 0:
				// Not a prime.
				continue outer
			}
		}
		// Found a new prime!
		it.p = append(it.p, x)
		if x >= start {
			it.items = append(it.items, x)
		}
	}

	return strconv.Itoa(start + pageSize), nil
}

func (it *SieveIterator) Next() (int, error) {
	if err := it.nextFunc(); err != nil {
		return 0, err
	}
	item := it.items[0]
	it.items = it.items[1:]
	return item, nil
}
