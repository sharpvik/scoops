package shared

var NumberOfBytesInOperand = [256]int{
    /* THE_END */           0,
    /* NOP */               0,
    /* LOAD_BYTES */       -1, // number of bytes in operand = first byte
    /* PRINT_OBJECT */      0,
    /* ARITHMETIC_OP */     1,
    /* BINARY_OP */         2,
    /* COMPARE_OP */        2,
    /* PRINT_NEWLINE */     0,
}
