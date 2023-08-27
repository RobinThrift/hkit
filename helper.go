package hkit

import (
	"bytes"
)

type Classes map[string]bool

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
