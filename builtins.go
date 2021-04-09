package hkit

// VarAttrs .
type VarAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Var HTML element.
func Var(attrs *VarAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "var",
		}
	}
}

// AreaAttrs .
type AreaAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
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

// Area HTML element.
func Area(attrs *AreaAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "area",
		}
	}
}

// DialogAttrs .
type DialogAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
	Open  string
}

// Dialog HTML element.
func Dialog(attrs *DialogAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "dialog",
		}
	}
}

// ProgressAttrs .
type ProgressAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
	Value string
	Max   string
}

// Progress HTML element.
func Progress(attrs *ProgressAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "progress",
		}
	}
}

// SampAttrs .
type SampAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Samp HTML element.
func Samp(attrs *SampAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "samp",
		}
	}
}

// RubyAttrs .
type RubyAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Ruby HTML element.
func Ruby(attrs *RubyAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "ruby",
		}
	}
}

// BAttrs .
type BAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// B HTML element.
func B(attrs *BAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "b",
		}
	}
}

// EmAttrs .
type EmAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Em HTML element.
func Em(attrs *EmAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "em",
		}
	}
}

// FigcaptionAttrs .
type FigcaptionAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Figcaption HTML element.
func Figcaption(attrs *FigcaptionAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "figcaption",
		}
	}
}

// LinkAttrs .
type LinkAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
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
	Color          string
	Disabled       string
}

// Link HTML element.
func Link(attrs *LinkAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "link",
		}
	}
}

// H5Attrs .
type H5Attrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// H5 HTML element.
func H5(attrs *H5Attrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "h5",
		}
	}
}

// LabelAttrs .
type LabelAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
	For   string
}

// Label HTML element.
func Label(attrs *LabelAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "label",
		}
	}
}

// OutputAttrs .
type OutputAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
	For   string
	Form  string
	Name  string
}

// Output HTML element.
func Output(attrs *OutputAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "output",
		}
	}
}

// PictureAttrs .
type PictureAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Picture HTML element.
func Picture(attrs *PictureAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "picture",
		}
	}
}

// TheadAttrs .
type TheadAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Thead HTML element.
func Thead(attrs *TheadAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "thead",
		}
	}
}

// DfnAttrs .
type DfnAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Dfn HTML element.
func Dfn(attrs *DfnAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "dfn",
		}
	}
}

// FooterAttrs .
type FooterAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Footer HTML element.
func Footer(attrs *FooterAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "footer",
		}
	}
}

// MenuAttrs .
type MenuAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Menu HTML element.
func Menu(attrs *MenuAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "menu",
		}
	}
}

// LiAttrs .
type LiAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
	Value string
}

// Li HTML element.
func Li(attrs *LiAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "li",
		}
	}
}

// MarkAttrs .
type MarkAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Mark HTML element.
func Mark(attrs *MarkAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "mark",
		}
	}
}

// SectionAttrs .
type SectionAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Section HTML element.
func Section(attrs *SectionAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "section",
		}
	}
}

// EmbedAttrs .
type EmbedAttrs struct {
	GlobalAttrs
	ID     string
	Class  string
	Style  string
	Src    string
	Type   string
	Width  string
	Height string
	Any    string
}

// Embed HTML element.
func Embed(attrs *EmbedAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "embed",
		}
	}
}

// HeaderAttrs .
type HeaderAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Header HTML element.
func Header(attrs *HeaderAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "header",
		}
	}
}

// LegendAttrs .
type LegendAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Legend HTML element.
func Legend(attrs *LegendAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "legend",
		}
	}
}

// OptgroupAttrs .
type OptgroupAttrs struct {
	GlobalAttrs
	ID       string
	Class    string
	Style    string
	Disabled string
	Label    string
}

// Optgroup HTML element.
func Optgroup(attrs *OptgroupAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "optgroup",
		}
	}
}

// HtmlAttrs .
type HtmlAttrs struct {
	GlobalAttrs
	ID       string
	Class    string
	Style    string
	Manifest string
}

