package main

import (
	"fmt"
	"github.com/nstrlabs/FeatureSDK/pkg/feature"
)

type PluginA struct {
	feature.DefaultStatelessFeature
	field string
}

func (pa *PluginA) Execute(msg feature.Message) (string, error) {
	// plugin a gets the first bytes of the raw message
	value, _ := feature.NewValueFromBytes(msg.GetText()[:4], feature.String)
	_ = msg.AddEventField(pa.field, value)
	return "default", nil
}

type FactoryPluginA struct {
}

var FactoryPlugin = FactoryPluginA{}

func (f *FactoryPluginA) NewFeature(params feature.FactoryParams) (feature.
Feature, error) {
	return &PluginA{field: params.Configuration["name"].(string)}, nil
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
