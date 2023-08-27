package hkit

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHFunc(t *testing.T) {
	var b bytes.Buffer

	markup := Hp(
		"div", Attrs("id", "outer", "class", "h-func-class-outer"),
		H(
			WithProps("div", &DivProps{ID: "nested"}),
			H(
				WithProps(Span, &SpanProps{Class: "h-func-class"}),
				Text("H func test"),
			),
		),
	)

	expectedIndent := `<div id="outer" class="h-func-class-outer">
    <div id="nested">
        <span class="h-func-class">
            H func test
        </span>
    </div>
</div>`

	_ = markup.RenderIndent(&b, "", "    ")
	assert.Equal(t, expectedIndent, b.String())
}
