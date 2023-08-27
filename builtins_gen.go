package hkit

import "io"

type AProps struct {
	*GenericProps
	X              map[string]string
	ID             string
	Class          string
	Style          map[string]string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

func (p *AProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Href != "" {
		err := writeAttr("href", p.Href, w)
		if err != nil {
			return err
		}
	}
	if p.Target != "" {
		err := writeAttr("target", p.Target, w)
		if err != nil {
			return err
		}
	}
	if p.Download != "" {
		err := writeAttr("download", p.Download, w)
		if err != nil {
			return err
		}
	}
	if p.Ping != "" {
		err := writeAttr("ping", p.Ping, w)
		if err != nil {
			return err
		}
	}
	if p.Rel != "" {
		err := writeAttr("rel", p.Rel, w)
		if err != nil {
			return err
		}
	}
	if p.Hreflang != "" {
		err := writeAttr("hreflang", p.Hreflang, w)
		if err != nil {
			return err
		}
	}
	if p.Type != "" {
		err := writeAttr("type", p.Type, w)
		if err != nil {
			return err
		}
	}
	if p.Referrerpolicy != "" {
		err := writeAttr("referrerpolicy", p.Referrerpolicy, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func A(props *AProps, children ...Component) Component {
	return &tag[*AProps]{
		children:    children,
		name:        "a",
		props:       props,
		selfClosing: false,
	}
}

type AbbrProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *AbbrProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Abbr(props *AbbrProps, children ...Component) Component {
	return &tag[*AbbrProps]{
		children:    children,
		name:        "abbr",
		props:       props,
		selfClosing: false,
	}
}

type AddressProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *AddressProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Address(props *AddressProps, children ...Component) Component {
	return &tag[*AddressProps]{
		children:    children,
		name:        "address",
		props:       props,
		selfClosing: false,
	}
}

type AreaProps struct {
	*GenericProps
	X              map[string]string
	ID             string
	Class          string
	Style          map[string]string
	Alt            string
	Coords         string
	Shape          string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Referrerpolicy string
}

func (p *AreaProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Alt != "" {
		err := writeAttr("alt", p.Alt, w)
		if err != nil {
			return err
		}
	}
	if p.Coords != "" {
		err := writeAttr("coords", p.Coords, w)
		if err != nil {
			return err
		}
	}
	if p.Shape != "" {
		err := writeAttr("shape", p.Shape, w)
		if err != nil {
			return err
		}
	}
	if p.Href != "" {
		err := writeAttr("href", p.Href, w)
		if err != nil {
			return err
		}
	}
	if p.Target != "" {
		err := writeAttr("target", p.Target, w)
		if err != nil {
			return err
		}
	}
	if p.Download != "" {
		err := writeAttr("download", p.Download, w)
		if err != nil {
			return err
		}
	}
	if p.Ping != "" {
		err := writeAttr("ping", p.Ping, w)
		if err != nil {
			return err
		}
	}
	if p.Rel != "" {
		err := writeAttr("rel", p.Rel, w)
		if err != nil {
			return err
		}
	}
	if p.Referrerpolicy != "" {
		err := writeAttr("referrerpolicy", p.Referrerpolicy, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Area(props *AreaProps, children ...Component) Component {
	return &tag[*AreaProps]{
		children:    children,
		name:        "area",
		props:       props,
		selfClosing: true,
	}
}

type ArticleProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *ArticleProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Article(props *ArticleProps, children ...Component) Component {
	return &tag[*ArticleProps]{
		children:    children,
		name:        "article",
		props:       props,
		selfClosing: false,
	}
}

type AsideProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *AsideProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Aside(props *AsideProps, children ...Component) Component {
	return &tag[*AsideProps]{
		children:    children,
		name:        "aside",
		props:       props,
		selfClosing: false,
	}
}

type AudioProps struct {
	*GenericProps
	X           map[string]string
	ID          string
	Class       string
	Style       map[string]string
	Src         string
	Crossorigin string
	Preload     string
	Autoplay    string
	Loop        string
	Muted       string
	Controls    string
}

func (p *AudioProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Src != "" {
		err := writeAttr("src", p.Src, w)
		if err != nil {
			return err
		}
	}
	if p.Crossorigin != "" {
		err := writeAttr("crossorigin", p.Crossorigin, w)
		if err != nil {
			return err
		}
	}
	if p.Preload != "" {
		err := writeAttr("preload", p.Preload, w)
		if err != nil {
			return err
		}
	}
	if p.Autoplay != "" {
		err := writeAttr("autoplay", p.Autoplay, w)
		if err != nil {
			return err
		}
	}
	if p.Loop != "" {
		err := writeAttr("loop", p.Loop, w)
		if err != nil {
			return err
		}
	}
	if p.Muted != "" {
		err := writeAttr("muted", p.Muted, w)
		if err != nil {
			return err
		}
	}
	if p.Controls != "" {
		err := writeAttr("controls", p.Controls, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Audio(props *AudioProps, children ...Component) Component {
	return &tag[*AudioProps]{
		children:    children,
		name:        "audio",
		props:       props,
		selfClosing: false,
	}
}

type BProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *BProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func B(props *BProps, children ...Component) Component {
	return &tag[*BProps]{
		children:    children,
		name:        "b",
		props:       props,
		selfClosing: false,
	}
}

type BaseProps struct {
	*GenericProps
	X      map[string]string
	ID     string
	Class  string
	Style  map[string]string
	Href   string
	Target string
}

func (p *BaseProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Href != "" {
		err := writeAttr("href", p.Href, w)
		if err != nil {
			return err
		}
	}
	if p.Target != "" {
		err := writeAttr("target", p.Target, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Base(props *BaseProps, children ...Component) Component {
	return &tag[*BaseProps]{
		children:    children,
		name:        "base",
		props:       props,
		selfClosing: true,
	}
}

type BdiProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *BdiProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Bdi(props *BdiProps, children ...Component) Component {
	return &tag[*BdiProps]{
		children:    children,
		name:        "bdi",
		props:       props,
		selfClosing: false,
	}
}

type BdoProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *BdoProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Bdo(props *BdoProps, children ...Component) Component {
	return &tag[*BdoProps]{
		children:    children,
		name:        "bdo",
		props:       props,
		selfClosing: false,
	}
}

type BlockquoteProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
	Cite  string
}

func (p *BlockquoteProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Cite != "" {
		err := writeAttr("cite", p.Cite, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Blockquote(props *BlockquoteProps, children ...Component) Component {
	return &tag[*BlockquoteProps]{
		children:    children,
		name:        "blockquote",
		props:       props,
		selfClosing: false,
	}
}

type BodyProps struct {
	*GenericProps
	X                    map[string]string
	ID                   string
	Class                string
	Style                map[string]string
	Onafterprint         string
	Onbeforeprint        string
	Onbeforeunload       string
	Onhashchange         string
	Onlanguagechange     string
	Onmessage            string
	Onmessageerror       string
	Onoffline            string
	Ononline             string
	Onpagehide           string
	Onpageshow           string
	Onpopstate           string
	Onrejectionhandled   string
	Onstorage            string
	Onunhandledrejection string
	Onunload             string
}

func (p *BodyProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Onafterprint != "" {
		err := writeAttr("onafterprint", p.Onafterprint, w)
		if err != nil {
			return err
		}
	}
	if p.Onbeforeprint != "" {
		err := writeAttr("onbeforeprint", p.Onbeforeprint, w)
		if err != nil {
			return err
		}
	}
	if p.Onbeforeunload != "" {
		err := writeAttr("onbeforeunload", p.Onbeforeunload, w)
		if err != nil {
			return err
		}
	}
	if p.Onhashchange != "" {
		err := writeAttr("onhashchange", p.Onhashchange, w)
		if err != nil {
			return err
		}
	}
	if p.Onlanguagechange != "" {
		err := writeAttr("onlanguagechange", p.Onlanguagechange, w)
		if err != nil {
			return err
		}
	}
	if p.Onmessage != "" {
		err := writeAttr("onmessage", p.Onmessage, w)
		if err != nil {
			return err
		}
	}
	if p.Onmessageerror != "" {
		err := writeAttr("onmessageerror", p.Onmessageerror, w)
		if err != nil {
			return err
		}
	}
	if p.Onoffline != "" {
		err := writeAttr("onoffline", p.Onoffline, w)
		if err != nil {
			return err
		}
	}
	if p.Ononline != "" {
		err := writeAttr("ononline", p.Ononline, w)
		if err != nil {
			return err
		}
	}
	if p.Onpagehide != "" {
		err := writeAttr("onpagehide", p.Onpagehide, w)
		if err != nil {
			return err
		}
	}
	if p.Onpageshow != "" {
		err := writeAttr("onpageshow", p.Onpageshow, w)
		if err != nil {
			return err
		}
	}
	if p.Onpopstate != "" {
		err := writeAttr("onpopstate", p.Onpopstate, w)
		if err != nil {
			return err
		}
	}
	if p.Onrejectionhandled != "" {
		err := writeAttr("onrejectionhandled", p.Onrejectionhandled, w)
		if err != nil {
			return err
		}
	}
	if p.Onstorage != "" {
		err := writeAttr("onstorage", p.Onstorage, w)
		if err != nil {
			return err
		}
	}
	if p.Onunhandledrejection != "" {
		err := writeAttr("onunhandledrejection", p.Onunhandledrejection, w)
		if err != nil {
			return err
		}
	}
	if p.Onunload != "" {
		err := writeAttr("onunload", p.Onunload, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Body(props *BodyProps, children ...Component) Component {
	return &tag[*BodyProps]{
		children:    children,
		name:        "body",
		props:       props,
		selfClosing: false,
	}
}

type BrProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *BrProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Br(props *BrProps, children ...Component) Component {
	return &tag[*BrProps]{
		children:    children,
		name:        "br",
		props:       props,
		selfClosing: true,
	}
}

type ButtonProps struct {
	*GenericProps
	X                   map[string]string
	ID                  string
	Class               string
	Style               map[string]string
	Disabled            string
	Form                string
	Formaction          string
	Formenctype         string
	Formmethod          string
	Formnovalidate      string
	Formtarget          string
	Name                string
	Popovertarget       string
	Popovertargetaction string
	Type                string
	Value               string
}

func (p *ButtonProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Disabled != "" {
		err := writeAttr("disabled", p.Disabled, w)
		if err != nil {
			return err
		}
	}
	if p.Form != "" {
		err := writeAttr("form", p.Form, w)
		if err != nil {
			return err
		}
	}
	if p.Formaction != "" {
		err := writeAttr("formaction", p.Formaction, w)
		if err != nil {
			return err
		}
	}
	if p.Formenctype != "" {
		err := writeAttr("formenctype", p.Formenctype, w)
		if err != nil {
			return err
		}
	}
	if p.Formmethod != "" {
		err := writeAttr("formmethod", p.Formmethod, w)
		if err != nil {
			return err
		}
	}
	if p.Formnovalidate != "" {
		err := writeAttr("formnovalidate", p.Formnovalidate, w)
		if err != nil {
			return err
		}
	}
	if p.Formtarget != "" {
		err := writeAttr("formtarget", p.Formtarget, w)
		if err != nil {
			return err
		}
	}
	if p.Name != "" {
		err := writeAttr("name", p.Name, w)
		if err != nil {
			return err
		}
	}
	if p.Popovertarget != "" {
		err := writeAttr("popovertarget", p.Popovertarget, w)
		if err != nil {
			return err
		}
	}
	if p.Popovertargetaction != "" {
		err := writeAttr("popovertargetaction", p.Popovertargetaction, w)
		if err != nil {
			return err
		}
	}
	if p.Type != "" {
		err := writeAttr("type", p.Type, w)
		if err != nil {
			return err
		}
	}
	if p.Value != "" {
		err := writeAttr("value", p.Value, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Button(props *ButtonProps, children ...Component) Component {
	return &tag[*ButtonProps]{
		children:    children,
		name:        "button",
		props:       props,
		selfClosing: false,
	}
}

type CanvasProps struct {
	*GenericProps
	X      map[string]string
	ID     string
	Class  string
	Style  map[string]string
	Width  string
	Height string
}

func (p *CanvasProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Width != "" {
		err := writeAttr("width", p.Width, w)
		if err != nil {
			return err
		}
	}
	if p.Height != "" {
		err := writeAttr("height", p.Height, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Canvas(props *CanvasProps, children ...Component) Component {
	return &tag[*CanvasProps]{
		children:    children,
		name:        "canvas",
		props:       props,
		selfClosing: false,
	}
}

type CaptionProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *CaptionProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Caption(props *CaptionProps, children ...Component) Component {
	return &tag[*CaptionProps]{
		children:    children,
		name:        "caption",
		props:       props,
		selfClosing: false,
	}
}

type CiteProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *CiteProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Cite(props *CiteProps, children ...Component) Component {
	return &tag[*CiteProps]{
		children:    children,
		name:        "cite",
		props:       props,
		selfClosing: false,
	}
}

type CodeProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *CodeProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Code(props *CodeProps, children ...Component) Component {
	return &tag[*CodeProps]{
		children:    children,
		name:        "code",
		props:       props,
		selfClosing: false,
	}
}

type ColProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
	Span  string
}

func (p *ColProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Span != "" {
		err := writeAttr("span", p.Span, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Col(props *ColProps, children ...Component) Component {
	return &tag[*ColProps]{
		children:    children,
		name:        "col",
		props:       props,
		selfClosing: true,
	}
}

type ColgroupProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
	Span  string
}

func (p *ColgroupProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Span != "" {
		err := writeAttr("span", p.Span, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Colgroup(props *ColgroupProps, children ...Component) Component {
	return &tag[*ColgroupProps]{
		children:    children,
		name:        "colgroup",
		props:       props,
		selfClosing: false,
	}
}

type DataProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
	Value string
}

func (p *DataProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Value != "" {
		err := writeAttr("value", p.Value, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Data(props *DataProps, children ...Component) Component {
	return &tag[*DataProps]{
		children:    children,
		name:        "data",
		props:       props,
		selfClosing: false,
	}
}

type DatalistProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *DatalistProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Datalist(props *DatalistProps, children ...Component) Component {
	return &tag[*DatalistProps]{
		children:    children,
		name:        "datalist",
		props:       props,
		selfClosing: false,
	}
}

type DdProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *DdProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Dd(props *DdProps, children ...Component) Component {
	return &tag[*DdProps]{
		children:    children,
		name:        "dd",
		props:       props,
		selfClosing: false,
	}
}

type DelProps struct {
	*GenericProps
	X        map[string]string
	ID       string
	Class    string
	Style    map[string]string
	Cite     string
	Datetime string
}

func (p *DelProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Cite != "" {
		err := writeAttr("cite", p.Cite, w)
		if err != nil {
			return err
		}
	}
	if p.Datetime != "" {
		err := writeAttr("datetime", p.Datetime, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Del(props *DelProps, children ...Component) Component {
	return &tag[*DelProps]{
		children:    children,
		name:        "del",
		props:       props,
		selfClosing: false,
	}
}

type DetailsProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
	Open  string
}

func (p *DetailsProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Open != "" {
		err := writeAttr("open", p.Open, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Details(props *DetailsProps, children ...Component) Component {
	return &tag[*DetailsProps]{
		children:    children,
		name:        "details",
		props:       props,
		selfClosing: false,
	}
}

type DfnProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *DfnProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Dfn(props *DfnProps, children ...Component) Component {
	return &tag[*DfnProps]{
		children:    children,
		name:        "dfn",
		props:       props,
		selfClosing: false,
	}
}

type DialogProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
	Open  string
}

func (p *DialogProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Open != "" {
		err := writeAttr("open", p.Open, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Dialog(props *DialogProps, children ...Component) Component {
	return &tag[*DialogProps]{
		children:    children,
		name:        "dialog",
		props:       props,
		selfClosing: false,
	}
}

type DivProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *DivProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Div(props *DivProps, children ...Component) Component {
	return &tag[*DivProps]{
		children:    children,
		name:        "div",
		props:       props,
		selfClosing: false,
	}
}

type DlProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *DlProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Dl(props *DlProps, children ...Component) Component {
	return &tag[*DlProps]{
		children:    children,
		name:        "dl",
		props:       props,
		selfClosing: false,
	}
}

type DtProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *DtProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Dt(props *DtProps, children ...Component) Component {
	return &tag[*DtProps]{
		children:    children,
		name:        "dt",
		props:       props,
		selfClosing: false,
	}
}

type EmProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *EmProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Em(props *EmProps, children ...Component) Component {
	return &tag[*EmProps]{
		children:    children,
		name:        "em",
		props:       props,
		selfClosing: false,
	}
}

type EmbedProps struct {
	*GenericProps
	X      map[string]string
	ID     string
	Class  string
	Style  map[string]string
	Src    string
	Type   string
	Width  string
	Height string
	Any    string
}

func (p *EmbedProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Src != "" {
		err := writeAttr("src", p.Src, w)
		if err != nil {
			return err
		}
	}
	if p.Type != "" {
		err := writeAttr("type", p.Type, w)
		if err != nil {
			return err
		}
	}
	if p.Width != "" {
		err := writeAttr("width", p.Width, w)
		if err != nil {
			return err
		}
	}
	if p.Height != "" {
		err := writeAttr("height", p.Height, w)
		if err != nil {
			return err
		}
	}
	if p.Any != "" {
		err := writeAttr("any", p.Any, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Embed(props *EmbedProps, children ...Component) Component {
	return &tag[*EmbedProps]{
		children:    children,
		name:        "embed",
		props:       props,
		selfClosing: true,
	}
}

type FieldsetProps struct {
	*GenericProps
	X        map[string]string
	ID       string
	Class    string
	Style    map[string]string
	Disabled string
	Form     string
	Name     string
}

func (p *FieldsetProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Disabled != "" {
		err := writeAttr("disabled", p.Disabled, w)
		if err != nil {
			return err
		}
	}
	if p.Form != "" {
		err := writeAttr("form", p.Form, w)
		if err != nil {
			return err
		}
	}
	if p.Name != "" {
		err := writeAttr("name", p.Name, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Fieldset(props *FieldsetProps, children ...Component) Component {
	return &tag[*FieldsetProps]{
		children:    children,
		name:        "fieldset",
		props:       props,
		selfClosing: false,
	}
}

type FigcaptionProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *FigcaptionProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Figcaption(props *FigcaptionProps, children ...Component) Component {
	return &tag[*FigcaptionProps]{
		children:    children,
		name:        "figcaption",
		props:       props,
		selfClosing: false,
	}
}

type FigureProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *FigureProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Figure(props *FigureProps, children ...Component) Component {
	return &tag[*FigureProps]{
		children:    children,
		name:        "figure",
		props:       props,
		selfClosing: false,
	}
}

type FooterProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *FooterProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Footer(props *FooterProps, children ...Component) Component {
	return &tag[*FooterProps]{
		children:    children,
		name:        "footer",
		props:       props,
		selfClosing: false,
	}
}

type FormProps struct {
	*GenericProps
	X             map[string]string
	ID            string
	Class         string
	Style         map[string]string
	AcceptCharset string
	Action        string
	Autocomplete  string
	Enctype       string
	Method        string
	Name          string
	Novalidate    string
	Rel           string
	Target        string
}

func (p *FormProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.AcceptCharset != "" {
		err := writeAttr("acceptcharset", p.AcceptCharset, w)
		if err != nil {
			return err
		}
	}
	if p.Action != "" {
		err := writeAttr("action", p.Action, w)
		if err != nil {
			return err
		}
	}
	if p.Autocomplete != "" {
		err := writeAttr("autocomplete", p.Autocomplete, w)
		if err != nil {
			return err
		}
	}
	if p.Enctype != "" {
		err := writeAttr("enctype", p.Enctype, w)
		if err != nil {
			return err
		}
	}
	if p.Method != "" {
		err := writeAttr("method", p.Method, w)
		if err != nil {
			return err
		}
	}
	if p.Name != "" {
		err := writeAttr("name", p.Name, w)
		if err != nil {
			return err
		}
	}
	if p.Novalidate != "" {
		err := writeAttr("novalidate", p.Novalidate, w)
		if err != nil {
			return err
		}
	}
	if p.Rel != "" {
		err := writeAttr("rel", p.Rel, w)
		if err != nil {
			return err
		}
	}
	if p.Target != "" {
		err := writeAttr("target", p.Target, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Form(props *FormProps, children ...Component) Component {
	return &tag[*FormProps]{
		children:    children,
		name:        "form",
		props:       props,
		selfClosing: false,
	}
}

type H1Props struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *H1Props) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func H1(props *H1Props, children ...Component) Component {
	return &tag[*H1Props]{
		children:    children,
		name:        "h1",
		props:       props,
		selfClosing: false,
	}
}

type H2Props struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *H2Props) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func H2(props *H2Props, children ...Component) Component {
	return &tag[*H2Props]{
		children:    children,
		name:        "h2",
		props:       props,
		selfClosing: false,
	}
}

type H3Props struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *H3Props) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func H3(props *H3Props, children ...Component) Component {
	return &tag[*H3Props]{
		children:    children,
		name:        "h3",
		props:       props,
		selfClosing: false,
	}
}

type H4Props struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *H4Props) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func H4(props *H4Props, children ...Component) Component {
	return &tag[*H4Props]{
		children:    children,
		name:        "h4",
		props:       props,
		selfClosing: false,
	}
}

type H5Props struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *H5Props) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func H5(props *H5Props, children ...Component) Component {
	return &tag[*H5Props]{
		children:    children,
		name:        "h5",
		props:       props,
		selfClosing: false,
	}
}

type H6Props struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *H6Props) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func H6(props *H6Props, children ...Component) Component {
	return &tag[*H6Props]{
		children:    children,
		name:        "h6",
		props:       props,
		selfClosing: false,
	}
}

type HeadProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *HeadProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Head(props *HeadProps, children ...Component) Component {
	return &tag[*HeadProps]{
		children:    children,
		name:        "head",
		props:       props,
		selfClosing: false,
	}
}

type HeaderProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *HeaderProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Header(props *HeaderProps, children ...Component) Component {
	return &tag[*HeaderProps]{
		children:    children,
		name:        "header",
		props:       props,
		selfClosing: false,
	}
}

type HgroupProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *HgroupProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Hgroup(props *HgroupProps, children ...Component) Component {
	return &tag[*HgroupProps]{
		children:    children,
		name:        "hgroup",
		props:       props,
		selfClosing: false,
	}
}

type HrProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *HrProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Hr(props *HrProps, children ...Component) Component {
	return &tag[*HrProps]{
		children:    children,
		name:        "hr",
		props:       props,
		selfClosing: true,
	}
}

type HtmlProps struct {
	*GenericProps
	X        map[string]string
	ID       string
	Class    string
	Style    map[string]string
	Manifest string
}

func (p *HtmlProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Manifest != "" {
		err := writeAttr("manifest", p.Manifest, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Html(props *HtmlProps, children ...Component) Component {
	return &tag[*HtmlProps]{
		children:    children,
		name:        "html",
		props:       props,
		selfClosing: false,
	}
}

type IProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *IProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func I(props *IProps, children ...Component) Component {
	return &tag[*IProps]{
		children:    children,
		name:        "i",
		props:       props,
		selfClosing: false,
	}
}

type IframeProps struct {
	*GenericProps
	X               map[string]string
	ID              string
	Class           string
	Style           map[string]string
	Src             string
	Srcdoc          string
	Name            string
	Sandbox         string
	Allow           string
	Allowfullscreen string
	Width           string
	Height          string
	Referrerpolicy  string
	Loading         string
}

func (p *IframeProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Src != "" {
		err := writeAttr("src", p.Src, w)
		if err != nil {
			return err
		}
	}
	if p.Srcdoc != "" {
		err := writeAttr("srcdoc", p.Srcdoc, w)
		if err != nil {
			return err
		}
	}
	if p.Name != "" {
		err := writeAttr("name", p.Name, w)
		if err != nil {
			return err
		}
	}
	if p.Sandbox != "" {
		err := writeAttr("sandbox", p.Sandbox, w)
		if err != nil {
			return err
		}
	}
	if p.Allow != "" {
		err := writeAttr("allow", p.Allow, w)
		if err != nil {
			return err
		}
	}
	if p.Allowfullscreen != "" {
		err := writeAttr("allowfullscreen", p.Allowfullscreen, w)
		if err != nil {
			return err
		}
	}
	if p.Width != "" {
		err := writeAttr("width", p.Width, w)
		if err != nil {
			return err
		}
	}
	if p.Height != "" {
		err := writeAttr("height", p.Height, w)
		if err != nil {
			return err
		}
	}
	if p.Referrerpolicy != "" {
		err := writeAttr("referrerpolicy", p.Referrerpolicy, w)
		if err != nil {
			return err
		}
	}
	if p.Loading != "" {
		err := writeAttr("loading", p.Loading, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Iframe(props *IframeProps, children ...Component) Component {
	return &tag[*IframeProps]{
		children:    children,
		name:        "iframe",
		props:       props,
		selfClosing: true,
	}
}

type ImgProps struct {
	*GenericProps
	X              map[string]string
	ID             string
	Class          string
	Style          map[string]string
	Alt            string
	Src            string
	Srcset         string
	Sizes          string
	Crossorigin    string
	Usemap         string
	Ismap          string
	Width          string
	Height         string
	Referrerpolicy string
	Decoding       string
	Loading        string
	Fetchpriority  string
}

func (p *ImgProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Alt != "" {
		err := writeAttr("alt", p.Alt, w)
		if err != nil {
			return err
		}
	}
	if p.Src != "" {
		err := writeAttr("src", p.Src, w)
		if err != nil {
			return err
		}
	}
	if p.Srcset != "" {
		err := writeAttr("srcset", p.Srcset, w)
		if err != nil {
			return err
		}
	}
	if p.Sizes != "" {
		err := writeAttr("sizes", p.Sizes, w)
		if err != nil {
			return err
		}
	}
	if p.Crossorigin != "" {
		err := writeAttr("crossorigin", p.Crossorigin, w)
		if err != nil {
			return err
		}
	}
	if p.Usemap != "" {
		err := writeAttr("usemap", p.Usemap, w)
		if err != nil {
			return err
		}
	}
	if p.Ismap != "" {
		err := writeAttr("ismap", p.Ismap, w)
		if err != nil {
			return err
		}
	}
	if p.Width != "" {
		err := writeAttr("width", p.Width, w)
		if err != nil {
			return err
		}
	}
	if p.Height != "" {
		err := writeAttr("height", p.Height, w)
		if err != nil {
			return err
		}
	}
	if p.Referrerpolicy != "" {
		err := writeAttr("referrerpolicy", p.Referrerpolicy, w)
		if err != nil {
			return err
		}
	}
	if p.Decoding != "" {
		err := writeAttr("decoding", p.Decoding, w)
		if err != nil {
			return err
		}
	}
	if p.Loading != "" {
		err := writeAttr("loading", p.Loading, w)
		if err != nil {
			return err
		}
	}
	if p.Fetchpriority != "" {
		err := writeAttr("fetchpriority", p.Fetchpriority, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Img(props *ImgProps, children ...Component) Component {
	return &tag[*ImgProps]{
		children:    children,
		name:        "img",
		props:       props,
		selfClosing: true,
	}
}

type InputProps struct {
	*GenericProps
	X                   map[string]string
	ID                  string
	Class               string
	Style               map[string]string
	Accept              string
	Alt                 string
	Autocomplete        string
	Checked             string
	Dirname             string
	Disabled            string
	Form                string
	Formaction          string
	Formenctype         string
	Formmethod          string
	Formnovalidate      string
	Formtarget          string
	Height              string
	List                string
	Max                 string
	Maxlength           string
	Min                 string
	Minlength           string
	Multiple            string
	Name                string
	Pattern             string
	Placeholder         string
	Popovertarget       string
	Popovertargetaction string
	Readonly            string
	Required            string
	Size                string
	Src                 string
	Step                string
	Type                string
	Value               string
	Width               string
}

func (p *InputProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Accept != "" {
		err := writeAttr("accept", p.Accept, w)
		if err != nil {
			return err
		}
	}
	if p.Alt != "" {
		err := writeAttr("alt", p.Alt, w)
		if err != nil {
			return err
		}
	}
	if p.Autocomplete != "" {
		err := writeAttr("autocomplete", p.Autocomplete, w)
		if err != nil {
			return err
		}
	}
	if p.Checked != "" {
		err := writeAttr("checked", p.Checked, w)
		if err != nil {
			return err
		}
	}
	if p.Dirname != "" {
		err := writeAttr("dirname", p.Dirname, w)
		if err != nil {
			return err
		}
	}
	if p.Disabled != "" {
		err := writeAttr("disabled", p.Disabled, w)
		if err != nil {
			return err
		}
	}
	if p.Form != "" {
		err := writeAttr("form", p.Form, w)
		if err != nil {
			return err
		}
	}
	if p.Formaction != "" {
		err := writeAttr("formaction", p.Formaction, w)
		if err != nil {
			return err
		}
	}
	if p.Formenctype != "" {
		err := writeAttr("formenctype", p.Formenctype, w)
		if err != nil {
			return err
		}
	}
	if p.Formmethod != "" {
		err := writeAttr("formmethod", p.Formmethod, w)
		if err != nil {
			return err
		}
	}
	if p.Formnovalidate != "" {
		err := writeAttr("formnovalidate", p.Formnovalidate, w)
		if err != nil {
			return err
		}
	}
	if p.Formtarget != "" {
		err := writeAttr("formtarget", p.Formtarget, w)
		if err != nil {
			return err
		}
	}
	if p.Height != "" {
		err := writeAttr("height", p.Height, w)
		if err != nil {
			return err
		}
	}
	if p.List != "" {
		err := writeAttr("list", p.List, w)
		if err != nil {
			return err
		}
	}
	if p.Max != "" {
		err := writeAttr("max", p.Max, w)
		if err != nil {
			return err
		}
	}
	if p.Maxlength != "" {
		err := writeAttr("maxlength", p.Maxlength, w)
		if err != nil {
			return err
		}
	}
	if p.Min != "" {
		err := writeAttr("min", p.Min, w)
		if err != nil {
			return err
		}
	}
	if p.Minlength != "" {
		err := writeAttr("minlength", p.Minlength, w)
		if err != nil {
			return err
		}
	}
	if p.Multiple != "" {
		err := writeAttr("multiple", p.Multiple, w)
		if err != nil {
			return err
		}
	}
	if p.Name != "" {
		err := writeAttr("name", p.Name, w)
		if err != nil {
			return err
		}
	}
	if p.Pattern != "" {
		err := writeAttr("pattern", p.Pattern, w)
		if err != nil {
			return err
		}
	}
	if p.Placeholder != "" {
		err := writeAttr("placeholder", p.Placeholder, w)
		if err != nil {
			return err
		}
	}
	if p.Popovertarget != "" {
		err := writeAttr("popovertarget", p.Popovertarget, w)
		if err != nil {
			return err
		}
	}
	if p.Popovertargetaction != "" {
		err := writeAttr("popovertargetaction", p.Popovertargetaction, w)
		if err != nil {
			return err
		}
	}
	if p.Readonly != "" {
		err := writeAttr("readonly", p.Readonly, w)
		if err != nil {
			return err
		}
	}
	if p.Required != "" {
		err := writeAttr("required", p.Required, w)
		if err != nil {
			return err
		}
	}
	if p.Size != "" {
		err := writeAttr("size", p.Size, w)
		if err != nil {
			return err
		}
	}
	if p.Src != "" {
		err := writeAttr("src", p.Src, w)
		if err != nil {
			return err
		}
	}
	if p.Step != "" {
		err := writeAttr("step", p.Step, w)
		if err != nil {
			return err
		}
	}
	if p.Type != "" {
		err := writeAttr("type", p.Type, w)
		if err != nil {
			return err
		}
	}
	if p.Value != "" {
		err := writeAttr("value", p.Value, w)
		if err != nil {
			return err
		}
	}
	if p.Width != "" {
		err := writeAttr("width", p.Width, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Input(props *InputProps, children ...Component) Component {
	return &tag[*InputProps]{
		children:    children,
		name:        "input",
		props:       props,
		selfClosing: true,
	}
}

type InsProps struct {
	*GenericProps
	X        map[string]string
	ID       string
	Class    string
	Style    map[string]string
	Cite     string
	Datetime string
}

func (p *InsProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Cite != "" {
		err := writeAttr("cite", p.Cite, w)
		if err != nil {
			return err
		}
	}
	if p.Datetime != "" {
		err := writeAttr("datetime", p.Datetime, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Ins(props *InsProps, children ...Component) Component {
	return &tag[*InsProps]{
		children:    children,
		name:        "ins",
		props:       props,
		selfClosing: false,
	}
}

type KbdProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *KbdProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Kbd(props *KbdProps, children ...Component) Component {
	return &tag[*KbdProps]{
		children:    children,
		name:        "kbd",
		props:       props,
		selfClosing: false,
	}
}

type LabelProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
	For   string
}

func (p *LabelProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.For != "" {
		err := writeAttr("for", p.For, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Label(props *LabelProps, children ...Component) Component {
	return &tag[*LabelProps]{
		children:    children,
		name:        "label",
		props:       props,
		selfClosing: false,
	}
}

type LegendProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *LegendProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Legend(props *LegendProps, children ...Component) Component {
	return &tag[*LegendProps]{
		children:    children,
		name:        "legend",
		props:       props,
		selfClosing: false,
	}
}

type LiProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
	Value string
}

func (p *LiProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Value != "" {
		err := writeAttr("value", p.Value, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Li(props *LiProps, children ...Component) Component {
	return &tag[*LiProps]{
		children:    children,
		name:        "li",
		props:       props,
		selfClosing: false,
	}
}

type LinkProps struct {
	*GenericProps
	X              map[string]string
	ID             string
	Class          string
	Style          map[string]string
	Href           string
	Crossorigin    string
	Rel            string
	As             string
	Media          string
	Hreflang       string
	Type           string
	Sizes          string
	Imagesrcset    string
	Imagesizes     string
	Referrerpolicy string
	Integrity      string
	Blocking       string
	Color          string
	Disabled       string
	Fetchpriority  string
}

func (p *LinkProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Href != "" {
		err := writeAttr("href", p.Href, w)
		if err != nil {
			return err
		}
	}
	if p.Crossorigin != "" {
		err := writeAttr("crossorigin", p.Crossorigin, w)
		if err != nil {
			return err
		}
	}
	if p.Rel != "" {
		err := writeAttr("rel", p.Rel, w)
		if err != nil {
			return err
		}
	}
	if p.As != "" {
		err := writeAttr("as", p.As, w)
		if err != nil {
			return err
		}
	}
	if p.Media != "" {
		err := writeAttr("media", p.Media, w)
		if err != nil {
			return err
		}
	}
	if p.Hreflang != "" {
		err := writeAttr("hreflang", p.Hreflang, w)
		if err != nil {
			return err
		}
	}
	if p.Type != "" {
		err := writeAttr("type", p.Type, w)
		if err != nil {
			return err
		}
	}
	if p.Sizes != "" {
		err := writeAttr("sizes", p.Sizes, w)
		if err != nil {
			return err
		}
	}
	if p.Imagesrcset != "" {
		err := writeAttr("imagesrcset", p.Imagesrcset, w)
		if err != nil {
			return err
		}
	}
	if p.Imagesizes != "" {
		err := writeAttr("imagesizes", p.Imagesizes, w)
		if err != nil {
			return err
		}
	}
	if p.Referrerpolicy != "" {
		err := writeAttr("referrerpolicy", p.Referrerpolicy, w)
		if err != nil {
			return err
		}
	}
	if p.Integrity != "" {
		err := writeAttr("integrity", p.Integrity, w)
		if err != nil {
			return err
		}
	}
	if p.Blocking != "" {
		err := writeAttr("blocking", p.Blocking, w)
		if err != nil {
			return err
		}
	}
	if p.Color != "" {
		err := writeAttr("color", p.Color, w)
		if err != nil {
			return err
		}
	}
	if p.Disabled != "" {
		err := writeAttr("disabled", p.Disabled, w)
		if err != nil {
			return err
		}
	}
	if p.Fetchpriority != "" {
		err := writeAttr("fetchpriority", p.Fetchpriority, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Link(props *LinkProps, children ...Component) Component {
	return &tag[*LinkProps]{
		children:    children,
		name:        "link",
		props:       props,
		selfClosing: true,
	}
}

type MainProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *MainProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Main(props *MainProps, children ...Component) Component {
	return &tag[*MainProps]{
		children:    children,
		name:        "main",
		props:       props,
		selfClosing: false,
	}
}

type MapProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
	Name  string
}

func (p *MapProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Name != "" {
		err := writeAttr("name", p.Name, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Map(props *MapProps, children ...Component) Component {
	return &tag[*MapProps]{
		children:    children,
		name:        "map",
		props:       props,
		selfClosing: false,
	}
}

type MarkProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *MarkProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Mark(props *MarkProps, children ...Component) Component {
	return &tag[*MarkProps]{
		children:    children,
		name:        "mark",
		props:       props,
		selfClosing: false,
	}
}

type MenuProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *MenuProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Menu(props *MenuProps, children ...Component) Component {
	return &tag[*MenuProps]{
		children:    children,
		name:        "menu",
		props:       props,
		selfClosing: false,
	}
}

type MetaProps struct {
	*GenericProps
	X         map[string]string
	ID        string
	Class     string
	Style     map[string]string
	Name      string
	HttpEquiv string
	Content   string
	Charset   string
	Media     string
}

func (p *MetaProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Name != "" {
		err := writeAttr("name", p.Name, w)
		if err != nil {
			return err
		}
	}
	if p.HttpEquiv != "" {
		err := writeAttr("httpequiv", p.HttpEquiv, w)
		if err != nil {
			return err
		}
	}
	if p.Content != "" {
		err := writeAttr("content", p.Content, w)
		if err != nil {
			return err
		}
	}
	if p.Charset != "" {
		err := writeAttr("charset", p.Charset, w)
		if err != nil {
			return err
		}
	}
	if p.Media != "" {
		err := writeAttr("media", p.Media, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Meta(props *MetaProps, children ...Component) Component {
	return &tag[*MetaProps]{
		children:    children,
		name:        "meta",
		props:       props,
		selfClosing: true,
	}
}

type MeterProps struct {
	*GenericProps
	X       map[string]string
	ID      string
	Class   string
	Style   map[string]string
	Value   string
	Min     string
	Max     string
	Low     string
	High    string
	Optimum string
}

func (p *MeterProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Value != "" {
		err := writeAttr("value", p.Value, w)
		if err != nil {
			return err
		}
	}
	if p.Min != "" {
		err := writeAttr("min", p.Min, w)
		if err != nil {
			return err
		}
	}
	if p.Max != "" {
		err := writeAttr("max", p.Max, w)
		if err != nil {
			return err
		}
	}
	if p.Low != "" {
		err := writeAttr("low", p.Low, w)
		if err != nil {
			return err
		}
	}
	if p.High != "" {
		err := writeAttr("high", p.High, w)
		if err != nil {
			return err
		}
	}
	if p.Optimum != "" {
		err := writeAttr("optimum", p.Optimum, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Meter(props *MeterProps, children ...Component) Component {
	return &tag[*MeterProps]{
		children:    children,
		name:        "meter",
		props:       props,
		selfClosing: false,
	}
}

type NavProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *NavProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Nav(props *NavProps, children ...Component) Component {
	return &tag[*NavProps]{
		children:    children,
		name:        "nav",
		props:       props,
		selfClosing: false,
	}
}

type NoscriptProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *NoscriptProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Noscript(props *NoscriptProps, children ...Component) Component {
	return &tag[*NoscriptProps]{
		children:    children,
		name:        "noscript",
		props:       props,
		selfClosing: false,
	}
}

type ObjectProps struct {
	*GenericProps
	X      map[string]string
	ID     string
	Class  string
	Style  map[string]string
	Data   string
	Type   string
	Name   string
	Form   string
	Width  string
	Height string
}

func (p *ObjectProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Data != "" {
		err := writeAttr("data", p.Data, w)
		if err != nil {
			return err
		}
	}
	if p.Type != "" {
		err := writeAttr("type", p.Type, w)
		if err != nil {
			return err
		}
	}
	if p.Name != "" {
		err := writeAttr("name", p.Name, w)
		if err != nil {
			return err
		}
	}
	if p.Form != "" {
		err := writeAttr("form", p.Form, w)
		if err != nil {
			return err
		}
	}
	if p.Width != "" {
		err := writeAttr("width", p.Width, w)
		if err != nil {
			return err
		}
	}
	if p.Height != "" {
		err := writeAttr("height", p.Height, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Object(props *ObjectProps, children ...Component) Component {
	return &tag[*ObjectProps]{
		children:    children,
		name:        "object",
		props:       props,
		selfClosing: false,
	}
}

type OlProps struct {
	*GenericProps
	X        map[string]string
	ID       string
	Class    string
	Style    map[string]string
	Reversed string
	Start    string
	Type     string
}

func (p *OlProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Reversed != "" {
		err := writeAttr("reversed", p.Reversed, w)
		if err != nil {
			return err
		}
	}
	if p.Start != "" {
		err := writeAttr("start", p.Start, w)
		if err != nil {
			return err
		}
	}
	if p.Type != "" {
		err := writeAttr("type", p.Type, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Ol(props *OlProps, children ...Component) Component {
	return &tag[*OlProps]{
		children:    children,
		name:        "ol",
		props:       props,
		selfClosing: false,
	}
}

type OptgroupProps struct {
	*GenericProps
	X        map[string]string
	ID       string
	Class    string
	Style    map[string]string
	Disabled string
	Label    string
}

func (p *OptgroupProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Disabled != "" {
		err := writeAttr("disabled", p.Disabled, w)
		if err != nil {
			return err
		}
	}
	if p.Label != "" {
		err := writeAttr("label", p.Label, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Optgroup(props *OptgroupProps, children ...Component) Component {
	return &tag[*OptgroupProps]{
		children:    children,
		name:        "optgroup",
		props:       props,
		selfClosing: false,
	}
}

type OptionProps struct {
	*GenericProps
	X        map[string]string
	ID       string
	Class    string
	Style    map[string]string
	Disabled string
	Label    string
	Selected string
	Value    string
}

func (p *OptionProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Disabled != "" {
		err := writeAttr("disabled", p.Disabled, w)
		if err != nil {
			return err
		}
	}
	if p.Label != "" {
		err := writeAttr("label", p.Label, w)
		if err != nil {
			return err
		}
	}
	if p.Selected != "" {
		err := writeAttr("selected", p.Selected, w)
		if err != nil {
			return err
		}
	}
	if p.Value != "" {
		err := writeAttr("value", p.Value, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Option(props *OptionProps, children ...Component) Component {
	return &tag[*OptionProps]{
		children:    children,
		name:        "option",
		props:       props,
		selfClosing: false,
	}
}

type OutputProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
	For   string
	Form  string
	Name  string
}

func (p *OutputProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.For != "" {
		err := writeAttr("for", p.For, w)
		if err != nil {
			return err
		}
	}
	if p.Form != "" {
		err := writeAttr("form", p.Form, w)
		if err != nil {
			return err
		}
	}
	if p.Name != "" {
		err := writeAttr("name", p.Name, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Output(props *OutputProps, children ...Component) Component {
	return &tag[*OutputProps]{
		children:    children,
		name:        "output",
		props:       props,
		selfClosing: false,
	}
}

type PProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *PProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func P(props *PProps, children ...Component) Component {
	return &tag[*PProps]{
		children:    children,
		name:        "p",
		props:       props,
		selfClosing: false,
	}
}

type PictureProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *PictureProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Picture(props *PictureProps, children ...Component) Component {
	return &tag[*PictureProps]{
		children:    children,
		name:        "picture",
		props:       props,
		selfClosing: false,
	}
}

type PreProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *PreProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Pre(props *PreProps, children ...Component) Component {
	return &tag[*PreProps]{
		children:    children,
		name:        "pre",
		props:       props,
		selfClosing: false,
	}
}

type ProgressProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
	Value string
	Max   string
}

func (p *ProgressProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Value != "" {
		err := writeAttr("value", p.Value, w)
		if err != nil {
			return err
		}
	}
	if p.Max != "" {
		err := writeAttr("max", p.Max, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Progress(props *ProgressProps, children ...Component) Component {
	return &tag[*ProgressProps]{
		children:    children,
		name:        "progress",
		props:       props,
		selfClosing: false,
	}
}

type QProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
	Cite  string
}

func (p *QProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Cite != "" {
		err := writeAttr("cite", p.Cite, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Q(props *QProps, children ...Component) Component {
	return &tag[*QProps]{
		children:    children,
		name:        "q",
		props:       props,
		selfClosing: false,
	}
}

type RpProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *RpProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Rp(props *RpProps, children ...Component) Component {
	return &tag[*RpProps]{
		children:    children,
		name:        "rp",
		props:       props,
		selfClosing: false,
	}
}

type RtProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *RtProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Rt(props *RtProps, children ...Component) Component {
	return &tag[*RtProps]{
		children:    children,
		name:        "rt",
		props:       props,
		selfClosing: false,
	}
}

type RubyProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *RubyProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Ruby(props *RubyProps, children ...Component) Component {
	return &tag[*RubyProps]{
		children:    children,
		name:        "ruby",
		props:       props,
		selfClosing: false,
	}
}

type SProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *SProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func S(props *SProps, children ...Component) Component {
	return &tag[*SProps]{
		children:    children,
		name:        "s",
		props:       props,
		selfClosing: false,
	}
}

type SampProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *SampProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Samp(props *SampProps, children ...Component) Component {
	return &tag[*SampProps]{
		children:    children,
		name:        "samp",
		props:       props,
		selfClosing: false,
	}
}

type ScriptProps struct {
	*GenericProps
	X              map[string]string
	ID             string
	Class          string
	Style          map[string]string
	Src            string
	Type           string
	Nomodule       string
	Async          string
	Defer          string
	Crossorigin    string
	Integrity      string
	Referrerpolicy string
	Blocking       string
	Fetchpriority  string
}

func (p *ScriptProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Src != "" {
		err := writeAttr("src", p.Src, w)
		if err != nil {
			return err
		}
	}
	if p.Type != "" {
		err := writeAttr("type", p.Type, w)
		if err != nil {
			return err
		}
	}
	if p.Nomodule != "" {
		err := writeAttr("nomodule", p.Nomodule, w)
		if err != nil {
			return err
		}
	}
	if p.Async != "" {
		err := writeAttr("async", p.Async, w)
		if err != nil {
			return err
		}
	}
	if p.Defer != "" {
		err := writeAttr("defer", p.Defer, w)
		if err != nil {
			return err
		}
	}
	if p.Crossorigin != "" {
		err := writeAttr("crossorigin", p.Crossorigin, w)
		if err != nil {
			return err
		}
	}
	if p.Integrity != "" {
		err := writeAttr("integrity", p.Integrity, w)
		if err != nil {
			return err
		}
	}
	if p.Referrerpolicy != "" {
		err := writeAttr("referrerpolicy", p.Referrerpolicy, w)
		if err != nil {
			return err
		}
	}
	if p.Blocking != "" {
		err := writeAttr("blocking", p.Blocking, w)
		if err != nil {
			return err
		}
	}
	if p.Fetchpriority != "" {
		err := writeAttr("fetchpriority", p.Fetchpriority, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Script(props *ScriptProps, children ...Component) Component {
	return &tag[*ScriptProps]{
		children:    children,
		name:        "script",
		props:       props,
		selfClosing: false,
	}
}

type SearchProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *SearchProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Search(props *SearchProps, children ...Component) Component {
	return &tag[*SearchProps]{
		children:    children,
		name:        "search",
		props:       props,
		selfClosing: false,
	}
}

type SectionProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *SectionProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Section(props *SectionProps, children ...Component) Component {
	return &tag[*SectionProps]{
		children:    children,
		name:        "section",
		props:       props,
		selfClosing: false,
	}
}

type SelectProps struct {
	*GenericProps
	X            map[string]string
	ID           string
	Class        string
	Style        map[string]string
	Autocomplete string
	Disabled     string
	Form         string
	Multiple     string
	Name         string
	Required     string
	Size         string
}

func (p *SelectProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Autocomplete != "" {
		err := writeAttr("autocomplete", p.Autocomplete, w)
		if err != nil {
			return err
		}
	}
	if p.Disabled != "" {
		err := writeAttr("disabled", p.Disabled, w)
		if err != nil {
			return err
		}
	}
	if p.Form != "" {
		err := writeAttr("form", p.Form, w)
		if err != nil {
			return err
		}
	}
	if p.Multiple != "" {
		err := writeAttr("multiple", p.Multiple, w)
		if err != nil {
			return err
		}
	}
	if p.Name != "" {
		err := writeAttr("name", p.Name, w)
		if err != nil {
			return err
		}
	}
	if p.Required != "" {
		err := writeAttr("required", p.Required, w)
		if err != nil {
			return err
		}
	}
	if p.Size != "" {
		err := writeAttr("size", p.Size, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Select(props *SelectProps, children ...Component) Component {
	return &tag[*SelectProps]{
		children:    children,
		name:        "select",
		props:       props,
		selfClosing: false,
	}
}

type SlotProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
	Name  string
}

func (p *SlotProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Name != "" {
		err := writeAttr("name", p.Name, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Slot(props *SlotProps, children ...Component) Component {
	return &tag[*SlotProps]{
		children:    children,
		name:        "slot",
		props:       props,
		selfClosing: false,
	}
}

type SmallProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *SmallProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Small(props *SmallProps, children ...Component) Component {
	return &tag[*SmallProps]{
		children:    children,
		name:        "small",
		props:       props,
		selfClosing: false,
	}
}

type SourceProps struct {
	*GenericProps
	X      map[string]string
	ID     string
	Class  string
	Style  map[string]string
	Type   string
	Media  string
	Src    string
	Srcset string
	Sizes  string
	Width  string
	Height string
}

func (p *SourceProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Type != "" {
		err := writeAttr("type", p.Type, w)
		if err != nil {
			return err
		}
	}
	if p.Media != "" {
		err := writeAttr("media", p.Media, w)
		if err != nil {
			return err
		}
	}
	if p.Src != "" {
		err := writeAttr("src", p.Src, w)
		if err != nil {
			return err
		}
	}
	if p.Srcset != "" {
		err := writeAttr("srcset", p.Srcset, w)
		if err != nil {
			return err
		}
	}
	if p.Sizes != "" {
		err := writeAttr("sizes", p.Sizes, w)
		if err != nil {
			return err
		}
	}
	if p.Width != "" {
		err := writeAttr("width", p.Width, w)
		if err != nil {
			return err
		}
	}
	if p.Height != "" {
		err := writeAttr("height", p.Height, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Source(props *SourceProps, children ...Component) Component {
	return &tag[*SourceProps]{
		children:    children,
		name:        "source",
		props:       props,
		selfClosing: true,
	}
}

type SpanProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *SpanProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Span(props *SpanProps, children ...Component) Component {
	return &tag[*SpanProps]{
		children:    children,
		name:        "span",
		props:       props,
		selfClosing: false,
	}
}

type StrongProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *StrongProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Strong(props *StrongProps, children ...Component) Component {
	return &tag[*StrongProps]{
		children:    children,
		name:        "strong",
		props:       props,
		selfClosing: false,
	}
}

type StyleProps struct {
	*GenericProps
	X        map[string]string
	ID       string
	Class    string
	Style    map[string]string
	Media    string
	Blocking string
}

func (p *StyleProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Media != "" {
		err := writeAttr("media", p.Media, w)
		if err != nil {
			return err
		}
	}
	if p.Blocking != "" {
		err := writeAttr("blocking", p.Blocking, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Style(props *StyleProps, children ...Component) Component {
	return &tag[*StyleProps]{
		children:    children,
		name:        "style",
		props:       props,
		selfClosing: false,
	}
}

type SubProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *SubProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Sub(props *SubProps, children ...Component) Component {
	return &tag[*SubProps]{
		children:    children,
		name:        "sub",
		props:       props,
		selfClosing: false,
	}
}

type SummaryProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *SummaryProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Summary(props *SummaryProps, children ...Component) Component {
	return &tag[*SummaryProps]{
		children:    children,
		name:        "summary",
		props:       props,
		selfClosing: false,
	}
}

type SupProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *SupProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Sup(props *SupProps, children ...Component) Component {
	return &tag[*SupProps]{
		children:    children,
		name:        "sup",
		props:       props,
		selfClosing: false,
	}
}

type TableProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *TableProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Table(props *TableProps, children ...Component) Component {
	return &tag[*TableProps]{
		children:    children,
		name:        "table",
		props:       props,
		selfClosing: false,
	}
}

type TbodyProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *TbodyProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Tbody(props *TbodyProps, children ...Component) Component {
	return &tag[*TbodyProps]{
		children:    children,
		name:        "tbody",
		props:       props,
		selfClosing: false,
	}
}

type TdProps struct {
	*GenericProps
	X       map[string]string
	ID      string
	Class   string
	Style   map[string]string
	Colspan string
	Rowspan string
	Headers string
}

func (p *TdProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Colspan != "" {
		err := writeAttr("colspan", p.Colspan, w)
		if err != nil {
			return err
		}
	}
	if p.Rowspan != "" {
		err := writeAttr("rowspan", p.Rowspan, w)
		if err != nil {
			return err
		}
	}
	if p.Headers != "" {
		err := writeAttr("headers", p.Headers, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Td(props *TdProps, children ...Component) Component {
	return &tag[*TdProps]{
		children:    children,
		name:        "td",
		props:       props,
		selfClosing: false,
	}
}

type TemplateProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *TemplateProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Template(props *TemplateProps, children ...Component) Component {
	return &tag[*TemplateProps]{
		children:    children,
		name:        "template",
		props:       props,
		selfClosing: true,
	}
}

type TextareaProps struct {
	*GenericProps
	X            map[string]string
	ID           string
	Class        string
	Style        map[string]string
	Autocomplete string
	Cols         string
	Dirname      string
	Disabled     string
	Form         string
	Maxlength    string
	Minlength    string
	Name         string
	Placeholder  string
	Readonly     string
	Required     string
	Rows         string
	Wrap         string
}

func (p *TextareaProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Autocomplete != "" {
		err := writeAttr("autocomplete", p.Autocomplete, w)
		if err != nil {
			return err
		}
	}
	if p.Cols != "" {
		err := writeAttr("cols", p.Cols, w)
		if err != nil {
			return err
		}
	}
	if p.Dirname != "" {
		err := writeAttr("dirname", p.Dirname, w)
		if err != nil {
			return err
		}
	}
	if p.Disabled != "" {
		err := writeAttr("disabled", p.Disabled, w)
		if err != nil {
			return err
		}
	}
	if p.Form != "" {
		err := writeAttr("form", p.Form, w)
		if err != nil {
			return err
		}
	}
	if p.Maxlength != "" {
		err := writeAttr("maxlength", p.Maxlength, w)
		if err != nil {
			return err
		}
	}
	if p.Minlength != "" {
		err := writeAttr("minlength", p.Minlength, w)
		if err != nil {
			return err
		}
	}
	if p.Name != "" {
		err := writeAttr("name", p.Name, w)
		if err != nil {
			return err
		}
	}
	if p.Placeholder != "" {
		err := writeAttr("placeholder", p.Placeholder, w)
		if err != nil {
			return err
		}
	}
	if p.Readonly != "" {
		err := writeAttr("readonly", p.Readonly, w)
		if err != nil {
			return err
		}
	}
	if p.Required != "" {
		err := writeAttr("required", p.Required, w)
		if err != nil {
			return err
		}
	}
	if p.Rows != "" {
		err := writeAttr("rows", p.Rows, w)
		if err != nil {
			return err
		}
	}
	if p.Wrap != "" {
		err := writeAttr("wrap", p.Wrap, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Textarea(props *TextareaProps, children ...Component) Component {
	return &tag[*TextareaProps]{
		children:    children,
		name:        "textarea",
		props:       props,
		selfClosing: false,
	}
}

type TfootProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *TfootProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Tfoot(props *TfootProps, children ...Component) Component {
	return &tag[*TfootProps]{
		children:    children,
		name:        "tfoot",
		props:       props,
		selfClosing: false,
	}
}

type ThProps struct {
	*GenericProps
	X       map[string]string
	ID      string
	Class   string
	Style   map[string]string
	Colspan string
	Rowspan string
	Headers string
	Scope   string
	Abbr    string
}

func (p *ThProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Colspan != "" {
		err := writeAttr("colspan", p.Colspan, w)
		if err != nil {
			return err
		}
	}
	if p.Rowspan != "" {
		err := writeAttr("rowspan", p.Rowspan, w)
		if err != nil {
			return err
		}
	}
	if p.Headers != "" {
		err := writeAttr("headers", p.Headers, w)
		if err != nil {
			return err
		}
	}
	if p.Scope != "" {
		err := writeAttr("scope", p.Scope, w)
		if err != nil {
			return err
		}
	}
	if p.Abbr != "" {
		err := writeAttr("abbr", p.Abbr, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Th(props *ThProps, children ...Component) Component {
	return &tag[*ThProps]{
		children:    children,
		name:        "th",
		props:       props,
		selfClosing: false,
	}
}

type TheadProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *TheadProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Thead(props *TheadProps, children ...Component) Component {
	return &tag[*TheadProps]{
		children:    children,
		name:        "thead",
		props:       props,
		selfClosing: false,
	}
}

type TimeProps struct {
	*GenericProps
	X        map[string]string
	ID       string
	Class    string
	Style    map[string]string
	Datetime string
}

func (p *TimeProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Datetime != "" {
		err := writeAttr("datetime", p.Datetime, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Time(props *TimeProps, children ...Component) Component {
	return &tag[*TimeProps]{
		children:    children,
		name:        "time",
		props:       props,
		selfClosing: false,
	}
}

type TitleProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *TitleProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Title(props *TitleProps, children ...Component) Component {
	return &tag[*TitleProps]{
		children:    children,
		name:        "title",
		props:       props,
		selfClosing: false,
	}
}

type TrProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *TrProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Tr(props *TrProps, children ...Component) Component {
	return &tag[*TrProps]{
		children:    children,
		name:        "tr",
		props:       props,
		selfClosing: false,
	}
}

type TrackProps struct {
	*GenericProps
	X       map[string]string
	ID      string
	Class   string
	Style   map[string]string
	Default string
	Kind    string
	Label   string
	Src     string
	Srclang string
}

func (p *TrackProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Default != "" {
		err := writeAttr("default", p.Default, w)
		if err != nil {
			return err
		}
	}
	if p.Kind != "" {
		err := writeAttr("kind", p.Kind, w)
		if err != nil {
			return err
		}
	}
	if p.Label != "" {
		err := writeAttr("label", p.Label, w)
		if err != nil {
			return err
		}
	}
	if p.Src != "" {
		err := writeAttr("src", p.Src, w)
		if err != nil {
			return err
		}
	}
	if p.Srclang != "" {
		err := writeAttr("srclang", p.Srclang, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Track(props *TrackProps, children ...Component) Component {
	return &tag[*TrackProps]{
		children:    children,
		name:        "track",
		props:       props,
		selfClosing: true,
	}
}

type UProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *UProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func U(props *UProps, children ...Component) Component {
	return &tag[*UProps]{
		children:    children,
		name:        "u",
		props:       props,
		selfClosing: false,
	}
}

type UlProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *UlProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Ul(props *UlProps, children ...Component) Component {
	return &tag[*UlProps]{
		children:    children,
		name:        "ul",
		props:       props,
		selfClosing: false,
	}
}

type VarProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *VarProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Var(props *VarProps, children ...Component) Component {
	return &tag[*VarProps]{
		children:    children,
		name:        "var",
		props:       props,
		selfClosing: false,
	}
}

type VideoProps struct {
	*GenericProps
	X           map[string]string
	ID          string
	Class       string
	Style       map[string]string
	Src         string
	Crossorigin string
	Poster      string
	Preload     string
	Autoplay    string
	Playsinline string
	Loop        string
	Muted       string
	Controls    string
	Width       string
	Height      string
}

func (p *VideoProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.Src != "" {
		err := writeAttr("src", p.Src, w)
		if err != nil {
			return err
		}
	}
	if p.Crossorigin != "" {
		err := writeAttr("crossorigin", p.Crossorigin, w)
		if err != nil {
			return err
		}
	}
	if p.Poster != "" {
		err := writeAttr("poster", p.Poster, w)
		if err != nil {
			return err
		}
	}
	if p.Preload != "" {
		err := writeAttr("preload", p.Preload, w)
		if err != nil {
			return err
		}
	}
	if p.Autoplay != "" {
		err := writeAttr("autoplay", p.Autoplay, w)
		if err != nil {
			return err
		}
	}
	if p.Playsinline != "" {
		err := writeAttr("playsinline", p.Playsinline, w)
		if err != nil {
			return err
		}
	}
	if p.Loop != "" {
		err := writeAttr("loop", p.Loop, w)
		if err != nil {
			return err
		}
	}
	if p.Muted != "" {
		err := writeAttr("muted", p.Muted, w)
		if err != nil {
			return err
		}
	}
	if p.Controls != "" {
		err := writeAttr("controls", p.Controls, w)
		if err != nil {
			return err
		}
	}
	if p.Width != "" {
		err := writeAttr("width", p.Width, w)
		if err != nil {
			return err
		}
	}
	if p.Height != "" {
		err := writeAttr("height", p.Height, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Video(props *VideoProps, children ...Component) Component {
	return &tag[*VideoProps]{
		children:    children,
		name:        "video",
		props:       props,
		selfClosing: false,
	}
}

type WbrProps struct {
	*GenericProps
	X     map[string]string
	ID    string
	Class string
	Style map[string]string
}

func (p *WbrProps) writeTo(w io.Writer) error {
	if p == nil {
		return nil
	}
	if p.GenericProps != nil {
		err := p.GenericProps.writeTo(w)
		if err != nil {
			return err
		}
	}
	if p.Style != nil {
		err := writeStyle(p.Style, w)
		if err != nil {
			return err
		}
	}
	if p.ID != "" {
		err := writeAttr("id", p.ID, w)
		if err != nil {
			return err
		}
	}
	if p.Class != "" {
		err := writeAttr("class", p.Class, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func Wbr(props *WbrProps, children ...Component) Component {
	return &tag[*WbrProps]{
		children:    children,
		name:        "wbr",
		props:       props,
		selfClosing: true,
	}
}
