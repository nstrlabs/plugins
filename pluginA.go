package main

import (
	"fmt"
	"github.com/nstrlabs/lib"
)

type PluginA struct {
	pluginName string
}

func (pa *PluginA) Execute() string {
	return "hello from plugin " + pa.pluginName
}

type FactoryPluginA struct {
}

var FactoryPlugin = FactoryPluginA{}

func (f *FactoryPluginA) New(configuration map[string]interface{}) lib.Feature {
	return &PluginA{pluginName: configuration["name"].(string)}
}

type PluginAValidator func(configuration map[string]interface{}) error

func (v PluginAValidator) Validate(configuration map[string]interface{}) error {
	name, ok := configuration["name"]
	if !ok {
		return fmt.Errorf("name not present")
	}
	_, ok = name.(string)
	if !ok {
		return fmt.Errorf("name should be a string")
	}
	return nil
}
