package shared

const ( // opcodes
    THE_END uint8 = iota          // 0
    NOP uint8 = iota              // 1
    PUSH_CONST uint8 = iota       // 2
    EMIT_CONST uint8 = iota       // 3
    BINARY_OP uint8 = iota        // 4
    COMPARE_OP uint8 = iota       // 5
    PRINT_NEWLINE uint8 = iota    // 6
)

var NumberOfBytesInOperand = [255]int{
    /* THE_END */       0,
    /* NOP */           0,
    /* PUSH_CONST */    -1, // number of bytes is specified in the first byte
    /* EMIT_CONST */    0,
    /* BINARY_OP */     2,
    /* COMPARE_OP */    2,
    /* PRINT_NEWLINE */ 0,
}