// Html HTML element.
func Html(attrs *HtmlAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "html",
		}
	}
}

// SourceAttrs .
type SourceAttrs struct {
	GlobalAttrs
	ID     string
	Class  string
	Style  string
	Src    string
	Type   string
	Srcset string
	Sizes  string
	Media  string
	Width  string
	Height string
}

// Source HTML element.
func Source(attrs *SourceAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "source",
		}
	}
}

// TrAttrs .
type TrAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Tr HTML element.
func Tr(attrs *TrAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "tr",
		}
	}
}

// AsideAttrs .
type AsideAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Aside HTML element.
func Aside(attrs *AsideAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "aside",
		}
	}
}

// BodyAttrs .
type BodyAttrs struct {
	GlobalAttrs
	ID                   string
	Class                string
	Style                string
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

// Body HTML element.
func Body(attrs *BodyAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "body",
		}
	}
}

// DetailsAttrs .
type DetailsAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
	Open  string
}

// Details HTML element.
func Details(attrs *DetailsAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "details",
		}
	}
}

// OlAttrs .
type OlAttrs struct {
	GlobalAttrs
	ID       string
	Class    string
	Style    string
	Reversed string
	Start    string
	Type     string
}

// Ol HTML element.
func Ol(attrs *OlAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "ol",
		}
	}
}

// UlAttrs .
type UlAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Ul HTML element.
func Ul(attrs *UlAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "ul",
		}
	}
}

// DataAttrs .
type DataAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
	Value string
}

// Data HTML element.
func Data(attrs *DataAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "data",
		}
	}
}

// H1Attrs .
type H1Attrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// H1 HTML element.
func H1(attrs *H1Attrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "h1",
		}
	}
}

// ParamAttrs .
type ParamAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
	Name  string
	Value string
}

// Param HTML element.
func Param(attrs *ParamAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "param",
		}
	}
}

// SubAttrs .
type SubAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Sub HTML element.
func Sub(attrs *SubAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "sub",
		}
	}
}

// BaseAttrs .
type BaseAttrs struct {
	GlobalAttrs
	ID     string
	Class  string
	Style  string
	Href   string
	Target string
}

// Base HTML element.
func Base(attrs *BaseAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "base",
		}
	}
}

// OptionAttrs .
type OptionAttrs struct {
	GlobalAttrs
	ID       string
	Class    string
	Style    string
	Disabled string
	Label    string
	Selected string
	Value    string
}

// Option HTML element.
func Option(attrs *OptionAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "option",
		}
	}
}

// HeadAttrs .
type HeadAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Head HTML element.
func Head(attrs *HeadAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "head",
		}
	}
}

// StrongAttrs .
type StrongAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Strong HTML element.
func Strong(attrs *StrongAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "strong",
		}
	}
}

// UAttrs .
type UAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// U HTML element.
func U(attrs *UAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "u",
		}
	}
}

// AddressAttrs .
type AddressAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Address HTML element.
func Address(attrs *AddressAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "address",
		}
	}
}

// MetaAttrs .
type MetaAttrs struct {
	GlobalAttrs
	ID        string
	Class     string
	Style     string
	Name      string
	HttpEquiv string
	Content   string
	Charset   string
}

// Meta HTML element.
func Meta(attrs *MetaAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "meta",
		}
	}
}

// SmallAttrs .
type SmallAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Small HTML element.
func Small(attrs *SmallAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "small",
		}
	}
}

// TrackAttrs .
type TrackAttrs struct {
	GlobalAttrs
	ID      string
	Class   string
	Style   string
	Default string
	Kind    string
	Label   string
	Src     string
	Srclang string
}

// Track HTML element.
func Track(attrs *TrackAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "track",
		}
	}
}

// BlockquoteAttrs .
type BlockquoteAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
	Cite  string
}

// Blockquote HTML element.
func Blockquote(attrs *BlockquoteAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "blockquote",
		}
	}
}

// DlAttrs .
type DlAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Dl HTML element.
func Dl(attrs *DlAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "dl",
		}
	}
}

