package hkit

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasics(t *testing.T) {
	var b bytes.Buffer

	markup := Div(&DivProps{ID: "test1"}, Div(nil, Text("Test")))
	_ = markup.Render(&b)
	assert.Equal(t, `<div id="test1"><div>Test</div></div>`, b.String())

	b.Reset()

	markup = Div(&DivProps{ID: "test2"},
		Div(nil, Text("Test2")),
	)

	expectedIndent := `<div id="test2">
    <div>
        Test2
    </div>
</div>`

	_ = markup.RenderIndent(&b, "", "    ")
	assert.Equal(t, expectedIndent, b.String())
}

func TestHFuncBasics(t *testing.T) {
	var b bytes.Buffer

	markup := H(
		Div, &DivProps{ID: "test2"},
		H(
			Div, nil,
			Text("H func test"),
		),
	)

	expectedIndent := `<div id="test2">
    <div>
        H func test
    </div>
</div>`

	_ = markup.RenderIndent(&b, "", "    ")
	assert.Equal(t, expectedIndent, b.String())
}

func TestCustomComponent(t *testing.T) {
	var b bytes.Buffer

	comp1 := func(text string) Component {
		return Div(nil, Main(nil,
			P(nil, Text(text)),
		))
	}

	markup := Div(&DivProps{ID: "comp_test"}, comp1("comp1 test"))
	_ = markup.Render(&b)
	assert.Equal(t, `<div id="comp_test"><div><main><p>comp1 test</p></main></div></div>`, b.String())

	b.Reset()

	comp2 := func(text string) Component {
		return Html(nil, Body(nil, comp1(text)))
	}

	markup = comp2("comp2 test")
	_ = markup.Render(&b)
	assert.Equal(t, `<html><body><div><main><p>comp2 test</p></main></div></body></html>`, b.String())
}

func TestSelfClosing(t *testing.T) {
	var b bytes.Buffer

	tt := []struct {
		comp Component
		exp  string
	}{
		{comp: Img(nil), exp: "<img />"},
		{comp: Input(nil), exp: "<input />"},
		{comp: Link(nil), exp: "<link />"},
	}

	for _, tt := range tt {
		b.Reset()
		_ = tt.comp.Render(&b)
		assert.Equal(t, tt.exp, b.String())
	}
}
