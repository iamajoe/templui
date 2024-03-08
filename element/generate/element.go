package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"regexp"
	"strings"
)

var (
	generatedExt      = "_generated.go"
	tmplLocation      = path.Join("element", "elementopts.go")
	tmplStructReplace = "*Element"
	commonIgnorePaths = []string{".git", ".github", "vendor", "node_modules"}
	ignoreExt         = []string{"_test.go", "_templ.go", generatedExt}
	structNameRegex   = regexp.MustCompile(`(?m).*type (.*) struct {\n.*element.Element`)
)

func getIgnorePaths(root string) []string {
	paths := commonIgnorePaths

	gitIgnorePath := path.Join(root, ".gitignore")
	if _, err := os.Stat(gitIgnorePath); errors.Is(err, os.ErrNotExist) {
		return paths
	}

	// we just ignore the error if anything, gitignore might have issues
	// no point in failing because of that
	file, err := os.ReadFile(gitIgnorePath)
	if err != nil {
		fmt.Println(fmt.Errorf("read file .gitignore: %s", err.Error()))
		return paths
	}

	gitIgnoreFiles := strings.Split(string(file), "\n")
	paths = append(paths, gitIgnoreFiles...)

	return paths
}

func template(filePath string, structName string, tmpl string) error {
	packageName := strings.ReplaceAll(path.Base(filePath), ".go", "")
	newFilePath := strings.Replace(filePath, ".go", generatedExt, 1)

	newTmpl := strings.ReplaceAll(tmpl, tmplStructReplace, fmt.Sprintf("*%s", structName))
	newTmpl = strings.ReplaceAll(newTmpl, "package element\n", "")
	newTmpl = fmt.Sprintf(
		"package %s\n\n// NOTE: Don't change this file. It is auto generated! You will lose any change\n",
		packageName,
	) + newTmpl

	return os.WriteFile(newFilePath, []byte(newTmpl), 0644)
}

func walkAndGenerate(root string) error {
	fileSystem := os.DirFS(root)
	ignorePaths := getIgnorePaths(root)

	// read the template to be used further down the line
	tmplFileRaw, err := os.ReadFile(path.Join(root, tmplLocation))
	if err != nil {
		return err
	}
	tmplFile := string(tmplFileRaw)

	return fs.WalkDir(fileSystem, ".", func(src string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		name := d.Name()

		// ignore any file starting with an _, it shouldnt generate anything
		if name[0] == '_' {
			if d.IsDir() {
				return fs.SkipDir
			}

			return nil
		}

		// check if this should be ignored
		for _, ignore := range ignorePaths {
			if name == ignore {
				return fs.SkipDir
			}
		}

		// directory will be recursed
		if d.IsDir() {
			return nil
		}

		// we only want .go files. those are the ones that will have the
		// struct to inherit Element
		if !strings.Contains(name, ".go") {
			return nil
		}

		// we want to ignore
		for _, ignore := range ignoreExt {
			if i := strings.Index(name, ignore); i != -1 && i == len(name)-len(ignore) {
				return nil
			}
		}

		srcPath := path.Join(root, src)
		fileRaw, err := os.ReadFile(srcPath)
		if err != nil {
			return err
		}

		file := string(fileRaw)
		match := structNameRegex.FindStringSubmatch(file)
		if len(match) < 2 {
			return nil
		}

		// generate the template
		return template(srcPath, match[1], tmplFile)
	})
}

func main() {
	fmt.Println("generating elements...")

	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = walkAndGenerate(path.Join(root, ".."))
	if err != nil {
		panic(err)
	}
}
