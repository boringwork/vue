// Package vue is the progressive framework for wasm applications.
package vue

import (
	"fmt"
	"reflect"
)

// Comp is a vue component.
type Comp struct {
	el       string
	tmpl     string
	data     interface{}
	methods  map[string]reflect.Value
	computed map[string]reflect.Value
	watchers map[string]reflect.Value
	props    map[string]struct{}
	subs     map[string]*Comp
	isSub    bool
}

// Component creates a new component from the given options.
func Component(options ...Option) *Comp {
	methods := make(map[string]reflect.Value, 0)
	computed := make(map[string]reflect.Value, 0)
	watches := make(map[string]reflect.Value, 0)
	props := make(map[string]struct{}, 0)
	subs := make(map[string]*Comp, 0)

	comp := &Comp{
		data:     struct{}{},
		methods:  methods,
		computed: computed,
		watchers: watches,
		props:    props,
		subs:     subs,
	}
	for _, option := range options {
		option(comp)
	}
	return comp
}

// newData creates new data from the function.
// Without a function the data of the component is returned.
func (comp *Comp) newData() reflect.Value {
	value := reflect.ValueOf(comp.data)
	if value.Type().Kind() != reflect.Func {
		return value
	}
	rets := value.Call(nil)
	if n := len(rets); n != 1 {
		must(fmt.Errorf("invalid return length: %d", n))
	}
	return rets[0]
}
