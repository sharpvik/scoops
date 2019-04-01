#pragma once

#include "ops.h"



char * exec(char *ip)
{
    switch (*ip)
    {
        case 10:
            op_hello();
            return ip + 1;
            break;
            
        case 97:
            op_foo();
            return ip + 1;
            break;
        
        default:
            op_nop();
            return ip + 1;
    }
}