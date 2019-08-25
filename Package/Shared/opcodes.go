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
    END             byte = iota	// 0
    PUSH_CONST	// 1
    MAKE_BLN	// 2
    MAKE_FLT	// 3
    MAKE_INT	// 4
    MAKE_NIL	// 5
    MAKE_RUNE	// 6
    MAKE_STRING	// 7
    STRING_CONCATENATE	// 8
    MAKE_SLICE	// 9
    SLICE_APPEND	// 10
    SLICE_GET_ITEM_BY_INDEX	// 11
    SLICE_POP	// 12
    MAKE_LINKED	// 13
    MAKE_SET	// 14
    MAKE_MAP	// 15
    MAKE_STACK	// 16
    STACK_PUSH	// 17
    STACK_POP	// 18
    STACK_CLEAR	// 19
    STACK_EMPTY	// 20
    STACK_PEEK	// 21
    MAKE_QUEUE	// 22
    QUEUE_PUSH	// 23
    QUEUE_POP	// 24
    QUEUE_CLEAR	// 25
    QUEUE_EMPTY	// 26
    QUEUE_PEEK	// 27
    MAKE_DEQUE	// 28
    MAKE_SCOOP	// 29
    CLONE_OBJ	// 30
    PRINT_OBJ	// 31
    GET_TYPE	// 32
    GET_SIZE	// 33
    STORE_VAR	// 34
    LOAD_VAR	// 35
    UNARY_NOT	// 36
    BINARY_AND	// 37
    BINARY_OR	// 38
    BINARY_XOR	// 39
    BINARY_ADD	// 40
    BINARY_SUB	// 41
    BINARY_MUL	// 42
    BINARY_MOD		// 43
    BINARY_DIV	// 44
    BINARY_POW	// 45
    PRINT_NEWLINE	// 46
    ABSOLUTE_JUMP	// 47
    RELATIVE_JUMP	// 48
    JUMP_IF_TRUE	// 49
    JUMP_IF_FALSE	// 50
    CALL_SCOOP	// 51
    RETURN	// 52
    POP	// 53
)
