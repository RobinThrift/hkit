package hkit

// ColgroupAttrs .
type ColgroupAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Colgroup HTML element.
func Colgroup(attrs *ColgroupAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "colgroup",
	}
}

// NoscriptAttrs .
type NoscriptAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Noscript HTML element.
func Noscript(attrs *NoscriptAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "noscript",
	}
}

// TextareaAttrs .
type TextareaAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Textarea HTML element.
func Textarea(attrs *TextareaAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "textarea",
	}
}

// H6Attrs .
type H6Attrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// H6 HTML element.
func H6(attrs *H6Attrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "h6",
	}
}

// LabelAttrs .
type LabelAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Label HTML element.
func Label(attrs *LabelAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "label",
	}
}

// ObjectAttrs .
type ObjectAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Object HTML element.
func Object(attrs *ObjectAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "object",
	}
}

// DelAttrs .
type DelAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Del HTML element.
func Del(attrs *DelAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "del",
	}
}

// PreAttrs .
type PreAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Pre HTML element.
func Pre(attrs *PreAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "pre",
	}
}

// SubAttrs .
type SubAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Sub HTML element.
func Sub(attrs *SubAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "sub",
	}
}

// VideoAttrs .
type VideoAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Video HTML element.
func Video(attrs *VideoAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "video",
	}
}

// AreaAttrs .
type AreaAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Area HTML element.
func Area(attrs *AreaAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "area",
	}
}

// DialogAttrs .
type DialogAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Dialog HTML element.
func Dialog(attrs *DialogAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "dialog",
	}
}

// HeaderAttrs .
type HeaderAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Header HTML element.
func Header(attrs *HeaderAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "header",
	}
}

// LinkAttrs .
type LinkAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Link HTML element.
func Link(attrs *LinkAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "link",
	}
}

// SectionAttrs .
type SectionAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Section HTML element.
func Section(attrs *SectionAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "section",
	}
}

// SummaryAttrs .
type SummaryAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Summary HTML element.
func Summary(attrs *SummaryAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "summary",
	}
}

// DataAttrs .
type DataAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Data HTML element.
func Data(attrs *DataAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "data",
	}
}

// DlAttrs .
type DlAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Dl HTML element.
func Dl(attrs *DlAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "dl",
	}
}

// NavAttrs .
type NavAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Nav HTML element.
func Nav(attrs *NavAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "nav",
	}
}

// RpAttrs .
type RpAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Rp HTML element.
func Rp(attrs *RpAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "rp",
	}
}

// ScriptAttrs .
type ScriptAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Script HTML element.
func Script(attrs *ScriptAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "script",
	}
}

// BrAttrs .
type BrAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Br HTML element.
func Br(attrs *BrAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "br",
	}
}

// DatalistAttrs .
type DatalistAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Datalist HTML element.
func Datalist(attrs *DatalistAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "datalist",
	}
}

// IframeAttrs .
type IframeAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Iframe HTML element.
func Iframe(attrs *IframeAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "iframe",
	}
}

// OptionAttrs .
type OptionAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Option HTML element.
func Option(attrs *OptionAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "option",
	}
}

// TitleAttrs .
type TitleAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Title HTML element.
func Title(attrs *TitleAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "title",
	}
}

// LiAttrs .
type LiAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Li HTML element.
func Li(attrs *LiAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "li",
	}
}

// ProgressAttrs .
type ProgressAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Progress HTML element.
func Progress(attrs *ProgressAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "progress",
	}
}

// CiteAttrs .
type CiteAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Cite HTML element.
func Cite(attrs *CiteAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "cite",
	}
}

// ColAttrs .
type ColAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Col HTML element.
func Col(attrs *ColAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "col",
	}
}

// DetailsAttrs .
type DetailsAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Details HTML element.
func Details(attrs *DetailsAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "details",
	}
}

// KbdAttrs .
type KbdAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Kbd HTML element.
func Kbd(attrs *KbdAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "kbd",
	}
}

// WbrAttrs .
type WbrAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Wbr HTML element.
func Wbr(attrs *WbrAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "wbr",
	}
}

// BaseAttrs .
type BaseAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Base HTML element.
func Base(attrs *BaseAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "base",
	}
}

// H4Attrs .
type H4Attrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// H4 HTML element.
func H4(attrs *H4Attrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "h4",
	}
}

// MenuAttrs .
type MenuAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Menu HTML element.
func Menu(attrs *MenuAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "menu",
	}
}

// OptgroupAttrs .
type OptgroupAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Optgroup HTML element.
func Optgroup(attrs *OptgroupAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "optgroup",
	}
}

