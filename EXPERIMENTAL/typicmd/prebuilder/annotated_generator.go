package prebuilder

import (
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/bash"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typicmd/prebuilder/golang"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typicmd/prebuilder/walker"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typictx"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typienv"
)

// AnnotatedGenerator responsible for generate annotated
type AnnotatedGenerator struct {
	*typictx.Context
	*walker.ProjectFiles
	Packages []string
}

// Generate the file
func (g *AnnotatedGenerator) Generate() (err error) {
	if g.check() {
		return g.generate()
	}
	return
}

func (g *AnnotatedGenerator) generate() (err error) {
	defer elapsed("Generate Annotated")()
	pkg := typienv.Dependency.Package
	src := golang.NewSourceCode(pkg)
	for _, pkg := range g.Packages {
		src.AddImport("", g.Context.Root+"/"+pkg)
	}
	src.AddConstructors(g.ProjectFiles.Autowires()...)
	src.AddMockTargets(g.ProjectFiles.Automocks()...)
	target := dependency + "/annotateds.go"
	err = src.Cook(target)
	if err != nil {
		return
	}
	return bash.GoImports(target)
}

func (g *AnnotatedGenerator) check() bool {
	return true
}