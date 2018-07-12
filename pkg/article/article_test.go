// Copyright (c) 2018 Davsk℠. All Rights Reserved.
// Use of this source code is governed by an ISC License (ISC)
// that can be found in the LICENSE file.

// article_test.go
package article

import (
	"fmt"
	"testing"
)

func TestArticle(t *testing.T) {
	this := New("Title", "Desc")
	fmt.Println(this.Title())
	fmt.Println(this.Description())
	fmt.Println((this.TimeStamp().Year() > 2017))

	// Output:
	// Title
	// Desc
	// true
}

func TestFailure(t *testing.T) {
	t.Fail()
}
