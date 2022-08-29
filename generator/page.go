package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"github.com/iancoleman/strcase"
)

func GeneratePage(templateFileDir string, outDir string, name string) error {
	var err error
	pathPrefix := outDir + "/" + getPageFolderName(name)
	err = os.MkdirAll(pathPrefix + "/pc", os.ModePerm)
	err = os.MkdirAll(pathPrefix + "/mobile", os.ModePerm)
	err = os.MkdirAll(pathPrefix + "/common/assets", os.ModePerm)
	err = os.MkdirAll(pathPrefix + "/common/components", os.ModePerm)
	err = os.MkdirAll(pathPrefix + "/common/constants", os.ModePerm)

	generateCommonFiles(templateFileDir, outDir, name)
	generatePlatformFiles(templateFileDir, outDir, name, "pc")
	generatePlatformFiles(templateFileDir, outDir, name, "mobile")
	generatePageIndex(templateFileDir, outDir, name)

	if err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}

func generatePageIndex(templateFileDir string, outDir string, name string) error {
	const pageIndexFileName = "page/index.ts"
	var err error
	f, err := ioutil.ReadFile(templateFileDir + "/" + pageIndexFileName)

	result := strings.Replace(string(f), "{name}", strcase.ToCamel(name), -1)
	err = ioutil.WriteFile(
		outDir + "/" + getPageFolderName(name) + "/" + "index.ts",
		[]byte(result),
		0644,
	)

	if err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}

func generatePlatformFiles(templateFileDir string, outDir string, name string, platform string) error {
	var err error
	camelName := strcase.ToCamel(name)
	currentTmplDir := templateFileDir + "/page" + "/" + platform + "/"
	currentOutDir := outDir + "/" + getPageFolderName(name) + "/" + platform + "/"

	f, err := ioutil.ReadFile(currentTmplDir + "index.tsx")
	result := strings.Replace(string(f), "{name}", camelName, -1)
	err = ioutil.WriteFile(currentOutDir + "index.tsx", []byte(result), 0644)

	f, err = ioutil.ReadFile(currentTmplDir +  "Component.tsx")
	result = strings.Replace(string(f), "{name}", camelName, -1)
	err = ioutil.WriteFile(currentOutDir + camelName + ".tsx", []byte(result), 0644)

	f, err = ioutil.ReadFile(currentTmplDir + "style.ts")
	err = ioutil.WriteFile(currentOutDir + "style.ts", []byte(string(f)), 0644)

	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

func generateCommonFiles(templateFileDir string, outDir string, name string) error {
	var err error
	camelName := strcase.ToCamel(name)
	currentTmplDir := templateFileDir + "/page" + "/common" + "/"
	currentOutDir := outDir + "/" + getPageFolderName(name) + "/" + "common" + "/"
	
	f, err := ioutil.ReadFile(currentTmplDir + "constants" + "/" +  "index.ts")
	err = ioutil.WriteFile(currentOutDir + "constants" + "/" + "index.ts", []byte(string(f)), 0644)

	f, err = ioutil.ReadFile(currentTmplDir + "constants" + "/" +  "tdk.ts")
	result := strings.Replace(string(f), "{name}", camelName, -1)
	err = ioutil.WriteFile(currentOutDir + "constants" + "/" + "tdk.ts", []byte(result), 0644)

	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

func getPageFolderName(name string) string {
	return strcase.ToKebab(name)
}
