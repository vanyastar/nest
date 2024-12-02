package main

import (
	"log"
	"os"
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
					log.Fatal(e)
					return
				}
				return
			}
			log.Println("- not enough parameters")
		}
	}
	cmd["default"]["help"].Call(nil)

}
