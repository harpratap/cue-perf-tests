package main

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
)

const conjunctionSchema = `
#level1: {
	text1: #level2
}

#level2: {
	text2: string
}
`

const singleConjunctionVal = `
v1: #level1 & {
	text1: text2: "hello"
}

v2: #level1 & {
	text1: text2: "hello"
}

v3: #level1 & {
	text1: text2: "hello"
}

v4: #level1 & {
	text1: text2: "hello"
}

v5: #level1 & {
	text1: text2: "hello"
}

v6: #level1 & {
	text1: text2: "hello"
}

v7: #level1 & {
	text1: text2: "hello"
}

v8: #level1 & {
	text1: text2: "hello"
}

v9: #level1 & {
	text1: text2: "hello"
}

v10: #level1 & {
	text1: text2: "hello"
}
`

const doubleConjunctionVal = `
v1: #level1 & {
	text1: #level2 & {
		text2: "hello"
	}
}

v2: #level1 & {
	text1: #level2 & {
		text2: "hello"
	}
}

v3: #level1 & {
	text1: #level2 & {
		text2: "hello"
	}
}

v4: #level1 & {
	text1: #level2 & {
		text2: "hello"
	}
}

v5: #level1 & {
	text1: #level2 & {
		text2: "hello"
	}
}

v6: #level1 & {
	text1: #level2 & {
		text2: "hello"
	}
}

v7: #level1 & {
	text1: #level2 & {
		text2: "hello"
	}
}

v8: #level1 & {
	text1: #level2 & {
		text2: "hello"
	}
}

v9: #level1 & {
	text1: #level2 & {
		text2: "hello"
	}
}

v10: #level1 & {
	text1: #level2 & {
		text2: "hello"
	}
}
`

func conjunctionSingle() {
	var (
		c *cue.Context
		s cue.Value
		v cue.Value
	)
	c = cuecontext.New()
	s = c.CompileString(conjunctionSchema)
	v = c.CompileString(singleConjunctionVal, cue.Scope(s))
	fmt.Sprint(v)
}

func conjunctionDouble() {
	var (
		c *cue.Context
		s cue.Value
		v cue.Value
	)
	c = cuecontext.New()
	s = c.CompileString(conjunctionSchema)
	v = c.CompileString(doubleConjunctionVal, cue.Scope(s))
	fmt.Sprint(v)
}
