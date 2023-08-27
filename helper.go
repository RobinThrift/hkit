package hkit

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

type Attributes map[string]any

func Attrs(keyVals ...any) Attributes {
	attrs := make(Attributes, len(keyVals))
	for i := 0; i < len(keyVals); i += 2 {
		var v any
		if i+1 < len(keyVals) {
			v = keyVals[i+1]
		}
		if s, ok := keyVals[i].(string); ok {
			attrs[s] = v
		} else {
			attrs[fmt.Sprint(keyVals[i])] = v
		}
	}

	return attrs
}

func (attrs Attributes) writeTo(w io.Writer) error {
	var b bytes.Buffer

	for k, v := range attrs {
		if reflect.ValueOf(v).IsZero() {
			continue
		}

		err := writeAttr(k, v, &b)
		if err != nil {
			return err
		}
	}

	_, err := b.WriteTo(w)

	return err
}

type Classes map[string]bool

func (c Classes) String() string {
	var b bytes.Buffer
	for k, t := range c {
		if t {
			b.WriteString(k)
		}
	}
	return b.String()
}

func Class(classes ...any) string {
	var b bytes.Buffer
	for _, c := range classes {
		switch c := c.(type) {
		case string:
			b.WriteString(c)
		case Classes:
			for k, t := range c {
				if t {
					b.WriteString(k)
				}
			}
		}

	}

	return b.String()
}

func If(cond bool, a, b Component) Component {
	if cond {
		return a
	}

	return b
}
