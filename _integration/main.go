package main

import (
	"fmt"
	"os"

	"github.com/ovrclk/akash/_integration/cmp"
	"github.com/ovrclk/gestalt"
	g "github.com/ovrclk/gestalt/builder"
	"github.com/ovrclk/gestalt/vars"
)

func main() {

	if len(os.Args) < 2 {
		usage()
	}

	var (
		suite    gestalt.Component
		defaults vars.Meta
	)

	switch os.Args[1] {
	case "standalone":
		suite, defaults = prepareStandalone()
	case "kube":
		suite, defaults = prepareKube()
	default:
		usage()
	}

	gestalt.RunWith(suite.WithMeta(defaults), os.Args[2:])
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %v <standalone|kube> [gestalt-options ...]\n")
	os.Exit(1)
}

func prepareStandalone() (gestalt.Component, vars.Meta) {
	defaults := detectDefaults().
		Default("provider-root", "./data/provider").
		Default("akashd-path", "../akashd").
		Default("akashd-root", "./data/node").
		Default("deployment-path", "./deployment.yml").
		Default("provider-path", "./provider.yml").
		Default("akash-root", "./data/client")

	return cmp.StandaloneSuite(), defaults

}

func prepareKube() (gestalt.Component, vars.Meta) {
	defaults := detectDefaults().
		Require("host-base").
		Default("node-hostname", "node-0").
		Default("run-dir", "../_run/multi").
		Default("akash-root", "../_run/multi/data/client")

	return cmp.KubeSuite(), defaults
}

func detectDefaults() vars.Meta {
	return g.
		Default("akash-path", "../akash")
}
