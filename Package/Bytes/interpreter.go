package bytes

import (
    "encoding/binary"
    "errors"
    "fmt"
    "github.com/sharpvik/scoops/Package/DataTypes/Primitives"
    "github.com/sharpvik/scoops/Package/Shared"
    "github.com/sharpvik/scoops/Package/Util"
)



func (interpreter *Interpreter) Evaluate() {
    instruction := *(interpreter.code[interpreter.ip])
    //fmt.Println(instruction)
    switch instruction.opcode {

    case shared.END:
        interpreter.running = false

    case shared.PUSH_CONST:
        interpreter.scope.data.Push(instruction.operand)
        
    case shared.MAKE_BLN:
        i := interpreter.scope.data.Pop().(byte)
        interpreter.scope.data.Push( primitives.NewBoolean(i) )

    case shared.MAKE_INT:
        var buffer []byte
        for i := 0; i < 8; i++ {
            buffer = append( buffer, interpreter.scope.data.Pop().(byte) )
        }
        n := int64( binary.LittleEndian.Uint64(buffer) )
        interpreter.scope.data.Push( primitives.NewInteger(n) )
        
    case shared.MAKE_NIL:
        interpreter.scope.data.Push(interpreter.thenil)

    case shared.MAKE_RUNE:
        var buffer []byte
        c := int(instruction.operand)
        for i := 0; i < c; i++ {
            buffer = append( buffer, interpreter.scope.data.Pop().(byte) )
        }
        interpreter.scope.data.Push( primitives.NewRune(buffer) )

    case shared.BINARY_ADD:
        x := interpreter.scope.data.Pop().(Object)
        y := interpreter.scope.data.Pop().(Object)
        _type := x.Type() + y.Type()
        switch _type {
        case "intint":
            a := x.(*primitives.Integer)
            b := y.(*primitives.Integer)
            c := primitives.AddInteger(a, b)
            interpreter.scope.data.Push(c)
        default:
            interpreter.err = errors.New("Unknown numeric data type.")
        }
        
    case shared.BINARY_SUB:
        x := interpreter.scope.data.Pop().(Object)
        y := interpreter.scope.data.Pop().(Object)
        _type := x.Type() + y.Type()
        switch _type {
        case "intint":
            a := x.(*primitives.Integer)
            b := y.(*primitives.Integer)
            c := primitives.SubInteger(a, b)
            interpreter.scope.data.Push(c)
        default:
            interpreter.err = errors.New("Unknown numeric data type.")
        }
        
    case shared.BINARY_MUL:
        x := interpreter.scope.data.Pop().(Object)
        y := interpreter.scope.data.Pop().(Object)
        _type := x.Type() + y.Type()
        switch _type {
        case "intint":
            a := x.(*primitives.Integer)
            b := y.(*primitives.Integer)
            c := primitives.MulInteger(a, b)
            interpreter.scope.data.Push(c)
        default:
            interpreter.err = errors.New("Unknown numeric data type.")
        }

    case shared.PRINT_OBJ:
        interpreter.scope.data.Peek().(Object).Print()

    case shared.PRINT_NEWLINE:
        fmt.Print("\n")

    case shared.POP:
        interpreter.scope.data.Pop()

    default:
        interpreter.err = errors.New("Unknown opcode.")

    }
    
    /* Instruction pointer is incremented by default. Use 'return' to force
     * it not to.
     */
    interpreter.ip++
}


func (interpreter *Interpreter) Execute() {
    for interpreter.running && interpreter.err == nil {
        interpreter.Evaluate()
    }
    if interpreter.err != nil {
        util.Error(interpreter.err)
        util.Log("Interpreter exited with non-zero return value.")
    }
}
