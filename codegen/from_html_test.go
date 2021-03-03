package codegen_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"testing"
	"text/template"

	"github.com/PuerkitoBio/goquery"
	"github.com/b3ntly/salesforce/codegen"
	"github.com/stretchr/testify/require"
)

// takes about 100s
func TestAllHTML2Go(t *testing.T) {
	t.Skip() 
	fileInfos, err := ioutil.ReadDir("../chromedp/all")
	require.Nil(t, err)

	var entities []*codegen.Struct

	// parse all .html files into []*codegen.Struct
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			contents, err := ioutil.ReadFile("../chromedp/all/" + fileInfo.Name())
			require.Nil(t, err)
			doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(contents))
			require.Nil(t, err)
			entity, err := codegen.FromHTML(doc)
			if err == nil {
				entities = append(entities, entity)
			}

			if errors.Is(err, codegen.ErrEmptyDocumentation) {
				continue
			}

			if err != nil {
				t.Log(err.Error(), fileInfo.Name())
				t.Fail()
				continue
			}
		}
	}

	var entityNames []string 
	entityMap := map[string]interface{}{}
	for i := 0; i < len(entities); i++ {
		entityNames = append(entityNames, entities[i].Name)
		entityMap[entities[i].Name] = entities[i]
	}

	seed := &codegen.Seed{
		OutputDirectory: "autogenerated",
		PackageName: "objects",
		DistinctSubpackages: entityNames,
		Options: []codegen.Option{
			codegen.WithTemplateMap(map[string]*template.Template{
				"object.go":template.Must(template.New("").Funcs(codegen.DefaultFuncMap).Parse(`
				package {{.SubPackageName}}
				
				{{.Data.String}}
				`)),
			}),
			codegen.WithDistinctData(entityMap),
		},
	}

	if err := codegen.Generate(seed); err != nil {
		t.Log(err.Error())
		t.Fail() 
	}
}