// HgroupAttrs .
type HgroupAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Hgroup HTML element.
func Hgroup(attrs *HgroupAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "hgroup",
		}
	}
}

// VideoAttrs .
type VideoAttrs struct {
	GlobalAttrs
	ID          string
	Class       string
	Style       string
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

// Video HTML element.
func Video(attrs *VideoAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "video",
		}
	}
}

// BdiAttrs .
type BdiAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Bdi HTML element.
func Bdi(attrs *BdiAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "bdi",
		}
	}
}

// InputAttrs .
type InputAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Accept         string
	Alt            string
	Autocomplete   string
	Checked        string
	Dirname        string
	Disabled       string
	Form           string
	Formaction     string
	Formenctype    string
	Formmethod     string
	Formnovalidate string
	Formtarget     string
	Height         string
	List           string
	Max            string
	Maxlength      string
	Min            string
	Minlength      string
	Multiple       string
	Name           string
	Pattern        string
	Placeholder    string
	Readonly       string
	Required       string
	Size           string
	Src            string
	Step           string
	Type           string
	Value          string
	Width          string
}

// Input HTML element.
func Input(attrs *InputAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "input",
		}
	}
}

// PreAttrs .
type PreAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Pre HTML element.
func Pre(attrs *PreAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "pre",
		}
	}
}

// TableAttrs .
type TableAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Table HTML element.
func Table(attrs *TableAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "table",
		}
	}
}

// H4Attrs .
type H4Attrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// H4 HTML element.
func H4(attrs *H4Attrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "h4",
		}
	}
}

// MeterAttrs .
type MeterAttrs struct {
	GlobalAttrs
	ID      string
	Class   string
	Style   string
	Value   string
	Min     string
	Max     string
	Low     string
	High    string
	Optimum string
}

// Meter HTML element.
func Meter(attrs *MeterAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "meter",
		}
	}
}

// RtAttrs .
type RtAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Rt HTML element.
func Rt(attrs *RtAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "rt",
		}
	}
}

// HrAttrs .
type HrAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Hr HTML element.
func Hr(attrs *HrAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "hr",
		}
	}
}

// IAttrs .
type IAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// I HTML element.
func I(attrs *IAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "i",
		}
	}
}

// PAttrs .
type PAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// P HTML element.
func P(attrs *PAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "p",
		}
	}
}

// SummaryAttrs .
type SummaryAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Summary HTML element.
func Summary(attrs *SummaryAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "summary",
		}
	}
}

// KbdAttrs .
type KbdAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Kbd HTML element.
func Kbd(attrs *KbdAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "kbd",
		}
	}
}

// H6Attrs .
type H6Attrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// H6 HTML element.
func H6(attrs *H6Attrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "h6",
		}
	}
}

// MainAttrs .
type MainAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Main HTML element.
func Main(attrs *MainAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "main",
		}
	}
}

// BdoAttrs .
type BdoAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Bdo HTML element.
func Bdo(attrs *BdoAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "bdo",
		}
	}
}

// BrAttrs .
type BrAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Br HTML element.
func Br(attrs *BrAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "br",
		}
	}
}

// DatalistAttrs .
type DatalistAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Datalist HTML element.
func Datalist(attrs *DatalistAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "datalist",
		}
	}
}

// FormAttrs .
type FormAttrs struct {
	GlobalAttrs
	ID            string
	Class         string
	Style         string
	AcceptCharset string
	Action        string
	Autocomplete  string
	Enctype       string
	Method        string
	Name          string
	Novalidate    string
	Target        string
}

// Form HTML element.
func Form(attrs *FormAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "form",
		}
	}
}

// NavAttrs .
type NavAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Nav HTML element.
func Nav(attrs *NavAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "nav",
		}
	}
}

// RpAttrs .
type RpAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Rp HTML element.
func Rp(attrs *RpAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "rp",
		}
	}
}

// TfootAttrs .
type TfootAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Tfoot HTML element.
func Tfoot(attrs *TfootAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "tfoot",
		}
	}
}

