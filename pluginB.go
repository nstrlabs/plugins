package main

import (
	"fmt"
	"github.com/nstrlabs/FeatureSDK/pkg/feature"
)

type PluginB struct {
	feature.DefaultStatelessFeature
	field string
}

func (pa *PluginB) Execute(msg feature.Message) (string, error) {
	// plugin a gets the first bytes of the raw message
	value, _ := feature.NewValueFromBytes(msg.GetText()[4:], feature.String)
	_ = msg.AddEventField(pa.field, value)
	return "default", nil
}

type factoryPluginB struct {
}

var FactoryPluginB = factoryPluginB{}

func (f *factoryPluginB) NewFeature(params feature.FactoryParams) (feature.
Feature, error) {
	return &PluginB{field: params.Configuration["name"].(string)}, nil
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
