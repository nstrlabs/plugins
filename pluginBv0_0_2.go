package main

import (
	"fmt"
	"github.com/nstrlabs/lib"
)

type pluginBv0_0_2 struct {
	pluginName string
}

func (pa *pluginBv0_0_2) Execute() string {
	return "hello from plugin B v0_0_2"
}

type factoryPluginBv0_0_2 struct {
}

var FactoryPluginBv_0_0_2 = factoryPluginBv0_0_2{}

func (f *factoryPluginBv0_0_2) New(configuration map[string]interface{}) lib.Feature {
	return &pluginBv0_0_2{pluginName: configuration["name"].(string)}
}

type PluginBv0_0_2Validator func(configuration map[string]interface{}) error

func (v PluginBv0_0_2Validator) Validate(configuration map[string]interface{}) error {
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

var ValidatorPluginB2 PluginBv0_0_2Validator
