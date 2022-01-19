package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

const longSchema = `
#schema: {
	spec: widgetMap: [string]: #Widget
	raw: widgetArray: [for w in spec.widgetMap {w}]
}

#Widget: {
	title: string
	type: string
	id: int
}
`

const longVal1 = `
v: #schema & {
	spec: widgetMap: "1": {
		title: "test"
		type: "test"
		id: 123
	}
`

const longVal10 = `
v: #schema & {
	spec: widgetMap: "1": {
		title: "test"
		type: "test"
		id: 123
	}
	spec: widgetMap: "2": {
		title: "test"
		type: "test"
		id: 123
	}
	spec: widgetMap: "3": {
		title: "test"
		type: "test"
		id: 123
	}
	spec: widgetMap: "4": {
		title: "test"
		type: "test"
		id: 123
	}
	spec: widgetMap: "5": {
		title: "test"
		type: "test"
		id: 123
	}
	spec: widgetMap: "6": {
		title: "test"
		type: "test"
		id: 123
	}
	spec: widgetMap: "7": {
		title: "test"
		type: "test"
		id: 123
	}
	spec: widgetMap: "8": {
		title: "test"
		type: "test"
		id: 123
	}
	spec: widgetMap: "9": {
		title: "test"
		type: "test"
		id: 123
	}
	spec: widgetMap: "10": {
		title: "test"
		type: "test"
		id: 123
	}
}
`

func listComprehensionLong10() {
	var (
		c *cue.Context
		s cue.Value
		v cue.Value
	)
	c = cuecontext.New()
	s = c.CompileString(longSchema)
	v = c.CompileString(longVal10, cue.Scope(s))
	fmt.Sprint(v)
}

func listComprehensionLong1() {
	var (
		c *cue.Context
		s cue.Value
		v cue.Value
	)
	c = cuecontext.New()
	s = c.CompileString(longSchema)
	v = c.CompileString(longVal1, cue.Scope(s))
	fmt.Sprint(v)
}

const shortSchema = `
#schema: {
	spec: widgetMap: [string]: string
	raw: widgetArray: [for w in spec.widgetMap {w}]
}
`
const shortVal1 = `
v: #schema & {
	spec: widgetMap: "1": "test"
`

const shortVal10 = `
v: #schema & {
	spec: widgetMap: "1": "test"
	spec: widgetMap: "2": "test"
	spec: widgetMap: "3": "test"
	spec: widgetMap: "4": "test"
	spec: widgetMap: "5": "test"
	spec: widgetMap: "6": "test"
	spec: widgetMap: "7": "test"
	spec: widgetMap: "8": "test"
	spec: widgetMap: "9": "test"
	spec: widgetMap: "10": "test"
`

func listComprehensionShort10() {
	var (
		c *cue.Context
		s cue.Value
		v cue.Value
	)
	c = cuecontext.New()
	s = c.CompileString(shortSchema)
	v = c.CompileString(shortVal10, cue.Scope(s))
	fmt.Sprint(v)
}

func listComprehensionShort1() {
	var (
		c *cue.Context
		s cue.Value
		v cue.Value
	)
	c = cuecontext.New()
	s = c.CompileString(shortSchema)
	v = c.CompileString(shortVal1, cue.Scope(s))
	fmt.Sprint(v)
}
