package main

import (
	"fmt"
	"github.com/nstrlabs/lib"
)

type pluginB struct {
	pluginName string
}

func (pa *pluginB) Execute() string {
	return "hello from plugin " + pa.pluginName
}

type factoryPluginB struct {
}

var FactoryPluginB = factoryPluginB{}

func (f *factoryPluginB) New(configuration map[string]interface{}) lib.Feature {
	return &pluginB{pluginName: configuration["name"].(string)}
}

type PluginBValidator func(configuration map[string]interface{}) error

func (v PluginBValidator) Validate(configuration map[string]interface{}) error {
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

var ValidatorPluginB PluginBValidator
