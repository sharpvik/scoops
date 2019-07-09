package bytes

import (
    "encoding/binary"
    "errors"
    "fmt"
    "github.com/sharpvik/scoops/Package/DataTypes/Primitives"
    "github.com/sharpvik/scoops/Package/DataTypes/Slice"
    "github.com/sharpvik/scoops/Package/DataTypes/String"
    "github.com/sharpvik/scoops/Package/Shared"
    "github.com/sharpvik/scoops/Package/Util"
)



func (interpreter *Interpreter) Evaluate() {
    instruction := *(interpreter.code[interpreter.ip])
    //fmt.Println(instruction)
    //interpreter.scope.data.Print()
    switch instruction.opcode {

    case shared.END:
        interpreter.running = false

    case shared.PUSH_CONST:
        interpreter.scope.data.Push( primitives.NewByte(instruction.operand) )
    
    case shared.MAKE_BLN:
        b := interpreter.scope.data.Pop().(*primitives.Byte)
        interpreter.scope.data.Push( primitives.NewBoolean(b) )
    
    case shared.MAKE_INT:
        var buffer []byte
        for i := 0; i < 8; i++ {
            b := interpreter.scope.data.Pop().(*primitives.Byte).Value
            buffer = append(buffer, b)
        }
        n := int64( binary.LittleEndian.Uint64(buffer) )
        interpreter.scope.data.Push( primitives.NewInteger(n) )
    
    case shared.MAKE_NIL:
        interpreter.scope.data.Push(interpreter.thenil)
    
    case shared.MAKE_RUNE:
        var buffer []byte
        c := int(instruction.operand)
        for i := 0; i < c; i++ {
            b := interpreter.scope.data.Pop().(*primitives.Byte).Value
            buffer = append(buffer, b)
        }
        interpreter.scope.data.Push( primitives.NewRune(buffer) )
    
    case shared.MAKE_STRING:
        c := interpreter.scope.data.Pop().(*primitives.Integer).Value
        var i int64
        var buffer []*primitives.Rune
        for i = 0; i < c; i++ {
            r := interpreter.scope.data.Pop().(*primitives.Rune)
            buffer = append(buffer, r)
        }
        interpreter.scope.data.Push( _string.New(buffer) )
    
    case shared.MAKE_SLICE:
        interpreter.scope.data.Push( slice.New() )
    
    case shared.UNARY_NOT:
        b := interpreter.scope.data.Pop().(*primitives.Boolean)
        interpreter.scope.data.Push( primitives.NotBoolean(b) )
        
    case shared.BINARY_AND:
        a := interpreter.scope.data.Pop().(*primitives.Boolean)
        b := interpreter.scope.data.Pop().(*primitives.Boolean)
        interpreter.scope.data.Push( primitives.AndBoolean(a, b) )
        
    case shared.BINARY_OR:
        a := interpreter.scope.data.Pop().(*primitives.Boolean)
        b := interpreter.scope.data.Pop().(*primitives.Boolean)
        interpreter.scope.data.Push( primitives.OrBoolean(a, b) )
        
    case shared.BINARY_XOR:
        a := interpreter.scope.data.Pop().(*primitives.Boolean)
        b := interpreter.scope.data.Pop().(*primitives.Boolean)
        interpreter.scope.data.Push( primitives.XorBoolean(a, b) )

    case shared.BINARY_ADD:
        x := interpreter.scope.data.Pop().(shared.Object)
        y := interpreter.scope.data.Pop().(shared.Object)
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
        x := interpreter.scope.data.Pop().(shared.Object)
        y := interpreter.scope.data.Pop().(shared.Object)
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
        x := interpreter.scope.data.Pop().(shared.Object)
        y := interpreter.scope.data.Pop().(shared.Object)
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
        interpreter.scope.data.Peek().Print()

    case shared.PRINT_NEWLINE:
        fmt.Print("\n")

    case shared.POP:
        interpreter.scope.data.Pop()

    default:
        interpreter.err = errors.New("Unknown opcode.")

    }
    
    /* Instruction pointer is incremented by default. Use 'return' inside the
       case body to force it not to. */
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
