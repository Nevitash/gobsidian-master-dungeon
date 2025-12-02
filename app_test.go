package main

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

const binaryName = "gmd"

func TestMain(m *testing.M) {
	testscript.Main(m, map[string]func(){
		"root": main,
	})
}

func TestScripts(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata_scripts",
	})
}
