package juice

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestInline(t *testing.T) {
	html := readFile("./examples/example.html")
	rules := ParseFile("./examples/example.css")
	inline, err := Inline(strings.NewReader(html), rules)
	assert.Nil(t, err)
	assert.Contains(t, inline, "style")
}
