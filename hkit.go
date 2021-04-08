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
	return xml.Marshal(c)
}

// RenderIndent .
func RenderIndent(c Component) ([]byte, error) {
	return xml.MarshalIndent(c, "", "    ")
}

// Component .
type Component interface {
}

type comp struct {
	tag      string
	attrs    interface{}
	children []Component
}

func (c *comp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	el := xml.StartElement{Name: xml.Name{Local: c.tag}, Attr: structToXMLAttrs(c.attrs)}
	err := e.EncodeToken(el)
	if err != nil {
		return err
	}

	for _, child := range c.children {
		if child.(*comp).tag == "_text_" {
			e.EncodeToken(xml.CharData(child.(*comp).attrs.(string)))
			continue
		}

		child.(*comp).MarshalXML(e, xml.StartElement{})
	}

	return e.EncodeToken(el.End())
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
