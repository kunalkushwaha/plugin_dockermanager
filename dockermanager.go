package dockermanager

import (
	"log"

	"github.com/itsabot/abot/shared/datatypes"
	"github.com/itsabot/abot/shared/plugin"
)

var p *dt.Plugin

func init() {
	// Create the plugin, setting it up to communicate with Abot through
	// the functions we specified.
	var err error
	p, err = plugin.New("github.com/kunalkushwaha/plugin_dockermanager")
	if err != nil {
		log.Fatalln("failed to build plugin.", err)
	}

	// Trigger & Identifier objects
	plugin.AppendTrigger(p, &dt.StructuredInput{
		Commands: []string{"DM"},
		Objects:  []string{"create", "list", "build"},
	})

	// State machine
	plugin.SetStates(p, [][]dt.State{[]dt.State{
		{
			OnEntry: func(in *dt.Msg) string {
				return "Hello world!"
			},
			OnInput: func(in *dt.Msg) {
			},
			Complete: func(in *dt.Msg) (bool, string) {
				return true, ""
			},
		},
	}})
}