// ButtonAttrs .
type ButtonAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Disabled       string
	Form           string
	Formaction     string
	Formenctype    string
	Formmethod     string
	Formnovalidate string
	Formtarget     string
	Name           string
	Type           string
	Value          string
}

// Button HTML element.
func Button(attrs *ButtonAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "button",
		}
	}
}

// TimeAttrs .
type TimeAttrs struct {
	GlobalAttrs
	ID       string
	Class    string
	Style    string
	Datetime string
}

// Time HTML element.
func Time(attrs *TimeAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "time",
		}
	}
}

// ScriptAttrs .
type ScriptAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Src            string
	Type           string
	Async          string
	Defer          string
	Crossorigin    string
	Integrity      string
	Referrerpolicy string
}

// Script HTML element.
func Script(attrs *ScriptAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "script",
		}
	}
}

// DtAttrs .
type DtAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Dt HTML element.
func Dt(attrs *DtAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "dt",
		}
	}
}

// FigureAttrs .
type FigureAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Figure HTML element.
func Figure(attrs *FigureAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "figure",
		}
	}
}

// ObjectAttrs .
type ObjectAttrs struct {
	GlobalAttrs
	ID     string
	Class  string
	Style  string
	Data   string
	Type   string
	Name   string
	Form   string
	Width  string
	Height string
}

// Object HTML element.
func Object(attrs *ObjectAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "object",
		}
	}
}

// AudioAttrs .
type AudioAttrs struct {
	GlobalAttrs
	ID          string
	Class       string
	Style       string
	Src         string
	Crossorigin string
	Preload     string
	Autoplay    string
	Loop        string
	Muted       string
	Controls    string
}

// Audio HTML element.
func Audio(attrs *AudioAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "audio",
		}
	}
}

// DdAttrs .
type DdAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Dd HTML element.
func Dd(attrs *DdAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "dd",
		}
	}
}

// IframeAttrs .
type IframeAttrs struct {
	GlobalAttrs
	ID              string
	Class           string
	Style           string
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

// Iframe HTML element.
func Iframe(attrs *IframeAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "iframe",
		}
	}
}

// StyleAttrs .
type StyleAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
	Media string
}

// Style HTML element.
func Style(attrs *StyleAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "style",
		}
	}
}

// TdAttrs .
type TdAttrs struct {
	GlobalAttrs
	ID      string
	Class   string
	Style   string
	Colspan string
	Rowspan string
	Headers string
}

// Td HTML element.
func Td(attrs *TdAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "td",
		}
	}
}

// WbrAttrs .
type WbrAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Wbr HTML element.
func Wbr(attrs *WbrAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "wbr",
		}
	}
}

// NoscriptAttrs .
type NoscriptAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Noscript HTML element.
func Noscript(attrs *NoscriptAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "noscript",
		}
	}
}

// SAttrs .
type SAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// S HTML element.
func S(attrs *SAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "s",
		}
	}
}

// SelectAttrs .
type SelectAttrs struct {
	GlobalAttrs
	ID           string
	Class        string
	Style        string
	Autocomplete string
	Disabled     string
	Form         string
	Multiple     string
	Name         string
	Required     string
	Size         string
}

// Select HTML element.
func Select(attrs *SelectAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "select",
		}
	}
}

// SlotAttrs .
type SlotAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
	Name  string
}

// Slot HTML element.
func Slot(attrs *SlotAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "slot",
		}
	}
}

// AbbrAttrs .
type AbbrAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Abbr HTML element.
func Abbr(attrs *AbbrAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "abbr",
		}
	}
}

// ColgroupAttrs .
type ColgroupAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
	Span  string
}

// Colgroup HTML element.
func Colgroup(attrs *ColgroupAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "colgroup",
		}
	}
}

// MapAttrs .
type MapAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
	Name  string
}

// Map HTML element.
func Map(attrs *MapAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "map",
		}
	}
}

// TbodyAttrs .
type TbodyAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Tbody HTML element.
func Tbody(attrs *TbodyAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "tbody",
		}
	}
}