// OlAttrs .
type OlAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Ol HTML element.
func Ol(attrs *OlAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "ol",
	}
}

// RtAttrs .
type RtAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Rt HTML element.
func Rt(attrs *RtAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "rt",
	}
}

// TrackAttrs .
type TrackAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Track HTML element.
func Track(attrs *TrackAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "track",
	}
}

// ParamAttrs .
type ParamAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Param HTML element.
func Param(attrs *ParamAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "param",
	}
}

// SAttrs .
type SAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// S HTML element.
func S(attrs *SAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "s",
	}
}

// H5Attrs .
type H5Attrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// H5 HTML element.
func H5(attrs *H5Attrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "h5",
	}
}

// ThAttrs .
type ThAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Th HTML element.
func Th(attrs *ThAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "th",
	}
}

// AudioAttrs .
type AudioAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Audio HTML element.
func Audio(attrs *AudioAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "audio",
	}
}

// BodyAttrs .
type BodyAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Body HTML element.
func Body(attrs *BodyAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "body",
	}
}

// SelectAttrs .
type SelectAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Select HTML element.
func Select(attrs *SelectAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "select",
	}
}

// FigcaptionAttrs .
type FigcaptionAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Figcaption HTML element.
func Figcaption(attrs *FigcaptionAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "figcaption",
	}
}

// OutputAttrs .
type OutputAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Output HTML element.
func Output(attrs *OutputAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "output",
	}
}

// TdAttrs .
type TdAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Td HTML element.
func Td(attrs *TdAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "td",
	}
}

// InputAttrs .
type InputAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Input HTML element.
func Input(attrs *InputAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "input",
	}
}

// MainAttrs .
type MainAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Main HTML element.
func Main(attrs *MainAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "main",
	}
}

// PictureAttrs .
type PictureAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Picture HTML element.
func Picture(attrs *PictureAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "picture",
	}
}

// SlotAttrs .
type SlotAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Slot HTML element.
func Slot(attrs *SlotAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "slot",
	}
}

// CaptionAttrs .
type CaptionAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Caption HTML element.
func Caption(attrs *CaptionAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "caption",
	}
}

// DfnAttrs .
type DfnAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Dfn HTML element.
func Dfn(attrs *DfnAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "dfn",
	}
}

// H3Attrs .
type H3Attrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// H3 HTML element.
func H3(attrs *H3Attrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "h3",
	}
}

// HeadAttrs .
type HeadAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Head HTML element.
func Head(attrs *HeadAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "head",
	}
}

// H1Attrs .
type H1Attrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// H1 HTML element.
func H1(attrs *H1Attrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "h1",
	}
}

// SpanAttrs .
type SpanAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Span HTML element.
func Span(attrs *SpanAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "span",
	}
}

// VarAttrs .
type VarAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Var HTML element.
func Var(attrs *VarAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "var",
	}
}

// InsAttrs .
type InsAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Ins HTML element.
func Ins(attrs *InsAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "ins",
	}
}

// StyleAttrs .
type StyleAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Style HTML element.
func Style(attrs *StyleAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "style",
	}
}

// TableAttrs .
type TableAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Table HTML element.
func Table(attrs *TableAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "table",
	}
}

// DivAttrs .
type DivAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Div HTML element.
func Div(attrs *DivAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "div",
	}
}

// DtAttrs .
type DtAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Dt HTML element.
func Dt(attrs *DtAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "dt",
	}
}

// EmbedAttrs .
type EmbedAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Embed HTML element.
func Embed(attrs *EmbedAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "embed",
	}
}

// ImgAttrs .
type ImgAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Img HTML element.
func Img(attrs *ImgAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "img",
	}
}

// TfootAttrs .
type TfootAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Tfoot HTML element.
func Tfoot(attrs *TfootAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "tfoot",
	}
}

// AddressAttrs .
type AddressAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Address HTML element.
func Address(attrs *AddressAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "address",
	}
}

// EmAttrs .
type EmAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Em HTML element.
func Em(attrs *EmAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "em",
	}
}

// MetaAttrs .
type MetaAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Meta HTML element.
func Meta(attrs *MetaAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "meta",
	}
}

// PAttrs .
type PAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// P HTML element.
func P(attrs *PAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "p",
	}
}

// SupAttrs .
type SupAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Sup HTML element.
func Sup(attrs *SupAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "sup",
	}
}

// UlAttrs .
type UlAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Ul HTML element.
func Ul(attrs *UlAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "ul",
	}
}

// SmallAttrs .
type SmallAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Small HTML element.
func Small(attrs *SmallAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "small",
	}
}

// StrongAttrs .
type StrongAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Strong HTML element.
func Strong(attrs *StrongAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "strong",
	}
}

