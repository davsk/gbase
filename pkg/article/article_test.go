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
