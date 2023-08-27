package hkit

import (
	"io"
)

type GenericProps struct {
	Accesskey       string
	Autocapitalize  string
	Autofocus       bool
	Contenteditable string
	Dir             string
	Draggable       bool
	Enterkeyhint    string
	Hidden          string
	Inputmode       string
	Is              string
	Itemid          string
	Itemprop        string
	Itemref         string
	Itemscope       bool
	Itemtype        string
	Lang            string
	Nonce           string
	Spellcheck      bool
	Tabindex        int
	Title           string
	Translate       bool
}

func Text(children ...string) Component {
	return textComponent{
		text: children,
	}
}

type textComponent struct {
	text []string
}

func (t textComponent) Render(w io.Writer) error {
	return t.RenderIndent(w, "", "")
}

func (t textComponent) RenderIndent(w io.Writer, prefix string, indent string) error {
	switch len(t.text) {
	case 0:
		return nil
	case 1:
		_, err := w.Write([]byte(prefix + t.text[0]))
		return err
	}

	_, err := w.Write([]byte(prefix + t.text[0]))
	if err != nil {
		return err
	}

	for _, s := range t.text[1:] {
		_, err = w.Write([]byte("\n" + prefix + s))
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *GenericProps) writeTo(w io.Writer) error { //nolint: gocyclo,gocognit,varnamelen // this is just a encoder, ignore
	if g == nil {
		return nil
	}

	if g.Accesskey != "" {
		err := writeAttr("accesskey", g.Accesskey, w)
		if err != nil {
			return err
		}
	}

	if g.Autocapitalize != "" {
		err := writeAttr("autocapitalize", g.Autocapitalize, w)
		if err != nil {
			return err
		}
	}

	if g.Autofocus {
		err := writeAttr("autofocus", g.Autofocus, w)
		if err != nil {
			return err
		}
	}

	if g.Contenteditable != "" {
		err := writeAttr("contenteditable", g.Contenteditable, w)
		if err != nil {
			return err
		}
	}

	if g.Dir != "" {
		err := writeAttr("dir", g.Dir, w)
		if err != nil {
			return err
		}
	}

	if g.Draggable {
		err := writeAttr("draggable", g.Draggable, w)
		if err != nil {
			return err
		}
	}

	if g.Enterkeyhint != "" {
		err := writeAttr("enterkeyhint", g.Enterkeyhint, w)
		if err != nil {
			return err
		}
	}

	if g.Hidden != "" {
		err := writeAttr("hidden", g.Hidden, w)
		if err != nil {
			return err
		}
	}

	if g.Inputmode != "" {
		err := writeAttr("inputmode", g.Inputmode, w)
		if err != nil {
			return err
		}
	}

	if g.Is != "" {
		err := writeAttr("is", g.Is, w)
		if err != nil {
			return err
		}
	}

	if g.Itemid != "" {
		err := writeAttr("itemid", g.Itemid, w)
		if err != nil {
			return err
		}
	}

	if g.Itemprop != "" {
		err := writeAttr("itemprop", g.Itemprop, w)
		if err != nil {
			return err
		}
	}

	if g.Itemref != "" {
		err := writeAttr("itemref", g.Itemref, w)
		if err != nil {
			return err
		}
	}

	if g.Itemscope {
		err := writeAttr("itemscope", g.Itemscope, w)
		if err != nil {
			return err
		}
	}

	if g.Itemtype != "" {
		err := writeAttr("itemtype", g.Itemtype, w)
		if err != nil {
			return err
		}
	}

	if g.Lang != "" {
		err := writeAttr("lang", g.Lang, w)
		if err != nil {
			return err
		}
	}

	if g.Nonce != "" {
		err := writeAttr("nonce", g.Nonce, w)
		if err != nil {
			return err
		}
	}

	if g.Spellcheck {
		err := writeAttr("spellcheck", g.Spellcheck, w)
		if err != nil {
			return err
		}
	}

	if g.Tabindex != 0 {
		err := writeAttr("tabindex", g.Tabindex, w)
		if err != nil {
			return err
		}
	}

	if g.Title != "" {
		err := writeAttr("title", g.Title, w)
		if err != nil {
			return err
		}
	}

	if g.Translate {
		err := writeAttr("translate", "yes", w)
		if err != nil {
			return err
		}
	}

	return nil
}
