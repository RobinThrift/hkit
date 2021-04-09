package main

import (
	"fmt"
	"log"

	h "github.com/RobinThrift/hkit"
)

func main() {
	ctx := h.NewContext()
	ctx.Set("Title", "From Context")
	ctx.Set("Nested", "Value")

	value, err := h.RenderWithContext(
		ctx,
		layout(compA()),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(value))
}

var compA = func() h.Component {
	return h.Div(
		nil,
		compB(),
	)
}

var compB = func() h.Component {
	return func(ctx h.Context) h.Element {
		return h.Div(
			nil,
			h.Span(
				&h.SpanAttrs{
					Class: "test",
				},
				h.Text(ctx.Get("Nested").(string)),
			),
		)
	}
}

var layout = func(children ...h.Component) h.Component {
	return func(ctx h.Context) h.Element {
		return h.Html(nil,
			h.Head(nil, h.Title(nil, h.Text(ctx.Get("Title").(string)))),
			h.Body(nil, children...),
		)
	}
}
