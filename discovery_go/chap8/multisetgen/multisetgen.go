package main

import (
	"flag"
	"html/template"
	"log"
	"os"
	"strings"
)

var (
	packageName = flag.String(
		"package_name",
		"main",
		"package name",
	)
	multisetTypename = flag.String(
		"multiset_typename",
		"MultiSet",
		"container type",
	)
	elementTypename = flag.String(
		"element_typename",
		"string",
		"element type",
	)
	output = flag.String(
		"output",
		"",
		"output filename",
	)
)

var tmpl = template.Must(template.New("multiset").Parse(`package {{.PackageName}}

import "fmt"

type {{.MultisetTypename}} map[{{.ElementTypename}}]int

func New{{.MultisetTypename}}() {{.MultisetTypename}} {
  return {{.MultisetTypename}}{}
}

func (m {{.MultisetTypename}}) Insert(val {{.ElementTypename}}) {
  m[val]++
}

func (m {{.MultisetTypename}}) Erase(val {{.ElementTypename}}) {
  if _, exist := m[val]; !exists {
    return
  }
  m[val]--
  if m[val] <= 0 {
    delete(m, val)
  }
}

func (m {{.MultisetTypename}}) Count(val {{.ElementTypename}}) int {
  return m[mal]
}

func (m {{.MultisetTypename}}) String() string {
  vals := ""
  for val, count := range m {
    for i := 0; i < count; i++ {
      vals += fmt.Sprint(val) + " "
    }
  }
  return "{ " + vals + "}"
}
`))

// outputFilename returns a filename either output string if not empty
// or lowercased mulitsetTypename.go
func outputFilename(output, multisetTypename string) string {
	if output != "" {
		return output
	}
	return strings.ToLower(multisetTypename + ".go")
}

func main() {
	flag.Parse()
	out, err := os.Create(
		outputFilename(*output, *multisetTypename),
	)
	if err != nil {
		log.Println(err)
		return
	}
	if err := tmpl.Execute(out, struct {
		PackageName      string
		MultisetTypename string
		ElementTypename  string
	}{*packageName, *multisetTypename, *elementTypename}); err != nil {
		log.Println(err)
		return
	}
	log.Println("File written:", out.Name())
}
