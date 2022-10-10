package main

import "github.com/nstrlabs/lib"

type pluginB struct {
}

func (pa *pluginB) Execute() string {
	return "hello from plugin B"
}

type factoryPluginB struct {
}

var FactoryPluginB = factoryPluginB{}

func (f *factoryPluginB) New() lib.Feature {
	return &pluginB{}
}
