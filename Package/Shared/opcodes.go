package shared

const ( // opcodes
    THE_END = iota          // 0
    NOP = iota              // 1
    PUSH_CONST = iota       // 2
    EMIT_CONST = iota       // 3
    BINARY_OP = iota        // 4
    COMPARE_OP = iota       // 5
    PRINT_NEWLINE = iota    // 6
)

var NumberOfBytesInOperand = [255]int{
    /* THE_END */       0, 
    /* NOP */           0, 
    /* PUSH_CONST */    -1,
    /* EMIT_CONST */    0, 
    /* BINARY_OP */     1, 
    /* COMPARE_OP */    2, 
    /* PRINT_NEWLINE */ 0, 
}

