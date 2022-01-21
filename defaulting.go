package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

const default0Schema = `
#default: {
	text1: string
	text2: string
	text3: string
	text4: string
	text5: string
	text6: string
	text7: string
	text8: string
	text9: string
	text10: string
}
`

const default0Val = `
v: #default & {
	text1: "hello"
	text2: "hello"
	text3: "hello"
	text4: "hello"
	text5: "hello"
	text6: "hello"
	text7: "hello"
	text8: "hello"
	text9: "hello"
	text10: "hello"
}
`

const default10Schema = `
#default: {
	text1: string | *"hello"
	text2: string | *"hello"
	text3: string | *"hello"
	text4: string | *"hello"
	text5: string | *"hello"
	text6: string | *"hello"
	text7: string | *"hello"
	text8: string | *"hello"
	text9: string | *"hello"
	text10: string | *"hello"
}
`

const default10Val = `
v: #default & {
	text1: "hello"
	text2: "hello"
	text3: "hello"
	text4: "hello"
	text5: "hello"
	text6: "hello"
	text7: "hello"
	text8: "hello"
	text9: "hello"
	text10: "hello"
}
`

func default0() {
	var (
		c *cue.Context
		s cue.Value
		v cue.Value
	)
	c = cuecontext.New()
	s = c.CompileString(default0Schema)
	v = c.CompileString(default0Val, cue.Scope(s))
	fmt.Sprint(v)
}

func default10() {
	var (
		c *cue.Context
		s cue.Value
		v cue.Value
	)
	c = cuecontext.New()
	s = c.CompileString(default10Schema)
	v = c.CompileString(default10Val, cue.Scope(s))
	fmt.Sprint(v)
}
