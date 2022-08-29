package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"github.com/iancoleman/strcase"
)

const fcComponentFileName = "fc/index.tsx"
const fcStyleFileName = "fc/style.ts"

func GenerateReactFc(templateFileDir string, outDir string, name string) error {
	var err error
	f, err := ioutil.ReadFile(templateFileDir + "/" + fcComponentFileName)
	if err != nil {
		fmt.Print(err)
		return err
	}
	result := strings.Replace(string(f), "{name}", getFcComponentName(name), -1)
	pathPrefix := outDir + "/" + getFcFolderName(name)
	err = os.Mkdir(pathPrefix, os.ModePerm)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(
		outDir+"/"+getFcFolderName(name)+"/"+getComponentFileName(name),
		[]byte(result),
		0644,
	)
	if err != nil {
		return err
	}

	f, err = ioutil.ReadFile(templateFileDir + "/" + fcStyleFileName)
	if err != nil {
		fmt.Print(err)
		return err
	}
	err = ioutil.WriteFile(
		outDir + "/" + getFcFolderName(name) + "/" + getCssFileName(name),
		[]byte(string(f)),
		0644,
	)
	return nil
}

func getFcFolderName(name string) string {
	return strcase.ToCamel(name)
}

func getComponentFileName(name string) string {
	return "index.tsx"
}

func getCssFileName(name string) string {
	return "style.ts"
}

func getFcComponentName(name string) string {
	return strcase.ToCamel(name)
}
