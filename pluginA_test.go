package main

import "testing"

func BenchmarkPlugin(b *testing.B) {
	feature := FactoryPlugin.New(map[string]interface{}{
		"name": "A",
	})
	m := NewMessage([]byte{0, 1, 2, 3, 4, 5, 6, 7})
	b.Run("pluginA", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = feature.Execute(m)
		}
	})
}

type msg struct {
	raw    []byte
	fields map[string][]byte
}

func (m *msg) Add(key string, value []byte) {
	m.fields[key] = value
}

func (m *msg) Get(key string) []byte {
	return m.fields[key]
}

func (m *msg) GetRaw() []byte {
	return m.raw
}

func NewMessage(raw []byte) *msg {
	return &msg{
		raw:    raw,
		fields: make(map[string][]byte, 16),
	}
}
