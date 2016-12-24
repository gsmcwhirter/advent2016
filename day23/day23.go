package day23

import (
	"fmt"
	"strings"

	"strconv"

	"github.com/gsmcwhirter/advent2016/lib"
)

type ArgumentType int
type RegisterName string
type InstructionType int

const (
	REGISTER ArgumentType = iota
	VALUE
)

const (
	CPY InstructionType = iota
	INC
	DEC
	JNZ
	TGL
)

type Argument struct {
	Type     ArgumentType
	Register RegisterName
	Value    int
}

func (a *Argument) ToString() string {
	switch a.Type {
	case REGISTER:
		return fmt.Sprintf("%v", a.Register)
	case VALUE:
		return fmt.Sprintf("%v", a.Value)
	default:
		return "unknown"
	}
}

type Instruction struct {
	Name    InstructionType
	Args    []Argument
	Toggled bool
}

func NewInstruction(itype InstructionType, args ...string) Instruction {
	if len(args) < 1 || len(args) > 2 {
		panic("Too few or too many arguments")
	}

	realArgs := make([]Argument, len(args))
	for i := 0; i < lib.IntMin(2, len(args)); i++ {
		val, err := strconv.Atoi(args[i])
		if err != nil {
			realArgs[i] = Argument{REGISTER, RegisterName(args[i]), 0}
		} else {
			realArgs[i] = Argument{VALUE, RegisterName(""), val}
		}
	}

	return Instruction{itype, realArgs, false}
}

func (i *Instruction) Copy() Instruction {
	return Instruction{
		Name:    i.Name,
		Args:    append([]Argument{}, i.Args...),
		Toggled: i.Toggled,
	}
}

func (i *Instruction) ToString() string {
	tStr := ""
	switch i.Name {
	case CPY:
		tStr = "CPY"
	case INC:
		tStr = "INC"
	case DEC:
		tStr = "DEC"
	case JNZ:
		tStr = "JNZ"
	case TGL:
		tStr = "TGL"
	}

	argStrings := make([]string, len(i.Args))
	for i, arg := range i.Args {
		argStrings[i] = arg.ToString()
	}

	return fmt.Sprintf("%v %v (toggled=%v)", tStr, argStrings, i.Toggled)
}

type Machine struct {
	InstructionPointer int
	Instructions       []Instruction
	Registers          map[RegisterName]int
}

func NewMachine(instructions []Instruction) Machine {
	return Machine{
		InstructionPointer: 0,
		Instructions:       instructions,
		Registers: map[RegisterName]int{
			RegisterName("a"): 0,
			RegisterName("b"): 0,
			RegisterName("c"): 0,
			RegisterName("d"): 0,
		},
	}
}

