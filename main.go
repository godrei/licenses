package main

import (
	"fmt"

	"github.com/bitrise-io/steps-xcode-test/pretty"
	"github.com/godrei/licenses/licenses"
)

func main() {
	licensesByPackage, err := licenses.Licenses("")
	if err != nil {
		panic(err)
	}
	fmt.Println(pretty.Object(licensesByPackage))
}
