package jlang

type VM struct{
	rawByteCode []byte
}

func NewVM(rawByteCode []byte) *VM{
	return &VM{
		rawByteCode:rawByteCode,
	}
}