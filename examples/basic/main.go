package main

import (
	"fmt"
	"log"

	h "github.com/RobinThrift/hkit"
)

func main() {
	value, err := h.Render(
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
		h.Input(&h.InputAttrs{
			Type:  "text",
			ID:    "something",
			Value: "value",
		}),
	)
}

var layout = func(children ...h.Component) h.Component {
	return h.Html(nil,
		h.Head(nil,
			h.Title(nil, h.Text("Hello")),
			h.Link(&h.LinkAttrs{
				Rel:  "stylesheet",
				Href: "test.css",
			}),
		),
		h.Body(nil, children...),
	)
}
