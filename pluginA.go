package main

import "github.com/nstrlabs/lib"

type PluginA struct {
}

func (pa *PluginA) Execute() string {
	return "hello from plugin A"
}

type FactoryPluginA struct {
}

var FactoryPlugin = FactoryPluginA{}

func (f *FactoryPluginA) New() lib.Feature {
	return &PluginA{}
}
