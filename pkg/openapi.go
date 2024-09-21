package openapi

import (
	"fmt"
	"os"

	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/renderer"
)

func RenderExample() {
	// create a new JSON mock generator
	mg := renderer.NewMockGenerator(renderer.JSON)

	// tell the mock generator to pretty print the output
	mg.SetPretty()

	exampleOpenapiFp, err := os.ReadFile("./test/data/fastapi-openapi.json")
	if err != nil {
		panic(err)
	}

	// create a new document from specification and build a v3 model.
	document, err := libopenapi.NewDocument(exampleOpenapiFp)
	if err != nil {
		panic(err)
	}
	v3Model, errs := document.BuildV3Model()
	if errs != nil {
		panic(errs)
	}
	dummyRequestModel := v3Model.Model.Components.Schemas["DummyRequest"]
	dummyRequest := dummyRequestModel.Schema()
	fakeDummyRequest, err := mg.GenerateMock(dummyRequest, "")
	if err != nil {
		panic(err)
	}

	// print the mock to stdout
	fmt.Println(string(fakeDummyRequest))
}
