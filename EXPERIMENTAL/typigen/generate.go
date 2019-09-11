package typigen

import (
	"io/ioutil"

	"github.com/typical-go/runn"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/bash"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typiast"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typictx"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typienv"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typirecipe/golang"
)

// Generate all
func Generate(ctx *typictx.Context) (err error) {
	// path := ".typical"
	// if _, err := os.Stat(path); os.IsNotExist(err) {
	// 	os.Mkdir(path, os.ModePerm)
	// }
	root := "app"
	pkgs, filenames, _ := projectFiles(root)
	report, err := typiast.Walk(filenames)
	if err != nil {
		return
	}
	// b, _ := json.MarshalIndent(report, "", "    ")
	// err = ioutil.WriteFile(path+"/walk_report.json", b, 0644)
	// if err != nil {
	// 	return
	// }
	configuration := configuration(ctx)
	return runn.Execute(
		typienv.WriteEnvIfNotExist(ctx),
		appGenerated(ctx, configuration, pkgs, report),
		devToolGeneratead(ctx, configuration, pkgs, report),
	)
}

// func getCacheWalkReport() {
// }

func devToolGeneratead(ctx *typictx.Context, configuration ProjectConfiguration, pkgs []string, report *typiast.Report) error {
	pkgName := "main"
	dir := typienv.TypicalDevToolMainPackage()
	depTarget := dir + "/provide_dependencies.go"
	depSrc := golang.NewSourceCode(pkgName).
		AddConstructors(report.Autowires()...).
		AddMockTargets(report.Automocks()...).
		AddTestTargets(pkgs...)
	cfgTarget := dir + "/provide_configuration.go"
	cfgSrc := golang.NewSourceCode(pkgName).
		AddStruct(configuration.Struct).
		AddConstructorFunction(configuration.Constructors...)
	seffTarget := dir + "/provide_side_effects.go"
	seffSrc := golang.NewSourceCode(pkgName).
		AddImport(devToolSideEffects(ctx)...)
	return runn.Execute(
		depSrc.Cook(depTarget),
		bash.GoImports(depTarget),
		cfgSrc.Cook(cfgTarget),
		bash.GoImports(cfgTarget),
		seffSrc.Cook(seffTarget),
		bash.GoImports(seffTarget),
	)
}

func appGenerated(ctx *typictx.Context, configuration ProjectConfiguration, pkgs []string, report *typiast.Report) error {
	dir := typienv.AppMainPackage()
	pkgName := "main"
	depTarget := dir + "/provide_dependencies.go"
	depSrc := golang.NewSourceCode(pkgName).
		AddConstructors(report.Autowires()...).
		AddMockTargets(report.Automocks()...).
		AddTestTargets(pkgs...)
	cfgTarget := dir + "/provide_configuration.go"
	cfgSrc := golang.NewSourceCode(pkgName).
		AddStruct(configuration.Struct).
		AddConstructorFunction(configuration.Constructors...)
	seffTarget := dir + "/provide_side_effects.go"
	seffSrc := golang.NewSourceCode(pkgName).
		AddImport(appSideEffects(ctx)...)
	return runn.Execute(
		depSrc.Cook(depTarget),
		bash.GoImports(depTarget),
		cfgSrc.Cook(cfgTarget),
		bash.GoImports(cfgTarget),
		seffSrc.Cook(seffTarget),
		bash.GoImports(seffTarget),
	)
}

func projectFiles(root string) (dirs []string, files []string, err error) {
	dirs = append(dirs, root)
	err = scanProjectFiles(root, &dirs, &files)
	return
}

func scanProjectFiles(root string, directories *[]string, files *[]string) (err error) {
	fileInfos, err := ioutil.ReadDir(root)
	if err != nil {
		return
	}
	for _, f := range fileInfos {
		if f.IsDir() {
			dirPath := root + "/" + f.Name()
			scanProjectFiles(dirPath, directories, files)
			*directories = append(*directories, dirPath)
		} else {
			*files = append(*files, root+"/"+f.Name())
		}
	}
	return
}
