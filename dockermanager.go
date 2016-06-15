package dockermanager

import (
	"github.com/itsabot/abot/core/log"
	"github.com/itsabot/abot/shared/datatypes"
	"github.com/itsabot/abot/shared/plugin"
)

var p *dt.Plugin

const memKey = "firstToken"

func init() {
	var err error
	p, err = plugin.New("github.com/kunalkushwaha/plugin_dock")
	if err != nil {
		log.Fatal("failed to build plugin plugin_dock", err)
	}
	plugin.SetKeywords(p,
		dt.KeywordHandler{
			Fn: kwDemo,
			Trigger: &dt.StructuredInput{
				Commands: []string{
					"docker", "dm", "dockermanager",
				},
				Objects: []string{
					"demo", "build", "list", "create",
				},
			},
		},
	)
	plugin.SetStates(p, [][]dt.State{[]dt.State{
		{
			OnEntry: func(in *dt.Msg) string {
				return "This is a demo."
			},
			OnInput: func(in *dt.Msg) {
				if len(in.Tokens) == 0 {
					return
				}
				p.SetMemory(in, memKey, in.Tokens[0])
			},
			Complete: func(in *dt.Msg) (bool, string) {
				return p.HasMemory(in, memKey), "I didn't understand that."
			},
		},
	}})
	p.SM.SetOnReset(func(in *dt.Msg) {
		p.DeleteMemory(in, memKey)
	})
	if err = plugin.Register(p); err != nil {
		p.Log.Fatalf("failed to register plugin plugin_dock. %s", err)
	}
}

func kwDemo(in *dt.Msg) string {
	return "It worked! You typed: " + in.Sentence
}
