package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

const openSchema = `
#def1: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}

#def2: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}

#def3: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}

#def4: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}

#def5: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}

#def6: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}

#def7: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}

#def8: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}

#def9: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}

#def10: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}
#def11: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}

#def12: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}

#def13: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}

#def14: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}

#def15: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}

#def16: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}

#def17: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}

#def18: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}

#def19: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}

#def20: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
	...
}
`

const closedSchema = `
#def1: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def2: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def3: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def4: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def5: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def6: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def7: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def8: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def9: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def10: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def11: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def12: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def13: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def14: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def15: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def16: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def17: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def18: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def19: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}

#def20: {
	text: string
	text1?: string
	text2?: string
	text3?: string
	text4?: string
}
`

const opennessVal = `
v1: #def1 & {
	text: "hello"
}

v2: #def2 & {
	text: "hello"
}

v3: #def3 & {
	text: "hello"
}

v4: #def4 & {
	text: "hello"
}

v5: #def5 & {
	text: "hello"
}

v6: #def6 & {
	text: "hello"
}

v7: #def7 & {
	text: "hello"
}

v8: #def8 & {
	text: "hello"
}

v9: #def9 & {
	text: "hello"
}

v10: #def10 & {
	text: "hello"
}

v11: #def11 & {
	text: "hello"
}

v12: #def12 & {
	text: "hello"
}

v13: #def13 & {
	text: "hello"
}

v14: #def14 & {
	text: "hello"
}

v15: #def15 & {
	text: "hello"
}

v16: #def16 & {
	text: "hello"
}

v17: #def17 & {
	text: "hello"
}

v18: #def18 & {
	text: "hello"
}

v19: #def19 & {
	text: "hello"
}

v20: #def20 & {
	text: "hello"
}
`

func definitionOpen() {
	var (
		c *cue.Context
		s cue.Value
		v cue.Value
	)
	c = cuecontext.New()
	s = c.CompileString(openSchema)
	v = c.CompileString(opennessVal, cue.Scope(s))
	fmt.Sprint(v)
}

func defintionClosed() {
	var (
		c *cue.Context
		s cue.Value
		v cue.Value
	)
	c = cuecontext.New()
	s = c.CompileString(closedSchema)
	v = c.CompileString(opennessVal, cue.Scope(s))
	fmt.Sprint(v)
}
