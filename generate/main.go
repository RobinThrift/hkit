package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dave/jennifer/jen"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var enCaser = cases.Title(language.English)

func main() {
	elements, err := getElements()
	if err != nil {
		log.Fatal(err)
	}

	f := jen.NewFile("hkit") // nolint: varnamelen

	for _, el := range elements { // nolint: varnamelen
		propsStructName := enCaser.String(el.name) + "Props"

		propsFields := make([]jen.Code, len(el.attrs)+5)

		propsFields[0] = jen.Id("*GenericProps")
		propsFields[1] = jen.Id("X").Id("Attributes")
		propsFields[2] = jen.Id("ID").String()
		propsFields[3] = jen.Id("Class").String()
		propsFields[4] = jen.Id("Style").Map(jen.String()).String()
		for i, a := range el.attrs {
			propsFields[i+5] = jen.Id(a).String()
		}

		f.Type().Id(propsStructName).Struct(
			propsFields...,
		)

		writeStatements := make([]jen.Code, 0, len(el.attrs)+6)

		writeStatements = append(writeStatements, jen.If(jen.Id("p").Op("==").Nil()).Block(jen.Return(jen.Nil())))

		writeStatements = append(writeStatements, jen.If(jen.Id("p").Dot("GenericProps").Op("!=").Nil()).Block(
			jen.Id("err").Op(":=").Id("p").Dot("GenericProps").Dot("writeTo").Call(jen.Id("w")),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Return(jen.Id("err")),
			),
		))

		writeStatements = append(writeStatements, jen.If(jen.Id("p").Dot("Style").Op("!=").Nil()).Block(
			jen.Id("err").Op(":=").Id("writeStyle").Call(jen.Id("p").Dot("Style"), jen.Id("w")),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Return(jen.Id("err")),
			),
		))

		writeStatements = append(writeStatements, jen.If(jen.Id("p").Dot("X").Op("!=").Nil()).Block(
			jen.Id("err").Op(":=").Id("p").Dot("X").Dot("writeTo").Call(jen.Id("w")),
			jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Return(jen.Id("err")),
			),
		))

		for _, p := range append(el.attrs, "ID", "Class") {
			writeStatements = append(writeStatements, jen.If(jen.Id("p").Dot(p).Op("!=").Lit("")).Block(
				jen.Id("err").Op(":=").Id("writeAttr").Call(jen.Lit(strings.ToLower(p)), jen.Id("p").Dot(p), jen.Id("w")),
				jen.If(jen.Id("err").Op("!=").Nil()).Block(
					jen.Return(jen.Id("err")),
				),
			))
		}

		writeStatements = append(writeStatements, jen.Return(jen.Nil()))

		f.Func().Params(jen.Id("p").Op("*").Id(propsStructName)).Id("writeTo").Params(jen.Id("w").Qual("io", "Writer")).Id("error").Block(
			writeStatements...,
		)

		f.Line()

		f.Func().Id(enCaser.String(el.name)).Params(jen.Id("props").Op("*").Id(propsStructName), jen.Id("children").Op("...").Id("Component")).Id("Component").Block(
			jen.Return(jen.Id("&tag").Types(jen.Op("*").Id(propsStructName)).Values(jen.Dict{
				jen.Id("name"):        jen.Lit(el.name),
				jen.Id("props"):       jen.Id("props"),
				jen.Id("children"):    jen.Id("children"),
				jen.Id("selfClosing"): jen.Lit(el.isSelfClosing),
			})),
		)

		f.Line()
	}

	err = f.Save("builtins_gen.go")
	if err != nil {
		log.Fatal(err)
	}
}

type element struct {
	name          string
	attrs         []string
	isSelfClosing bool
}

func getElements() ([]element, error) { // nolint: gocognit
	elements := make([]element, 0, 110)

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

			currEl := element{
				name:  name,
				attrs: []string{},
			}

			tr.Find("td").EachWithBreak(func(j int, td *goquery.Selection) bool {
				if j == 3 {
					if td.Text() == "empty" {
						currEl.isSelfClosing = true
					}
				}
				if j == 4 {
					for _, attr := range strings.Split(strings.TrimSpace(td.Text()), ";") {
						if attr == "globals" {
							continue
						}
						currEl.attrs = append(currEl.attrs, toPascalCase(attr))
					}
					return false
				}

				return true
			})

			if name == "h1, h2, h3, h4, h5, h6" {
				for _, n := range strings.Split(name, ", ") {
					headerEl := element{name: n, attrs: currEl.attrs}
					elements = append(elements, headerEl)
				}
				return
			}

			elements = append(elements, currEl)
		})
	})

	return elements, nil
}

// from https://github.com/iancoleman/strcase/blob/master/camel.go
func toPascalCase(s string) string { // nolint: gocognit
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
