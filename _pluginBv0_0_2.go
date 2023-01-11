package main

import (
	"FeatureSDK/pkg/feature"
	"fmt"
)

type pluginBv0_0_2 struct {
	field string
}

func (pa *pluginBv0_0_2) Execute(msg feature.InMsg) error {
	// pluginBv0_0_2 gets the last bytes of the raw message
	return nil
}

type factoryPluginBv0_0_2 struct {
}

var FactoryPluginBv_0_0_2 = factoryPluginBv0_0_2{}

func (f *factoryPluginBv0_0_2) New(configuration map[string]interface{}) feature.Feature {
	return &pluginBv0_0_2{field: configuration["name"].(string)}
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
