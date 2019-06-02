package assembly

import "github.com/sharpvik/scoops/Package/Shared"

// Based on ../Shared/opcodes.go/const(opcodes)
var OpcodeMap = map[string]uint8{
	"THE_END": shared.THE_END,
	"NOP": shared.NOP,
	"LOAD_BYTES": shared.LOAD_BYTES,
	"PRINT_OBJECT": shared.PRINT_OBJECT,
	"ARITHMETIC_OP": shared.ARITHMETIC_OP,
	"BINARY_OP": shared.BINARY_OP,
	"COMPARE_OP": shared.COMPARE_OP,
	"PRINT_NEWLINE": shared.PRINT_NEWLINE,
}