// ColAttrs .
type ColAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
	Span  string
}

// Col HTML element.
func Col(attrs *ColAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "col",
		}
	}
}

// DelAttrs .
type DelAttrs struct {
	GlobalAttrs
	ID       string
	Class    string
	Style    string
	Cite     string
	Datetime string
}

// Del HTML element.
func Del(attrs *DelAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "del",
		}
	}
}

// ImgAttrs .
type ImgAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
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
}

// Img HTML element.
func Img(attrs *ImgAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "img",
		}
	}
}

// CodeAttrs .
type CodeAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Code HTML element.
func Code(attrs *CodeAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "code",
		}
	}
}

// SpanAttrs .
type SpanAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Span HTML element.
func Span(attrs *SpanAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "span",
		}
	}
}

// CanvasAttrs .
type CanvasAttrs struct {
	GlobalAttrs
	ID     string
	Class  string
	Style  string
	Width  string
	Height string
}

// Canvas HTML element.
func Canvas(attrs *CanvasAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "canvas",
		}
	}
}

// DivAttrs .
type DivAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Div HTML element.
func Div(attrs *DivAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "div",
		}
	}
}

// AAttrs .
type AAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// A HTML element.
func A(attrs *AAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "a",
		}
	}
}

// H3Attrs .
type H3Attrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// H3 HTML element.
func H3(attrs *H3Attrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "h3",
		}
	}
}

// ThAttrs .
type ThAttrs struct {
	GlobalAttrs
	ID      string
	Class   string
	Style   string
	Colspan string
	Rowspan string
	Headers string
	Scope   string
	Abbr    string
}

// Th HTML element.
func Th(attrs *ThAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "th",
		}
	}
}

// TitleAttrs .
type TitleAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Title HTML element.
func Title(attrs *TitleAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "title",
		}
	}
}

// CaptionAttrs .
type CaptionAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Caption HTML element.
func Caption(attrs *CaptionAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "caption",
		}
	}
}

// QAttrs .
type QAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
	Cite  string
}

// Q HTML element.
func Q(attrs *QAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "q",
		}
	}
}

// TextareaAttrs .
type TextareaAttrs struct {
	GlobalAttrs
	ID          string
	Class       string
	Style       string
	Cols        string
	Dirname     string
	Disabled    string
	Form        string
	Maxlength   string
	Minlength   string
	Name        string
	Placeholder string
	Readonly    string
	Required    string
	Rows        string
	Wrap        string
}

// Textarea HTML element.
func Textarea(attrs *TextareaAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "textarea",
		}
	}
}

// CiteAttrs .
type CiteAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Cite HTML element.
func Cite(attrs *CiteAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "cite",
		}
	}
}

// InsAttrs .
type InsAttrs struct {
	GlobalAttrs
	ID       string
	Class    string
	Style    string
	Cite     string
	Datetime string
}

// Ins HTML element.
func Ins(attrs *InsAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "ins",
		}
	}
}

// TemplateAttrs .
type TemplateAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Template HTML element.
func Template(attrs *TemplateAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "template",
		}
	}
}

// ArticleAttrs .
type ArticleAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Article HTML element.
func Article(attrs *ArticleAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "article",
		}
	}
}

// FieldsetAttrs .
type FieldsetAttrs struct {
	GlobalAttrs
	ID       string
	Class    string
	Style    string
	Disabled string
	Form     string
	Name     string
}

// Fieldset HTML element.
func Fieldset(attrs *FieldsetAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "fieldset",
		}
	}
}

// H2Attrs .
type H2Attrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// H2 HTML element.
func H2(attrs *H2Attrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "h2",
		}
	}
}

// SupAttrs .
type SupAttrs struct {
	GlobalAttrs
	ID    string
	Class string
	Style string
}

// Sup HTML element.
func Sup(attrs *SupAttrs, children ...Component) Component {
	return func(ctx Context) Element {
		return &el{
			attrs:    attrs,
			children: children,
			ctx:      ctx,
			tag:      "sup",
		}
	}
}
