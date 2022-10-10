package main

import "github.com/nstrlabs/lib"

type pluginBv0_0_2 struct {
}

func (pa *pluginBv0_0_2) Execute() string {
	return "hello from plugin B v0_0_2"
}

type factoryPluginBv0_0_2 struct {
}

var FactoryPluginBv_0_0_2 = factoryPluginBv0_0_2{}

func (f *factoryPluginBv0_0_2) New() lib.Feature {
	return &pluginBv0_0_2{}
}
