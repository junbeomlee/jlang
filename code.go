package jlang

type Instructions []byte

type Opcode byte

const (
	OpConstant Opcode = iota
)