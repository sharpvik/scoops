package bytes

import (
    //"errors"
    "fmt"
    //"github.com/sharpvik/scoops/Package/DataTypes/Primitives"
    "github.com/sharpvik/scoops/Package/Shared"
    "github.com/sharpvik/scoops/Package/Util"
)



func (interpreter *Interpreter) Evaluate() {
    inst := *(interpreter.code[interpreter.ip])
    switch inst.opcode {
    case shared.END:
        fmt.Println(inst)
        interpreter.running = false
    default:
        fmt.Println(inst)
        interpreter.ip++
    }
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
