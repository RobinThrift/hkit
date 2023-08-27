package hkit

import (
	"bytes"
	"fmt"
	"io"
)

type ComponentFunc[P Props] func(P, ...Component) Component

type Component interface {
	Render(w io.Writer) error
	RenderIndent(w io.Writer, prefix string, indent string) error
}

type Props interface {
	writeTo(w io.Writer) error
}

type tag[P Props] struct {
	name        string
	props       P
	children    []Component
	selfClosing bool
}

func Tag[P Props](name string, props P, children ...Component) Component {
	return &tag[P]{
		name:     name,
		props:    props,
		children: children,
	}
}

func (t *tag[P]) Render(w io.Writer) error {
	if t == nil {
		return nil
	}

	err := writeTag(t.name, t.props, t.selfClosing, w)
	if err != nil {
		return err
	}

	if t.selfClosing {
		return nil
	}

	for _, c := range t.children {
		err = c.Render(w)
		if err != nil {
			return err
		}
	}

	return writeEndTag(t.name, w)
}

func (t *tag[P]) RenderIndent(w io.Writer, prefix string, indent string) error { //nolint: varnamelen
	if t == nil {
		return nil
	}

	_, err := w.Write([]byte(prefix))
	if err != nil {
		return err
	}

	err = writeTag(t.name, t.props, t.selfClosing, w)
	if err != nil {
		return err
	}

	if t.selfClosing {
		return nil
	}

	for _, c := range t.children {
		if c == nil {
			continue
		}

		_, err = w.Write([]byte("\n"))
		if err != nil {
			return err
		}

		err = c.RenderIndent(w, prefix+indent, indent)
		if err != nil {
			return err
		}

		_, err = w.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}

	_, err = w.Write([]byte(prefix))
	if err != nil {
		return err
	}

	return writeEndTag(t.name, w)
}

func writeTag[P Props](tag string, props P, selfClosing bool, w io.Writer) error {
	_, err := w.Write([]byte("<" + tag))
	if err != nil {
		return err
	}

	err = props.writeTo(w)
	if err != nil {
		return err
	}

	if selfClosing {
		_, err = w.Write([]byte(" />"))
	} else {
		_, err = w.Write([]byte(">"))
	}

	return err
}

func writeEndTag(tag string, w io.Writer) error {
	_, err := w.Write([]byte("</" + tag + ">"))
	return err
}

func writeAttr(name string, value any, w io.Writer) error {
	_, err := fmt.Fprintf(w, ` %s="%v"`, name, value)
	return err
}

func writeStyle(style map[string]string, w io.Writer) error {
	var b bytes.Buffer

	for k, v := range style {
		_, err := b.WriteString(k + ":" + v + ";")
		if err != nil {
			return err
		}
	}

	_, err := b.WriteTo(w)

	return err
}