func (m *Machine) ExecuteInstruction(inst Instruction) {
	var tmp int
	var newInst Instruction

	switch inst.Name {
	case INC:
		if inst.Toggled {
			newInst = inst.Copy()
			newInst.Name = DEC
			newInst.Toggled = false
			m.ExecuteInstruction(newInst)
		} else {
			if inst.Args[0].Type == REGISTER {
				m.Registers[inst.Args[0].Register]++
			}
			m.InstructionPointer++
		}

	case DEC:
		if inst.Toggled {
			newInst = inst.Copy()
			newInst.Name = INC
			newInst.Toggled = false
			m.ExecuteInstruction(newInst)
		} else {
			if inst.Args[0].Type == REGISTER {
				m.Registers[inst.Args[0].Register]--
			}
			m.InstructionPointer++
		}

	case CPY:
		if inst.Toggled {
			newInst = inst.Copy()
			newInst.Name = JNZ
			newInst.Toggled = false
			m.ExecuteInstruction(newInst)
		} else {
			if inst.Args[0].Type == REGISTER {
				tmp = m.Registers[inst.Args[0].Register]
			} else {
				tmp = inst.Args[0].Value
			}

			if inst.Args[1].Type == REGISTER {
				m.Registers[inst.Args[1].Register] = tmp
			}
			m.InstructionPointer++
		}

	case JNZ:
		if inst.Toggled {
			newInst = inst.Copy()
			newInst.Name = CPY
			newInst.Toggled = false
			m.ExecuteInstruction(newInst)
		} else {
			if inst.Args[0].Type == REGISTER {
				tmp = m.Registers[inst.Args[0].Register]
			} else {
				tmp = inst.Args[0].Value
			}

			//issue here?
			var mod int
			if inst.Args[1].Type == REGISTER {
				mod = m.Registers[inst.Args[1].Register]
			} else {
				mod = inst.Args[1].Value
			}

			if tmp != 0 {
				m.InstructionPointer += mod
			} else {
				m.InstructionPointer++
			}
		}

	case TGL:
		if inst.Toggled {
			newInst = inst.Copy()
			newInst.Name = INC
			newInst.Toggled = false
			m.ExecuteInstruction(newInst)
		} else {
			var togglePointer int
			if inst.Args[0].Type == REGISTER {
				togglePointer = m.InstructionPointer + m.Registers[inst.Args[0].Register]
			} else {
				togglePointer = m.InstructionPointer + inst.Args[0].Value
			}

			if togglePointer < len(m.Instructions) && togglePointer >= 0 {
				m.Instructions[togglePointer].Toggled = !m.Instructions[togglePointer].Toggled
			}
			m.InstructionPointer++
		}
	}
}

func (m *Machine) ExecuteNextInstruction() bool {
	if m.InstructionPointer >= len(m.Instructions) {
		return false
	}

	inst := m.Instructions[m.InstructionPointer]
	m.ExecuteInstruction(inst)

	return true
}

func (m *Machine) CurrentInstruction() string {
	if m.InstructionPointer >= len(m.Instructions) {
		return "(done)"
	}

	return m.Instructions[m.InstructionPointer].ToString()
}

func (m *Machine) PrintState() {
	fmt.Printf("IPointer: %v\n", m.InstructionPointer)
	fmt.Printf("Instruction: %v\n", m.CurrentInstruction())
	fmt.Print("Registers: ")
	fmt.Printf("a=%v ", m.Registers[RegisterName("a")])
	fmt.Printf("b=%v ", m.Registers[RegisterName("b")])
	fmt.Printf("c=%v ", m.Registers[RegisterName("c")])
	fmt.Printf("d=%v ", m.Registers[RegisterName("d")])
	fmt.Println()
}

func (m *Machine) PrintInstructions() {
	for i, inst := range m.Instructions {
		fmt.Printf("%v: %v\n", i, inst.ToString())
	}
}

func ParseDataString(dat string, initA int) Machine {
	lines := strings.Split(dat, "\n")
	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")

		switch parts[0] {
		case "cpy":
			instructions[i] = NewInstruction(CPY, parts[1:]...)
		case "inc":
			instructions[i] = NewInstruction(INC, parts[1:]...)
		case "dec":
			instructions[i] = NewInstruction(DEC, parts[1:]...)
		case "jnz":
			instructions[i] = NewInstruction(JNZ, parts[1:]...)
		case "tgl":
			instructions[i] = NewInstruction(TGL, parts[1:]...)
		}
	}

	machine := NewMachine(instructions)
	machine.Registers[RegisterName("a")] = initA
	return machine
}

func LoadData(filename string, initA int) Machine {
	dat := lib.ReadFileData(filename)

	return ParseDataString(strings.Trim(string(dat), "\n"), initA)
}

func RunPartA(filename string) {
	machine := LoadData(filename, 7)

	machine.PrintState()
	machine.PrintInstructions()

	for machine.ExecuteNextInstruction() {
		machine.PrintState()
	}
	machine.PrintState()
}

func RunPartB(filename string) {
	machine := LoadData(filename, 12)

	machine.PrintState()
	machine.PrintInstructions()

	for machine.ExecuteNextInstruction() {
		// machine.PrintState()
	}
	machine.PrintState()
}
