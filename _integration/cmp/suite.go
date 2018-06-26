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

func KubeSuite() gestalt.Component {
	// key := newKey("master")
	// daddr := g.Ref("deployment-id")

	return g.Suite("main").
		Run(kubeInstall())
	// Run(groupAccountSend(key))
}
