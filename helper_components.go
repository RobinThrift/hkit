package hkit

// Text .
func Text(t string) Component {
	return &comp{
		tag:   "_text_",
		attrs: t,
	}
}

// If .
func If(cond bool, a, b Component) Component {
	if cond {
		return a
	}

	return b
}
