package cmp

import (
	"github.com/ovrclk/gestalt"
	g "github.com/ovrclk/gestalt/builder"
	"github.com/ovrclk/gestalt/vars"
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

func KubeSuite() gestalt.Component {
	key := newKey("master")
	daddr := g.Ref("deployment-id")

	genAkashNode := g.FN("generate-node-host", func(e gestalt.Evaluator) error {
		host := vars.Expand(e.Vars(), "{{node-hostname}}.{{host-base}}:80")
		e.Vars().Put("akash-node", host)
		return nil
	}).WithMeta(g.Require("host-base", "node-hostname").Export("akash-node"))

	return g.Suite("main").
		Run(genAkashNode).
		Run(kubeInstall()).
		Run(keyGetAddress(key)).
		Run(groupAccountSend(key)).
		Run(groupDeploy(key.name, daddr))
}
