package jlang

type Compiler struct {
}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (c Compiler) Compile() []byte {
	return []byte{}
}
