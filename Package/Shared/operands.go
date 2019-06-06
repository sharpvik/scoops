package shared

var NumberOfOperands = [256]uint32{
    /* THE_END */               0,
    /* NOP */                   0,
    /* LOAD_CONST */            1,
    /* MAKE_PRIMITIVE */        1,
    /* STORE_VAR */             1,
    /* LOAD_VAR */              1,
    /* BINARY_OP */             1,
    /* MAKE_COLLECTION */       2,
    /* PRINT_OBJ */             0,
    /* PRINT_NEWLINE */         0,
    /* ABSOLUTE_JUMP */         1,
    /* RELATIVE_JUMP */         1,
    /* JUMP_IF_TRUE */          1,
    /* JUMP_IF_FALSE */         1,
    /* MAKE_SCOOP */            2,
    /* CALL_SCOOP */            1,
    /* RETURN */                1,
}
