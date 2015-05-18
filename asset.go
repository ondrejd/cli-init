package main

import (
	"fmt"
	"io/ioutil"
)

// This file is here because of this:
// https://github.com/tcnksm/cli-init/issues/4

const assets_path = "/home/ondrejd/GoWorkspace/src/github.com/ondrejd/cli-init/templates/"

// bindata_read reads the given file from disk. It returns
// an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

// assets_commands_tmpl reads file data from disk.
func assets_commands_tmpl() ([]byte, error) {
	return bindata_read(
		fmt.Sprintf("%scommands.tmpl", assets_path),
		"templates/commands.tmpl",
	)
}

// assets_changelog_tmpl reads file data from disk.
func assets_changelog_tmpl() ([]byte, error) {
	return bindata_read(
		fmt.Sprintf("%sCHANGELOG.tmpl", assets_path),
		"templates/CHANGELOG.tmpl",
	)
}

// assets_main_tmpl reads file data from disk.
func assets_main_tmpl() ([]byte, error) {
	return bindata_read(
		fmt.Sprintf("%smain.tmpl", assets_path),
		"templates/main.tmpl",
	)
}

// assets_readme_tmpl reads file data from disk.
func assets_readme_tmpl() ([]byte, error) {
	return bindata_read(
		fmt.Sprintf("%sREADME.tmpl", assets_path),
		"templates/README.tmpl",
	)
}

// assets_version_tmpl reads file data from disk.
func assets_version_tmpl() ([]byte, error) {
	return bindata_read(
		fmt.Sprintf("%sversion.tmpl", assets_path),
		"templates/version.tmpl",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	if f, ok := _bindata[name]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"templates/commands.tmpl": assets_commands_tmpl,
	"templates/CHANGELOG.tmpl": assets_changelog_tmpl,
	"templates/main.tmpl": assets_main_tmpl,
	"templates/README.tmpl": assets_readme_tmpl,
	"templates/version.tmpl": assets_version_tmpl,
}
