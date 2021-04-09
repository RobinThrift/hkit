package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dave/jennifer/jen"
)

func main() {
	elements, err := getElements()
	if err != nil {
		log.Fatal(err)
	}

	f := jen.NewFile("hkit")

	for el, attrs := range elements {
		f.Comment(strings.Title(el) + "Attrs .")

		attrFields := make([]jen.Code, len(attrs)+4)
		attrFields[0] = jen.Id("GlobalAttrs")
		attrFields[1] = jen.Id("ID").String()
		attrFields[2] = jen.Id("Class").String()
		attrFields[3] = jen.Id("Style").String()
		for i, a := range attrs {
			attrFields[i+4] = jen.Id(a).String()
		}

		f.Type().Id(strings.Title(el) + "Attrs").Struct(
			attrFields...,
		)

		f.Line()

		f.Comment(strings.Title(el) + " HTML element.")
		f.Func().Id(strings.Title(el)).Params(jen.Id("attrs").Id("*"+strings.Title(el)+"Attrs"), jen.Id("children").Op("...").Id("Component")).Id("Component").Block(
			jen.Return(jen.Func().Params(
				jen.Id("ctx").Id("Context"),
			).Id("Element").Block(
				jen.Return(jen.Id("&el").Values(jen.Dict{
					jen.Id("ctx"):      jen.Id("ctx"),
					jen.Id("tag"):      jen.Lit(el),
					jen.Id("attrs"):    jen.Id("attrs"),
					jen.Id("children"): jen.Id("children"),
				},
				)),
			)),
		)

		f.Line()
	}

	err = f.Save("builtins.go")
	if err != nil {
		log.Fatal(err)
	}
}

func getElements() (map[string][]string, error) {
	elements := map[string][]string{
		"html": {},
	}

	res, err := http.Get("https://html.spec.whatwg.org/multipage/indices.html")
	if err != nil {
		return elements, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	defer res.Body.Close()
	if err != nil {
		return elements, err
	}

	doc.Find("#elements-3").Each(func(_ int, s *goquery.Selection) {
		el := s.NextUntil("table").Next()

		el.Find("tr").Each(func(i int, tr *goquery.Selection) {
			if i == 0 {
				return
			}

			name := strings.TrimSpace(tr.Find("th").First().Text())
			if name == "" || name == "autonomous custom elements" || strings.HasPrefix(name, "SVG") || strings.HasPrefix(name, "Math") {
				return
			}

			attrs := []string{}

			tr.Find("td").EachWithBreak(func(j int, td *goquery.Selection) bool {
				if j == 4 {
					for _, attr := range strings.Split(strings.TrimSpace(td.Text()), ";") {
						if attr == "globals" {
							continue
						}
						attrs = append(attrs, toPascalCase(attr))
					}
					return false
				}

				return true
			})

			if name == "h1, h2, h3, h4, h5, h6" {
				for _, n := range strings.Split(name, ", ") {
					elements[n] = attrs
				}
				return
			}

			elements[name] = attrs
		})
	})

	return elements, nil
}

// from https://github.com/iancoleman/strcase/blob/master/camel.go
func toPascalCase(s string) string {
	if s == "" {
		return s
	}

	n := strings.Builder{}
	n.Grow(len(s))

	capNext := true
	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if capNext {
			if vIsLow {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsCap {
				v += 'a'
				v -= 'A'
			}
		}
		if vIsCap || vIsLow {
			n.WriteByte(v)
			capNext = false
		} else if vIsNum := v >= '0' && v <= '9'; vIsNum {
			n.WriteByte(v)
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}

	return n.String()
}
