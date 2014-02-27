package juice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	blocks := Parse("./examples/style.css")

	body := blocks[0]
	h1 := blocks[1]
	margin := body.Properties[0]
	fontSize := h1.Properties[0]

	assert.Equal(t, body.Selector, "body")
	assert.Equal(t, h1.Selector, "h1")

	assert.Equal(t, margin.Name, "margin")
	assert.Equal(t, margin.Value, " 0")

	assert.Equal(t, fontSize.Name, "font-size")
	assert.Equal(t, fontSize.Value, " 21px")
}
