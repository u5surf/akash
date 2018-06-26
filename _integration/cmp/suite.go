package cmp

import (
	"github.com/ovrclk/gestalt"
	g "github.com/ovrclk/gestalt/builder"
)

func StandaloneSuite() gestalt.Component {
	key := newKey("master")
	paddr := g.Ref("provider-id")
	daddr := g.Ref("deployment-id")
	return g.Suite("main").
		Run(groupKey(key)).
		Run(groupNodeRun(key)).
		Run(groupAccountSend(key)).
		Run(groupProvider(paddr)).
		Run(groupDeploy(key.name, daddr))
}

func kubeInit() gestalt.Component {
	return g.EXEC("init", "./run.sh", "init").
		Dir("{{run-dir}}").
		WithMeta(g.Require("run-dir"))
}

func kubeInstallNodes() gestalt.Component {
	return g.Group("install-nodes").
		Run(
			g.SH("install", "make", "install-nodes").
				Dir("{{run-dir}}").
				AddEnv("DOMAIN", "{{host-base}}")).
		Run(
			g.Retry(10).Run(
				g.SH("check", "make", "check-nodes").
					Dir("{{run-dir}}").
					AddEnv("DOMAIN", "{{host-base}}"))).
		WithMeta(g.Require("run-dir", "host-base"))
}

func kubeInstallProviders() gestalt.Component {
	return g.Group("install-nodes").
		Run(
			g.SH("install", "make", "install-nodes").
				Dir("{{run-dir}}").
				AddEnv("DOMAIN", "{{host-base}}")).
		Run(
			g.Retry(10).Run(
				g.SH("check", "make", "check-providers").
					Dir("{{run-dir}}").
					AddEnv("DOMAIN", "{{host-base}}"))).
		WithMeta(g.Require("run-dir", "host-base"))
}

func KubeSuite() gestalt.Component {
	key := newKey("master")
	// daddr := g.Ref("deployment-id")

	return g.Suite("main").
		Run(kubeInit()).
		Run(kubeInstallNodes()).
		Run(kubeInstallProviders()).
		Run(groupAccountSend(key))
}
