package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

var (
	fPackageName     = flag.String("package", "", "package name (required)")
	fStructName      = flag.String("struct", "", "struct name (required)")
	fTmplPath        = flag.String("tmpl", "elementopts.go", "template path")
	fTmplPackageName = flag.String("tmpl-package", "templui", "template package name")
	fTmplStructName  = flag.String("tmpl-struct", "Element", "template struct name")
	fOutput          = flag.String("output", "*_generated.go", "generated file")
)

func template(
	packageName string,
	structName string,
	tmplPackage string,
	tmplStruct string,
	tmpl string,
) string {
	newTmpl := strings.ReplaceAll(
		tmpl, tmplStruct, structName)
	newTmpl = strings.ReplaceAll(
		newTmpl,
		fmt.Sprintf("package %s\n", tmplPackage),
		"",
	)

	return fmt.Sprintf(
		"package %s\n\n// NOTE: Don't change this file. It is auto generated! You will lose any change\n%s",
		packageName,
		newTmpl,
	)
}

func getTemplateData(tmplPath string, maxParentRecursion int) (string, error) {
	for i := 0; i < maxParentRecursion; i++ {
		if _, err := os.Stat(tmplPath); errors.Is(err, os.ErrNotExist) {
			tmplPath = path.Join("..", tmplPath)
			continue
		}

		// at this point, we know the file exists
		file, err := os.ReadFile(tmplPath)
		if err != nil {
			return "", fmt.Errorf("read file \"%s\": %s", tmplPath, err.Error())
		}

		return string(file), nil
	}

	return "", fmt.Errorf("template \"%s\" not found", tmplPath)
}

func main() {
	flag.Parse()

	packageName := *fPackageName
	if len(packageName) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	structName := *fStructName
	if len(structName) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	// no point in going further than 5 folders up, if that happens,
	// the file doesn't exist
	tmpl, err := getTemplateData(*fTmplPath, 5)
	if err != nil {
		flag.Usage()
		panic(err)
	}

	outputFileName := strings.ReplaceAll(*fOutput, "*", packageName)
	fileData := template(
		packageName,
		structName,
		*fTmplPackageName,
		*fTmplStructName,
		tmpl,
	)

	err = os.WriteFile(outputFileName, []byte(fileData), 0644)
	if err != nil {
		panic(err)
	}
}
