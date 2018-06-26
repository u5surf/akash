package cmp

import (
	"github.com/ovrclk/akash/cmd/akash/constants"
	g "github.com/ovrclk/gestalt/builder"
	gx "github.com/ovrclk/gestalt/exec"
	"github.com/ovrclk/gestalt/vars"
)

var defaultAkashRoot = g.Ref("akash-root")

func akash(name string, args ...string) gx.Cmd {
	return akash_(defaultAkashRoot, name, args...)
}

func akash_n(name string, args ...string) gx.Cmd {
	return akash_n_(defaultAkashRoot, name, args...)
}

func akashd(name string, args ...string) gx.Cmd {
	cmd := g.EXEC("akashd-"+name,
		"{{akashd-path}}",
		append([]string{"-d", "{{akashd-root}}"}, args...)...)
	cmd.WithMeta(g.Require("akashd-path", "akashd-root"))
	return cmd
}

func akash_n_(root vars.Ref, name string, args ...string) gx.Cmd {
	cmd := akash_(root, name,
		append([]string{"-n", "{{akash-node}}"}, args...)...)
	cmd.WithMeta(g.Default("akash-node", constants.DefaultNode))
	return cmd
}

func akash_(root vars.Ref, name string, args ...string) gx.Cmd {
	cmd := g.EXEC("akash-"+name,
		"{{akash-path}}",
		append([]string{"-d", root.Var()}, args...)...)

	cmd.WithMeta(g.Require("akash-path", root.Name()))
	return cmd
}
