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
	log.Println("In docker plugin")
	if err != nil {
		log.Fatalln("failed to build plugin.", err)
	}

	// Trigger & Identifier objects
	plugin.AppendTrigger(p, &dt.StructuredInput{
		Commands: []string{"DM", "DockerManager", "Docker"},
		Objects:  []string{"create", "list", "build"},
	})

	// State machine
	plugin.SetStates(p, [][]dt.State{[]dt.State{
		{
			OnEntry: func(in *dt.Msg) string {
				log.Println("In docker plugin OnEntry")
				return "Hello world!"
			},
			OnInput: func(in *dt.Msg) {
			},
			Complete: func(in *dt.Msg) (bool, string) {
				log.Println("In docker plugin Complete")
				return true, "Docker has performed!"
			},
		},
	}})
	p.SM.SetOnReset(func(in *dt.Msg) {
		//p.DeleteMemory(in, memKey)
	})
	if err = plugin.Register(p); err != nil {
		p.Log.Fatalf("failed to register plugin plugin_dock. %s", err)
	}
}
