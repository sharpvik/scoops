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

    case shared.MAKE_INT:
        var buffer []byte
        for i := 0; i < 8; i++ {
            v := interpreter.scope.data.Pop()
            b := v.(byte)
            buffer = append(buffer, b)
        }
        n := int64( binary.LittleEndian.Uint64(buffer) )
        o := primitives.NewInteger(n)
        interpreter.scope.data.Push(o)

    case shared.PRINT_OBJ:
        i := interpreter.scope.data.Peek()
        o := i.(*primitives.Integer)
        o.Print()

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
