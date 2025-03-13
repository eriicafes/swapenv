package main

import (
	"embed"
	"encoding/json"

	"github.com/eriicafes/swapenv/cmd"
)

//go:embed package.json
var content embed.FS

func main() {
	// get version from package.json
	b, err := content.Open("package.json")
	if err != nil {
		panic(err)
	}
	var packageJson struct{ Version string }
	err = json.NewDecoder(b).Decode(&packageJson)
	if err != nil {
		panic(err)
	}

	cmd.Execute(packageJson.Version)
}
