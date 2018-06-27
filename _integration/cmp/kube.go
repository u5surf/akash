package cmp

import (
	"os"

	"github.com/ovrclk/gestalt"
	g "github.com/ovrclk/gestalt/builder"
)

func kubeCheckNodes() gestalt.Component {
	return g.Group("kube-check-nodes").
		Run(
			g.Retry(30).Run(
				g.SH("check", "make", "helm-check-nodes").
					Dir("{{run-dir}}").
					AddEnv("DOMAIN", "{{host-base}}").
					AddEnv("PATH", os.Getenv("PATH")).
					AddEnv("HOME", os.Getenv("HOME")))).
		WithMeta(g.Require("run-dir", "host-base"))
}

func kubeCheckProviders() gestalt.Component {
	return g.Group("kube-check-providers").
		Run(
			g.Retry(30).Run(
				g.SH("check", "make", "helm-check-providers").
					Dir("{{run-dir}}").
					AddEnv("DOMAIN", "{{host-base}}").
					AddEnv("PATH", os.Getenv("PATH")).
					AddEnv("HOME", os.Getenv("HOME")))).
		WithMeta(g.Require("run-dir", "host-base"))
}

func kubeInstall() gestalt.Component {
	return g.Group("install-components").
		Run(kubeCheckNodes()).
		Run(kubeCheckProviders())
}
