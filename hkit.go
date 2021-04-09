package hkit

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// Render .
func Render(c Component) ([]byte, error) {
	return RenderWithContext(NewContext(), c)
}

// RenderWithContext .
func RenderWithContext(ctx Context, c Component) ([]byte, error) {
	el := c(ctx)

	if r, ok := el.(Renderable); ok {
		return renderToBytes(r.HTML())
	}

	for {
		c, ok := el.(Component)
		if !ok {
			break
		}

		el = c(ctx)
	}

	r, ok := el.(Renderable)
	if !ok {
		return nil, errors.New("element is not renderable")
	}

	return renderToBytes(r.HTML())
}

func renderToBytes(t *html.Node) ([]byte, error) {
	var buf bytes.Buffer

	err := html.Render(&buf, t)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Component .
type Component func(ctx Context) Element

// Element .
type Element interface{}

// Renderable .
type Renderable interface {
	HTML() *html.Node
}

type el struct {
	ctx      Context
	tag      string
	attrs    interface{}
	children []Component
}

func (e *el) HTML() *html.Node {
	if e.tag == "_text_" {
		return &html.Node{
			Type: html.TextNode,
			Data: html.EscapeString(strings.Join(e.attrs.([]string), "")),
		}
	}

	node := &html.Node{
		Type: html.ElementNode,
		Data: e.tag,
		Attr: structToAttrs(e.attrs),
	}

	for _, child := range e.children {
		childEl := child(e.ctx)

		for {
			c, ok := childEl.(Component)
			if !ok {
				break
			}

			childEl = c(e.ctx)
		}

		if childEl == nil {
			continue
		}

		if r, ok := childEl.(Renderable); ok {
			node.AppendChild(r.HTML())
		}
	}

	return node
}

func structToAttrs(s interface{}) []html.Attribute {
	attrs := []html.Attribute{}

	if s == nil {
		return attrs
	}

	val := reflect.ValueOf(s)
	elem := val

	if val.Kind() == reflect.Ptr {
		elem = val.Elem()
	}

	if elem.Kind() != reflect.Struct {
		return attrs
	}

	numfields := elem.NumField()
	elemType := elem.Type()
	for i := 0; i < numfields; i++ {
		field := elem.Field(i)

		if !field.CanSet() {
			continue
		}

		if elem.Field(i).Kind() == reflect.Struct {
			attrs = append(attrs, structToAttrs(field.Interface())...)
			continue

		}

		value := fmt.Sprint(field.Interface())
		if value == "" {
			continue
		}

		name := toKebabCase(elemType.Field(i).Name)
		attrs = append(attrs, html.Attribute{Key: name, Val: value})
	}

	return attrs
}

// inspired by https://gist.github.com/stoewer/fbe273b711e6a06315d19552dd4d33e6
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toKebabCase(str string) string {
	kebab := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	kebab = matchAllCap.ReplaceAllString(kebab, "${1}-${2}")
	return strings.ToLower(kebab)
}
