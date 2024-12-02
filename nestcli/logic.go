package main

import (
	"reflect"
	"strings"
)

var arguments = new(argumentsStack)

type argumentsStack struct {
	values []string
}

func (a *argumentsStack) Pop() string {
	if len(a.values) == 0 {
		return ""
	}
	t := a.values[:1][0]
	a.values = a.values[1:]
	return t
}

func (a *argumentsStack) Length() int {
	return len(a.values)
}

func (a *argumentsStack) Push(args []string) {
	for _, s := range args {
		a.values = append(a.values, strings.ToLower(s))
	}
}

var cmd = make(map[string]map[string]reflect.Value)

func registerStruct(i any) {
	sType := reflect.TypeOf(i)
	structName := strings.ToLower(sType.Name())
	sValue := reflect.ValueOf(i)

	for i := range sType.NumMethod() {
		if cmd[structName] == nil {
			cmd[structName] = make(map[string]reflect.Value)
		}
		methodName := strings.ToLower(sType.Method(i).Name)
		cmd[structName][methodName] = sValue.Method(i)
	}
}
