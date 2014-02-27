package juice

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestInline(t *testing.T) {
	html := readFile("./examples/example.html")
	rules := ParseFile("./examples/example.css")
	output := Inline(strings.NewReader(html), rules)
	assert.Contains(t, output, "style")
	assert.NotContains(t, output, "class")
}
