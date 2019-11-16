package shared

/*
 * If you were to ever edit the opcodes down here, make sure that you
 *     1. Run scoops/Tools/assembly_instructions_update.py
 *     2. Possibly run scoops/Tools/enumerate_opcodes.py
 *     3. Update if necessary...
 *         │ scoops/Package/Assembly/assembler_test.go
 *         │ scoops/InputFiles/scpa/*
 */

const (
	END byte = iota
	MAKE_BYTE
	MAKE_BLN
	MAKE_FLT
	MAKE_INT
	MAKE_NIL
	MAKE_RUNE
	MAKE_ASCII_RUNE
	MAKE_STRING
	STRING_CONCATENATE
	MAKE_SLICE
	SLICE_APPEND
	SLICE_GET_ITEM_BY_INDEX
	SLICE_POP
	MAKE_LINKED
	MAKE_SET
	MAKE_MAP
	MAKE_STACK
	STACK_PUSH
	STACK_POP
	STACK_CLEAR
	STACK_EMPTY
	STACK_PEEK
	MAKE_QUEUE
	QUEUE_PUSH
	QUEUE_POP
	QUEUE_CLEAR
	QUEUE_EMPTY
	QUEUE_PEEK
	MAKE_DEQUE
	CLONE_OBJ
	PRINT_OBJ
	GET_TYPE
	GET_SIZE
	MAKE_SCOOP
	INLINE_SCOOP
	SCOOP_CALL
	SCOOP_ACCEPT_ARG
	SCOOP_RETURN
	STORE_CONST
	LOAD_CONST
	STORE_VAR
	LOAD_VAR
	UNARY_NOT
	BINARY_AND
	BINARY_OR
	BINARY_XOR
	BINARY_ADD
	BINARY_SUB
	BINARY_MUL
	BINARY_DIV
	BINARY_MOD
	BINARY_POW
	PRINT_NEWLINE
	ABSOLUTE_JUMP
	RELATIVE_JUMP
	ABSOLUTE_JUMP_IF_FALSE
	RELATIVE_JUMP_IF_FALSE
	POP
)
