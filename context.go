package hkit

// Context .
type Context interface {
	Get(k string) interface{}
	Set(k string, v interface{})
}

// NewContext .
func NewContext() Context {
	return context{}
}

type context map[string]interface{}

func (c context) Get(k string) interface{} {
	return c[k]
}

func (c context) Set(k string, v interface{}) {
	c[k] = v
}
