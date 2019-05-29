package shared

/*
 * If you were to ever edit the opcodes down here, make sure that you
 * also update... 
 *     │ scoops/Package/Assembly/
 *     ├──── assembler.go
 *     └──── assembler_test.go
 *
 *     │ scoops/InputFiles/scpa/*
 */
const ( /* opcodes */
    THE_END uint8 = iota            // 0
    NOP uint8 = iota                // 1
    LOAD_BYTES uint8 = iota         // 2
    PRINT_OBJECT uint8 = iota       // 3
    ARITHMETIC_OP uint8 = iota      // 4
    BINARY_OP uint8 = iota          // 5
    COMPARE_OP uint8 = iota         // 6
    PRINT_NEWLINE uint8 = iota      // 7
)

var NumberOfBytesInOperand = [255]int{
    /* THE_END */           0,
    /* NOP */               0,
    /* LOAD_BYTES */       -1, // number of bytes in operand = first byte
    /* PRINT_OBJECT */      0,
    /* ARITHMETIC_OP */     1,
    /* BINARY_OP */         2,
    /* COMPARE_OP */        2,
    /* PRINT_NEWLINE */     0,
}

