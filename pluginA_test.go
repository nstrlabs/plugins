package main

import (
	"github.com/nstrlabs/FeatureSDK/pkg/feature"
	"testing"
	"time"
)

func BenchmarkPlugin(b *testing.B) {
	feature := FactoryPlugin.New(feature.
	FactoryParams{Configuration: map[string]interface{}{
		"name": "A",
	}})
	m := NewMessage([]byte{0, 1, 2, 3, 4, 5, 6, 7})
	b.Run("pluginA", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = feature.Execute(m)
		}
	})
}

type msg struct {
	raw    []byte
	fields map[string]feature.Value
}

func (m *msg) AddEventField(key string, value feature.Value) bool {
	m.fields[key] = value
	return true
}

func (m *msg) GetEventField(key string) (feature.Value, bool) {
	result, ok := m.fields[key]
	return result, ok
}

func (m *msg) GetProtocolField(key string) (feature.Value, bool) {
	return feature.Value{}, false
}

func (m *msg) GetProtocol() feature.Protocol {
	return feature.Syslog
}

func (m *msg) GetTimestamp() time.Time {
	return time.Now()
}

func (m *msg) GetText() []byte {
	return m.raw
}

func NewMessage(raw []byte) *msg {
	return &msg{
		raw:    raw,
		fields: make(map[string]feature.Value, 16),
	}
}
