package bytes

import (
    "encoding/binary"
    "fmt"
    "github.com/sharpvik/scoops/Package/DataTypes/Primitives"
    "github.com/sharpvik/scoops/Package/DataTypes/Queue"
    "github.com/sharpvik/scoops/Package/DataTypes/Slice"
    "github.com/sharpvik/scoops/Package/DataTypes/Stack"
    "github.com/sharpvik/scoops/Package/DataTypes/String"
    "github.com/sharpvik/scoops/Package/Shared"
    "github.com/sharpvik/scoops/Package/Util"
    "math"
)



func (interpreter *Interpreter) Evaluate() {
    instruction := *(interpreter.code[interpreter.ip])

    //fmt.Println(instruction) // debug

    switch instruction.opcode {

    case shared.END:
        interpreter.running = false

    case shared.PUSH_CONST:
        interpreter.scope.data.Push( primitives.NewByte(instruction.operand) )

    case shared.MAKE_BLN:
        b := interpreter.scope.data.Pop().(*primitives.Byte)
        interpreter.scope.data.Push( primitives.NewBoolean(b) )

    /*
     * It's important to remember that the sequence of bytes that are pushed
     * onto the data stack to become a flt or an int in the future must be
     * generated using the BigEndian byte order. Stack reverses them and here
     * interpreter uses the LittleEndian to instantiate primitives.
     */
    case shared.MAKE_FLT:
        var buffer []byte
        for i := 0; i < 8; i++ {
            b := interpreter.scope.data.Pop().(*primitives.Byte).Value
            buffer = append(buffer, b)
        }
        bits := binary.LittleEndian.Uint64(buffer)
        float := math.Float64frombits(bits)
        interpreter.scope.data.Push( primitives.NewFloat(float) )

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

    case shared.STRING_CONCATENATE:
        b := interpreter.scope.data.Pop().(*_string.String)
        a := interpreter.scope.data.Pop().(*_string.String)
        c := _string.Concatenate(a, b)
        interpreter.scope.data.Push(c)

    case shared.MAKE_SLICE:
        interpreter.scope.data.Push( slice.New() )

    case shared.SLICE_APPEND:
        obj := interpreter.scope.data.Pop()
        slice := interpreter.scope.data.Peek().(*slice.Slice)
        slice.Append(obj)

    case shared.SLICE_GET_ITEM_BY_INDEX:
        i := interpreter.scope.data.Pop().(*primitives.Integer).Value
        index := uint64(i)
        slice := interpreter.scope.data.Peek().(*slice.Slice)
        obj := slice.GetItemByIndex(index)
        interpreter.scope.data.Push(obj)

    case shared.SLICE_POP:
        i := interpreter.scope.data.Pop().(*primitives.Integer).Value
        index := uint64(i)
        slice := interpreter.scope.data.Peek().(*slice.Slice)
        obj := slice.Pop(index)
        interpreter.scope.data.Push(obj)

    case shared.MAKE_STACK:
        interpreter.scope.data.Push( stack.New() )

    case shared.STACK_PUSH:
        obj := interpreter.scope.data.Pop()
        interpreter.scope.data.Peek().(*stack.Stack).Push(obj)

    case shared.STACK_POP:
        obj := interpreter.scope.data.Peek().(*stack.Stack).Pop()
        interpreter.scope.data.Push(obj)

    case shared.STACK_CLEAR:
        interpreter.scope.data.Peek().(*stack.Stack).Clear()

    case shared.STACK_EMPTY:
        bln := interpreter.scope.data.Peek().(*stack.Stack).Empty()
        interpreter.scope.data.Push( primitives.FromBoolean(bln) )

    case shared.STACK_PEEK:
        obj := interpreter.scope.data.Peek().(*stack.Stack).Peek()
        interpreter.scope.data.Push(obj)

    case shared.MAKE_QUEUE:
        interpreter.scope.data.Push( queue.New() )

    case shared.QUEUE_PUSH:
        obj := interpreter.scope.data.Pop()
        interpreter.scope.data.Peek().(*queue.Queue).Push(obj)

    case shared.QUEUE_POP:
        obj := interpreter.scope.data.Peek().(*queue.Queue).Pop()
        interpreter.scope.data.Push(obj)

    case shared.QUEUE_CLEAR:
        interpreter.scope.data.Peek().(*queue.Queue).Clear()

    case shared.QUEUE_EMPTY:
        bln := interpreter.scope.data.Peek().(*queue.Queue).Empty()
        interpreter.scope.data.Push( primitives.FromBoolean(bln) )

    case shared.QUEUE_PEEK:
        obj := interpreter.scope.data.Peek().(*queue.Queue).Peek()
        interpreter.scope.data.Push(obj)

    case shared.CLONE_OBJ:
        interpreter.scope.data.Push(
            interpreter.scope.data.Peek().Clone(),
        )

    case shared.PRINT_OBJ:
        interpreter.scope.data.Peek().Print(interpreter.writer)

    case shared.GET_TYPE:
        interpreter.scope.data.Push(
            _string.FromString( interpreter.scope.data.Peek().Type() ),
        )

    case shared.GET_SIZE:
        collection := interpreter.scope.data.Peek().(shared.Collection)
        size := int64( collection.Size() )
        interpreter.scope.data.Push( primitives.NewInteger(size) )

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
        x := interpreter.scope.data.Pop()
        y := interpreter.scope.data.Pop()
        _type := x.Type() + y.Type()
        switch _type {
        case "intint":
            a := x.(*primitives.Integer)
            b := y.(*primitives.Integer)
            interpreter.scope.data.Push( primitives.AddInteger(a, b) )
        default:
            interpreter.err = primitives.NewError(
                shared.RuntimeError,
                "Unknown numeric data type.",
            )
        }

    case shared.BINARY_SUB:
        x := interpreter.scope.data.Pop()
        y := interpreter.scope.data.Pop()
        _type := x.Type() + y.Type()
        switch _type {
        case "intint":
            a := x.(*primitives.Integer)
            b := y.(*primitives.Integer)
            c := primitives.SubInteger(a, b)
            interpreter.scope.data.Push(c)
        default:
            interpreter.err = primitives.NewError(
                shared.RuntimeError,
                "Unknown numeric data type.",
            )
        }

    case shared.BINARY_MUL:
        x := interpreter.scope.data.Pop()
        y := interpreter.scope.data.Pop()
        _type := x.Type() + y.Type()
        switch _type {
        case "intint":
            a := x.(*primitives.Integer)
            b := y.(*primitives.Integer)
            c := primitives.MulInteger(a, b)
            interpreter.scope.data.Push(c)
        default:
            interpreter.err = primitives.NewError(
                shared.RuntimeError,
                "Unknown numeric data type.",
            )
        }

    case shared.PRINT_NEWLINE:
        interpreter.writer.WriteString("\n")
        interpreter.writer.Flush()

    case shared.POP:
        interpreter.scope.data.Pop()

    default:
        interpreter.err = primitives.NewError(
            shared.OpcodeError,
            fmt.Sprintf(
                "Instruction #%d: Unknown opcode %d.",
                interpreter.ip,
                instruction.opcode,
            ),
        )

    }

    /* Instruction pointer is incremented by default. Use 'return' inside the
       case body to force it not to. */
    interpreter.ip++

    //interpreter.scope.data.Print(interpreter.writer) // debug
}


func (interpreter *Interpreter) Execute() {
    for interpreter.running && interpreter.err == nil {
        interpreter.Evaluate()
    }
    if interpreter.err != nil {
        interpreter.err.Print(interpreter.writer)
        util.Log("Interpreter exited with non-zero return value.")
    }
}
