package codegen_test

import (
	"testing"
	"text/template"

	"github.com/beeekind/go-salesforce-sdk/codegen"
	"github.com/beeekind/go-salesforce-sdk/metadata"
	"github.com/stretchr/testify/require"
)

var fromDescriptionTests = map[*metadata.Describe][]codegen.Struct{
	{
		Name: "foo",
		Fields: []*metadata.Field{
			{Name: "Title", Type: "string", ReferenceTo: nil, RelationshipName: ""},
			{Name: "OwnerId", Type: "reference", ReferenceTo: []string{"Group", "User"}, RelationshipName: "Owner"},
			{Name: "CreatedById", Type: "reference", ReferenceTo: []string{"User"}, RelationshipName: "CreatedBy"},
		},
		ChildRelationships: []*metadata.ChildRelationship{
			{RelationshipName: "Events", JunctionReferenceTo: nil, ChildSObject: "Event"},
		},
	}: {
		{Name: "Foo", Dependencies: []string{"Group", "User", "Event"}, Properties: []*codegen.Property{
			{ParentName: "Foo", Name: "CreatedBy", Type: "*User", Tag: codegen.Tag{"json": []string{"CreatedBy"}}},
			{ParentName: "", Name: "CreatedByID", Type: "string", Tag: codegen.Tag{"json": []string{"CreatedById"}}},
			{ParentName: "Foo", Name: "Events", Type: "struct {\n\tDone bool `json:\"done\"`\n\tCount int `json:\"count\"`\n\tTotalSize int `json:\"totalSize\"`\n\tRecords []*Event `json:\"records\"`\n}", Tag: codegen.Tag{"json": []string{"Events"}}},
			{ParentName: "Foo", Name: "Owner", Type: "*GroupUser", Tag: codegen.Tag{"json": []string{"Owner"}}},
			{ParentName: "", Name: "OwnerID", Type: "string", Tag: codegen.Tag{"json": []string{"OwnerId"}}},
			{ParentName: "", Name: "Title", Type: "string", Tag: codegen.Tag{"json": []string{"Title"}}},
		}},
		{Name: "GroupUser", Dependencies: []string{"Group", "User"}, Properties: []*codegen.Property{
			{Name: "Group", IsEmbedded: true},
			{Name: "User", IsEmbedded: true},
		}},
	},
}

func TestFromDescription(t *testing.T) {
	for desc, expected := range fromDescriptionTests {
		t.Run(desc.Name, func(t *testing.T) {
			result, err := codegen.FromDescribe(desc, false)
			require.Nil(t, err)
			require.Equal(t, len(expected), len(result))

			result = result.Sort() 

			for i := 0; i < len(expected); i++ {
				s1 := expected[i]
				s2 := result[i]

				require.Equal(t, s1.Name, s2.Name)

				require.Equal(t, len(s1.Properties), len(s2.Properties))

				for j := 0; j < len(s1.Properties); j++ {
					p1 := s1.Properties[j]
					p2 := s2.Properties[j]

					require.Equal(t, p1, p2)
				}
			}
		})
	}
}

func TestFromDescriptionSeed(t *testing.T) {
	for desc := range fromDescriptionTests {
		t.Run(desc.Name, func(t *testing.T) {
			result, err := codegen.FromDescribe(desc, false)
			require.Nil(t, err)

			var subpackages []string
			entityMap := map[string]interface{}{}
			for _, entity := range result {
				uncapturedVariable := entity
				subpackages = append(subpackages, entity.Name)
				entityMap[entity.Name] = &uncapturedVariable
			}

			seed := &codegen.Seed{
				OutputDirectory:     "autogenerated",
				PackageName:         "objects",
				DistinctSubpackages: subpackages,
				Options: []codegen.Option{
					codegen.WithTemplateMap(map[string]*template.Template{
						"object.go": template.Must(template.New("").Funcs(codegen.DefaultFuncMap).Parse(`
						package {{.SubPackageName}}
						
						{{.Data.String}}
						`)),
					}),
					codegen.WithDistinctData(entityMap),
				},
			}

			if _, err := codegen.Render(seed); err != nil {
				t.Log(err.Error())
				t.FailNow() 
			}

		})
	}
}
