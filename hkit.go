package hkit

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// Render .
func Render(c Component) ([]byte, error) {
	return RenderWithContext(NewContext(), c)
}

// RenderWithContext .
func RenderWithContext(ctx Context, c Component) ([]byte, error) {
	el := c(ctx)

	if _, ok := el.(xml.Marshaler); ok {
		return xml.Marshal(el)
	}

	for {
		c, ok := el.(Component)
		if !ok {
			break
		}

		el = c(ctx)
	}

	return xml.Marshal(el)
}

// RenderIndent .
func RenderIndent(c Component) ([]byte, error) {
	return RenderIndentWithContext(NewContext(), c)
}

// RenderIndentWithContext .
func RenderIndentWithContext(ctx Context, c Component) ([]byte, error) {
	el := c(ctx)

	if _, ok := el.(xml.Marshaler); ok {
		return xml.MarshalIndent(el, "", "    ")
	}

	for {
		c, ok := el.(Component)
		if !ok {
			break
		}

		el = c(ctx)
	}

	return xml.MarshalIndent(el, "", "    ")
}

// Component .
type Component func(ctx Context) Element

// Element .
type Element interface{}

type el struct {
	ctx      Context
	tag      string
	attrs    interface{}
	children []Component
}

func (e *el) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	xmlEl := xml.StartElement{Name: xml.Name{Local: e.tag}, Attr: structToXMLAttrs(e.attrs)}
	err := enc.EncodeToken(xmlEl)
	if err != nil {
		return err
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

		if childEl.(*el).tag == "_text_" {
			enc.EncodeToken(xml.CharData(strings.Join(childEl.(*el).attrs.([]string), "")))
			continue
		}

		childEl.(*el).MarshalXML(enc, xml.StartElement{})
	}

	return enc.EncodeToken(xmlEl.End())
}

func structToXMLAttrs(s interface{}) []xml.Attr {
	attrs := []xml.Attr{}

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
			attrs = append(attrs, structToXMLAttrs(field.Interface())...)
			continue

		}

		value := fmt.Sprint(field.Interface())
		if value == "" {
			continue
		}

		name := toKebabCase(elemType.Field(i).Name)
		attrs = append(attrs, xml.Attr{Name: xml.Name{Local: name}, Value: value})
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
