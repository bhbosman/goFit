package fitDecl

import strconv "strconv"

type DiveGasMode byte

const (
	DiveGasMode_OpenCircuit          DiveGasMode = 0
	DiveGasMode_ClosedCircuitDiluent DiveGasMode = 1
	DiveGasMode_Invalid              DiveGasMode = 255
)

var DiveGasModeNames = map[DiveGasMode]string{
	DiveGasMode_OpenCircuit:          "OpenCircuit",
	DiveGasMode_ClosedCircuitDiluent: "ClosedCircuitDiluent",
}

func (k DiveGasMode) String() string {
	if uint(k) < uint(len(DiveGasModeNames)) {
		if v, ok := DiveGasModeNames[k]; ok {
			return v
		}
	}
	return "DiveGasMode(" + strconv.Itoa(int(k)) + ")"
}

var DiveGasModeValues = map[string]DiveGasMode{
	"OpenCircuit":          DiveGasMode_OpenCircuit,
	"ClosedCircuitDiluent": DiveGasMode_ClosedCircuitDiluent,
}

func DiveGasModeFromString(s string) DiveGasMode {
	if v, ok := DiveGasModeValues[s]; ok {
		return v
	}
	return DiveGasMode(DiveGasMode_Invalid)
}