// TbodyAttrs .
type TbodyAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Tbody HTML element.
func Tbody(attrs *TbodyAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "tbody",
	}
}

// MeterAttrs .
type MeterAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Meter HTML element.
func Meter(attrs *MeterAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "meter",
	}
}

// SourceAttrs .
type SourceAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Source HTML element.
func Source(attrs *SourceAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "source",
	}
}

// BdoAttrs .
type BdoAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Bdo HTML element.
func Bdo(attrs *BdoAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "bdo",
	}
}

// HgroupAttrs .
type HgroupAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Hgroup HTML element.
func Hgroup(attrs *HgroupAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "hgroup",
	}
}

// LegendAttrs .
type LegendAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Legend HTML element.
func Legend(attrs *LegendAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "legend",
	}
}

// QAttrs .
type QAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Q HTML element.
func Q(attrs *QAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "q",
	}
}

// AsideAttrs .
type AsideAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Aside HTML element.
func Aside(attrs *AsideAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "aside",
	}
}

// FooterAttrs .
type FooterAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Footer HTML element.
func Footer(attrs *FooterAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "footer",
	}
}

// FormAttrs .
type FormAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Form HTML element.
func Form(attrs *FormAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "form",
	}
}

// H2Attrs .
type H2Attrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// H2 HTML element.
func H2(attrs *H2Attrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "h2",
	}
}

// BAttrs .
type BAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// B HTML element.
func B(attrs *BAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "b",
	}
}

// CodeAttrs .
type CodeAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Code HTML element.
func Code(attrs *CodeAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "code",
	}
}

// IAttrs .
type IAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// I HTML element.
func I(attrs *IAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "i",
	}
}

// SampAttrs .
type SampAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Samp HTML element.
func Samp(attrs *SampAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "samp",
	}
}

// TheadAttrs .
type TheadAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Thead HTML element.
func Thead(attrs *TheadAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "thead",
	}
}

// UAttrs .
type UAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// U HTML element.
func U(attrs *UAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "u",
	}
}

// AAttrs .
type AAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
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
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "a",
	}
}

// ArticleAttrs .
type ArticleAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Article HTML element.
func Article(attrs *ArticleAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "article",
	}
}

// CanvasAttrs .
type CanvasAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Canvas HTML element.
func Canvas(attrs *CanvasAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "canvas",
	}
}

// MarkAttrs .
type MarkAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Mark HTML element.
func Mark(attrs *MarkAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "mark",
	}
}

// FigureAttrs .
type FigureAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Figure HTML element.
func Figure(attrs *FigureAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "figure",
	}
}

// HrAttrs .
type HrAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Hr HTML element.
func Hr(attrs *HrAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "hr",
	}
}

// TrAttrs .
type TrAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Tr HTML element.
func Tr(attrs *TrAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "tr",
	}
}

// TimeAttrs .
type TimeAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Time HTML element.
func Time(attrs *TimeAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "time",
	}
}

// BdiAttrs .
type BdiAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Bdi HTML element.
func Bdi(attrs *BdiAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "bdi",
	}
}

// ButtonAttrs .
type ButtonAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Button HTML element.
func Button(attrs *ButtonAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "button",
	}
}

// MapAttrs .
type MapAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Map HTML element.
func Map(attrs *MapAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "map",
	}
}

// TemplateAttrs .
type TemplateAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Template HTML element.
func Template(attrs *TemplateAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "template",
	}
}

// DdAttrs .
type DdAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Dd HTML element.
func Dd(attrs *DdAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "dd",
	}
}

// HtmlAttrs .
type HtmlAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Html HTML element.
func Html(attrs *HtmlAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "html",
	}
}

// RubyAttrs .
type RubyAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Ruby HTML element.
func Ruby(attrs *RubyAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "ruby",
	}
}

// AbbrAttrs .
type AbbrAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Abbr HTML element.
func Abbr(attrs *AbbrAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "abbr",
	}
}

// BlockquoteAttrs .
type BlockquoteAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Blockquote HTML element.
func Blockquote(attrs *BlockquoteAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "blockquote",
	}
}

// FieldsetAttrs .
type FieldsetAttrs struct {
	GlobalAttrs
	ID             string
	Class          string
	Style          string
	Globals        string
	Href           string
	Target         string
	Download       string
	Ping           string
	Rel            string
	Hreflang       string
	Type           string
	Referrerpolicy string
}

// Fieldset HTML element.
func Fieldset(attrs *FieldsetAttrs, children ...Component) Component {
	return &comp{
		attrs:    attrs,
		children: children,
		tag:      "fieldset",
	}
}
