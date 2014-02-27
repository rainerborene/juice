package juice

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var Html = `<html><body><h1>It works</h1></body></html>`
var Expected = `<html><head></head><body style="margin: 0;"><h1 style="color: #fff; font-size: 18px; font-size: 21px;">It works</h1></body></html>`

func TestInline(t *testing.T) {
	rules := ParseFile("./examples/style.css")
	inline, err := Inline(strings.NewReader(Html), rules)
	assert.Equal(t, inline, Expected)
	assert.Nil(t, err)
}
