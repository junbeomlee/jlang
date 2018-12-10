package jlang

type Instructions []byte

type Opcode byte

const (

	// Constant opcode represents constant value
	OpConstant Opcode = iota
)

// Description for opcode
type OpcodeDesc struct {

	// Name of opcodes
	Name string

	// List of widths of operands
	// Example1) Opcode operand1(2 byte size) -> OperandsWidth will be []int{2}
	// Example2) Opcode operand1(1 byte size) operand2(2 byte size) -> OperandsWidth will be []int{1,2}s
	OperandsWidth []int
}

var opDictionary map[Opcode]OpcodeDesc

func init() {
	opDictionary = make(map[Opcode]OpcodeDesc)
	opDictionary[OpConstant] = OpcodeDesc{"OpConstant", []int{2}}
}

//func Encode() []byte{
//
//}
