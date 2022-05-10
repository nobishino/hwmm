package hwmm

type Machine struct {
	Threads      Thread
	SharedMemory SharedMemory
}

func NewMachine() *Machine {
	return &Machine{
		SharedMemory: NewSharedMemory(),
	}
}

func (m *Machine) AddThread(ii ...Instruction) {
	m.Threads = Thread{
		Registers:    map[int]int{},
		Instructions: ii,
		Memory:       m.SharedMemory,
	}
}

func (m *Machine) Run() Result {
	for _, i := range m.Threads.Instructions {
		switch i.Type {
		case W:
			m.Threads.Memory[i.VarName] = i.Value
		}
	}
	return Result{
		SharedMemory: m.Threads.Memory,
	}
}

type Thread struct {
	Registers    map[int]int
	Instructions []Instruction
	Memory       SharedMemory
}

type Instruction struct {
	Type    InstructionType // R or W
	VarName string          // Destination for W, Source for R
	Value   int             // For W only
}

type InstructionType int

const (
	R InstructionType = iota
	W
)

func Write(varName string, value int) Instruction {
	return Instruction{
		Type:    W,
		VarName: varName,
		Value:   value,
	}
}

type Result struct {
	SharedMemory
}

type SharedMemory map[string]int

func NewSharedMemory() SharedMemory {
	return SharedMemory{}
}
