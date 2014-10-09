package search

import (
	"fmt"
	"testing"
	"time"

	"appengine/aetest"
)

type Doc struct {
	Author  string
	Comment string
}

func TestSearchCursor(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	index, err := Open(fmt.Sprintf("TestIndex%s", time.Now().Format(time.RFC3339)))
	if err != nil {
		t.Fatalf("Unable to open index %v", err)
	}
	for i := 0; i < 10; i++ {
		d := &Doc{
			Author:  fmt.Sprintf("User %d", i),
			Comment: fmt.Sprintf("Comment %d", i),
		}
		id, err := index.Put(c, fmt.Sprintf("user-%d", i), d)
		if err != nil {
			t.Errorf("Unable to index document %v: %v", d, err)
		}
		err = index.Get(c, id, d)
		if err != nil {
			t.Errorf("Unable to get by id %v: %v", id, err)
		}
	}

	var cursor string
	var done bool
	var loopCount = 0
	var totalResCount = 0
	for !done {
		opts := &SearchOptions{
			Limit:  3,
			Cursor: cursor,
		}
		resCount := 0
		for it := index.Search(c, "user", opts); ; {
			var res Doc
			id, err := it.Next(&res)
			if err == Done {
				if resCount < 3 {
					done = true
				}
				cursor = it.Cursor()
				t.Logf("Using cursor %v", cursor)
				break
			}
			if err != nil {
				t.Errorf("Error fething results %v", err)
				break
			}
			resCount++
			totalResCount++
			t.Logf("Search result %s -> %#v", id, res)
		}
		if loopCount++; loopCount > 10 {
			break
		}
	}

	if totalResCount != 10 {
		t.Errorf("Invalid results returned when using cursor %v, expected 10", totalResCount)
	}
}
