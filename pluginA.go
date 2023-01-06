package main

import (
	"fmt"
	"github.com/nstrlabs/lib"
)

type PluginA struct {
	field string
}

func (pa *PluginA) Execute(msg lib.Msg) error {
	// plugin a gets the first bytes of the raw message
	msg.Add(pa.field, msg.GetRaw()[:4])
	return nil
}

type FactoryPluginA struct {
}

var FactoryPlugin = FactoryPluginA{}

func (f *FactoryPluginA) New(configuration map[string]interface{}) lib.Feature {
	return &PluginA{field: configuration["name"].(string)}
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

var ValidatorPluginA PluginAValidator
