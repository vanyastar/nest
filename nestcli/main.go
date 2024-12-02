package main

import (
	"os"

	"github.com/vanyastar/nest/nestlog"
)

func main() {
	arguments.Push(os.Args[1:])

	registerStruct(Default{})
	registerStruct(Generate{})

	if st, ok := cmd[arguments.Pop()]; ok {
		if m, ok := st[arguments.Pop()]; ok {
			if arguments.Length() >= 2 {
				err := m.Call(nil)
				if !err[0].IsNil() {
					e := err[0].Interface().(error)
					nestlog.Error("CLI", e.Error())
					return
				}
				return
			}
			nestlog.Error("CLI", "Not enough parameters")
			return
		}
	}
	cmd["default"]["help"].Call(nil)

}
