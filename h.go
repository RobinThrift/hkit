package hkit

type HCompFunc func(...Component) Component

type HComp interface {
	HCompFunc | string
}

func H[C HComp](comp C, children ...Component) Component {
	var c any = comp
	switch c := c.(type) {
	case string:
		return &tag[*GenericProps]{
			name:     c,
			children: children,
		}
	case HCompFunc:
		return c(children...)
	}

	return nil
}

type HpComp[P Props] interface {
	ComponentFunc[P] | string
}

func Hp[P Props, C HpComp[P]](comp C, props P, children ...Component) Component {
	var c any = comp
	switch c := c.(type) {
	case string:
		return &tag[P]{
			name:     c,
			props:    props,
			children: children,
		}
	case ComponentFunc[P]:
		return c(props, children...)
	}

	return nil
}

func WithProps[P Props, C interface {
	string | ~func(P, ...Component) Component
}](comp C, props P) HCompFunc {
	var c any = comp
	switch c := c.(type) {
	case string:
		return func(children ...Component) Component {
			return &tag[P]{
				name:     c,
				props:    props,
				children: children,
			}
		}
	case func(P, ...Component) Component:
		return func(children ...Component) Component {
			return c(props, children...)
		}
	}

	return func(...Component) Component {
		return nil
	}
}
