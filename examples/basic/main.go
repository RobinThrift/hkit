package main

import (
	"fmt"
	"log"

	h "github.com/RobinThrift/hkit"
)

func main() {
	value, err := h.RenderIndent(
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
		h.Span(
			&h.SpanAttrs{
				Class: "test",
			},
			h.Text("hello"),
		),
	)
}

var layout = func(children ...h.Component) h.Component {
	return h.Html(nil,
		h.Head(nil, h.Title(nil, h.Text("Hello"))),
		h.Body(nil, children...),
	)
}
