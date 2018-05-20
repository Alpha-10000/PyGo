package main

import "fmt"
import "os"

type VirtualMachine struct {
	calls Stack
	current Frame
}

func (vm* VirtualMachine)parseInstruction(code []byte) Instruction {
	cur := &vm.current

	instr := Instruction{}
	instr.opcode = int(code[cur.offset])
	cur.offset += 1

	if instr.opcode < HAS_ARGS {
		return instr
	}

	instr.arg = int(code[cur.offset]) + int(code[cur.offset+1])*256 
	cur.offset += 2

	return instr
}

func (vm* VirtualMachine)Run(code []byte) {
	fmt.Println("VM is running")
	vm.parseInstruction(code)
}

func interp(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Impossible to open file %s\n", fileName)
		os.Exit(1)
	}
	byteCode := make([]byte, 1024)
	_, err = f.Read(byteCode)
	if err != nil {
		fmt.Printf("Impossible to read file %s\n", fileName)
		os.Exit(1)
	}
	vm := VirtualMachine{}
	fmt.Printf("%+v\n", vm)
	vm.Run(byteCode)
}
